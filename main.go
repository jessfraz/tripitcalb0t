package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/jessfraz/tripitcalb0t/tripit"
	"github.com/jessfraz/tripitcalb0t/version"
	"github.com/mmcloughlin/openflights"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	calendar "google.golang.org/api/calendar/v3"
)

const (
	// BANNER is what is printed for help/info output.
	BANNER = ` _        _       _ _            _ _      ___  _
| |_ _ __(_)_ __ (_) |_ ___ __ _| | |__  / _ \| |_
| __| '__| | '_ \| | __/ __/ _` + "`" + ` | | '_ \| | | | __|
| |_| |  | | |_) | | || (_| (_| | | |_) | |_| | |_
 \__|_|  |_| .__/|_|\__\___\__,_|_|_.__/ \___/ \__|
           |_|

 Bot to automatically create Google Calendar events from TripIt flight data.
 Version: %s
 Build: %s

`
)

var (
	googleCalendarKeyfile string
	calendarName          string
	credsDir              string

	tripitUsername string
	tripitToken    string

	interval time.Duration
	once     bool

	debug bool
	vrsn  bool
)

func init() {
	// Get home directory.
	home, err := getHome()
	if err != nil {
		logrus.Fatal(err)
	}
	credsDir = filepath.Join(home, ".tripitcalb0t")

	// parse flags
	flag.StringVar(&googleCalendarKeyfile, "google-keyfile", filepath.Join(credsDir, "google.json"), "Path to Google Calendar keyfile")
	flag.StringVar(&calendarName, "calendar", os.Getenv("GOOGLE_CALENDAR_ID"), "Calendar name to add events to (or env var GOOGLE_CALENDAR_ID)")

	flag.StringVar(&tripitUsername, "tripit-username", os.Getenv("TRIPIT_USERNAME"), "TripIt Username for authentication (or env var TRIPIT_USERNAME)")
	flag.StringVar(&tripitToken, "tripit-token", os.Getenv("TRIPIT_TOKEN"), "TripIt Token for authentication (or env var TRIPIT_TOKEN)")

	flag.DurationVar(&interval, "interval", time.Minute, "update interval (ex. 5ms, 10s, 1m, 3h)")
	flag.BoolVar(&once, "once", false, "run once and exit, do not run as a daemon")

	flag.BoolVar(&vrsn, "version", false, "print version and exit")
	flag.BoolVar(&vrsn, "v", false, "print version and exit (shorthand)")
	flag.BoolVar(&debug, "d", false, "run in debug mode")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(BANNER, version.VERSION, version.GITCOMMIT))
		flag.PrintDefaults()
	}

	flag.Parse()

	if vrsn {
		fmt.Printf("tripitcalb0t version %s, build %s", version.VERSION, version.GITCOMMIT)
		os.Exit(0)
	}

	// set log level
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if tripitUsername == "" {
		usageAndExit("tripit username cannot be empty", 1)
	}

	if tripitToken == "" {
		usageAndExit("tripit token cannot be empty", 1)
	}

	if _, err := os.Stat(googleCalendarKeyfile); os.IsNotExist(err) {
		usageAndExit(fmt.Sprintf("Google Calendar keyfile %q does not exist", googleCalendarKeyfile), 1)
	}

	if calendarName == "" {
		usageAndExit("calendar name cannot be empty", 1)
	}
}

func main() {
	ticker := time.NewTicker(interval)

	// On ^C, or SIGTERM handle exit.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for sig := range c {
			ticker.Stop()
			logrus.Infof("Received %s, exiting.", sig.String())
			os.Exit(0)
		}
	}()

	// Create the TripIt API client.
	tripitClient := tripit.New(tripitUsername, tripitToken)

	// Create the Google calendar API client.
	gcalData, err := ioutil.ReadFile(googleCalendarKeyfile)
	if err != nil {
		logrus.Fatalf("reading file %s failed: %v", googleCalendarKeyfile, err)
	}
	gcalTokenSource, err := google.JWTConfigFromJSON(gcalData, calendar.CalendarScope)
	if err != nil {
		logrus.Fatalf("creating google calendar token source from file %s failed: %v", googleCalendarKeyfile, err)
	}

	// Create our context.
	ctx := context.Background()

	// Create the Google calendar client.
	gcalClient, err := calendar.New(gcalTokenSource.Client(ctx))
	if err != nil {
		logrus.Fatalf("creating google calendar client failed: %v", err)
	}

	// If the user passed the once flag, just do the run once and exit.
	if once {
		run(tripitClient, gcalClient, calendarName)
		logrus.Infof("Updated TripIt calendar entries in Google calendar %s", calendarName)
		os.Exit(0)
	}

	logrus.Infof("Starting bot to update TripIt calendar entries in Google calendar %s every %s", calendarName, interval)
	for range ticker.C {
		run(tripitClient, gcalClient, calendarName)
	}
}

func run(tripitClient *tripit.Client, gcalClient *calendar.Service, calendarName string) {
	// Get a list of events from Google calendar.
	t := time.Now().AddDate(-4, 0, 0).Format(time.RFC3339)
	events, err := gcalClient.Events.List(calendarName).ShowDeleted(false).SingleEvents(true).TimeMin(t).OrderBy("startTime").Q("Flight").MaxResults(2500).Do()
	if err != nil {
		logrus.Fatalf("getting events from google calendar %s failed: %v", calendarName, err)
	}

	trips, err := getTripItEvents(tripitClient, 1, "true")
	if err != nil {
		logrus.Fatalf("getting tripit events failed: %v", err)
	}

	// Iterate over the trip and see if we already have a matching calendar event.
	// If not make one and/or update the old one.
	for _, trip := range trips {
		if trip.ConfirmationNumber == "" {
			logrus.Warnf("skipping trip that has no confirmation number: %#v", trip)
			continue
		}

		var matchingEvent *calendar.Event
		for _, e := range events.Items {
			// We only care about TripIt events that match our tripID or segmentID.
			if (strings.Contains(strings.ToLower(e.Description), "tripit") ||
				strings.Contains(strings.ToLower(e.Summary), "flight")) &&
				strings.Contains(e.Description, trip.SegmentID) {
				matchingEvent = e
				break
			}
		}

		// Get airport information.
		airport := getAirportName(trip.AirportCode)
		if airport == "" {
			logrus.Errorf("getting airport information from iata database for %s returned no match", trip.AirportCode)
			continue
		}

		if matchingEvent == nil {
			// No event was found for this trip, let's create one.
			matchingEvent = &calendar.Event{
				Summary:     trip.Title,
				Description: trip.Description,
				Start:       &trip.Start,
				End:         &trip.End,
				Location:    airport,
			}

			// Insert the event.
			_, err = gcalClient.Events.Insert(calendarName, matchingEvent).Do()
			if err != nil {
				logrus.Errorf("inserting google calendar event failed: %v", err)
			}
			continue
		}

		// Update our matching event.
		matchingEvent.Summary = trip.Title
		matchingEvent.Description = trip.Description
		matchingEvent.Start = &trip.Start
		matchingEvent.End = &trip.End
		matchingEvent.Location = airport

		// Update the event.
		_, err = gcalClient.Events.Update(calendarName, matchingEvent.Id, matchingEvent).Do()
		if err != nil {
			logrus.Errorf("updating google calendar event %s failed: %v", matchingEvent.Id, err)
		}
	}
}

func getTripItEvents(tripitClient *tripit.Client, page int, pastFilter string) ([]tripit.Event, error) {
	// Get a list of trips.
	resp, err := tripitClient.ListTrips(
		tripit.Filter{
			Type:  tripit.FilterPast,
			Value: pastFilter,
		},
		tripit.Filter{
			Type:  tripit.FilterIncludeObjects,
			Value: "true",
		},
		tripit.Filter{
			Type:  tripit.FilterPageNum,
			Value: fmt.Sprintf("%d", page),
		},
		tripit.Filter{
			Type:  tripit.FilterPageSize,
			Value: "25",
		})
	if err != nil {
		return nil, fmt.Errorf("listing trips from TripIt failed: %v", err)
	}

	var events []tripit.Event

	// Iterate over our flights and create/update calendar entries in Google calendar.
	for _, flight := range resp.Flights {
		// Create the events for the flight.
		evs, err := flight.GetFlightSegmentsAsEvents()
		if err != nil {
			// Warn on error and continue iterating through the flights.
			logrus.Warn(err)
			continue
		}

		// Add to our events array.
		events = append(events, evs...)
	}

	// Paginate.
	pageNum, err := strconv.Atoi(resp.PageNum)
	if err != nil {
		return nil, err
	}
	maxPage, err := strconv.Atoi(resp.MaxPage)
	if err != nil {
		return nil, err
	}

	if pageNum < maxPage {
		pageNum++

		evs, err := getTripItEvents(tripitClient, pageNum, pastFilter)
		if err != nil {
			return nil, err
		}

		return append(events, evs...), nil
	}

	if pastFilter == "true" {
		// Get future events as well.
		evs, err := getTripItEvents(tripitClient, 1, "false")
		if err != nil {
			return nil, err
		}

		return append(events, evs...), nil
	}

	return events, nil
}

func getAirportName(code string) string {
	for _, airport := range openflights.Airports {
		if airport.IATA == code {
			return airport.Name
		}
	}

	return ""
}

func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}

func getHome() (string, error) {
	home := os.Getenv(homeKey)
	if home != "" {
		return home, nil
	}

	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return u.HomeDir, nil
}
