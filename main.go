package main

import (
	"context"
	"errors"
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

	"github.com/genuinetools/pkg/cli"
	"github.com/jessfraz/tripitcalb0t/tripit"
	"github.com/jessfraz/tripitcalb0t/version"
	"github.com/mmcloughlin/openflights"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	calendar "google.golang.org/api/calendar/v3"
)

var (
	googleCalendarKeyfile string
	calendarName          string
	credsDir              string
  	pastFilter            string

	tripitUsername string
	tripitPassword string

	interval time.Duration
	once     bool
  	past     bool

	debug bool
)

func main() {
	// Get home directory.
	home, err := getHome()
	if err != nil {
		logrus.Fatal(err)
	}
	credsDir = filepath.Join(home, ".tripitcalb0t")

	// Create a new cli program.
	p := cli.NewProgram()
	p.Name = "tripitcalb0t"
	p.Description = "Bot to automatically create Google Calendar events from TripIt flight data"

	// Set the GitCommit and Version.
	p.GitCommit = version.GITCOMMIT
	p.Version = version.VERSION

	// Setup the global flags.
	p.FlagSet = flag.NewFlagSet("global", flag.ExitOnError)
	p.FlagSet.StringVar(&googleCalendarKeyfile, "google-keyfile", filepath.Join(credsDir, "google.json"), "Path to Google Calendar keyfile")
	p.FlagSet.StringVar(&calendarName, "calendar", os.Getenv("GOOGLE_CALENDAR_ID"), "Calendar name to add events to (or env var GOOGLE_CALENDAR_ID)")

	p.FlagSet.StringVar(&tripitUsername, "tripit-username", os.Getenv("TRIPIT_USERNAME"), "TripIt Username for authentication (or env var TRIPIT_USERNAME)")
	p.FlagSet.StringVar(&tripitPassword, "tripit-password", os.Getenv("TRIPIT_PASSWORD"), "TripIt Password for authentication (or env var TRIPIT_PASSWORD)")


	p.FlagSet.DurationVar(&interval, "interval", time.Minute, "Update interval (ex. 5ms, 10s, 1m, 3h)")
  	p.FlagSet.BoolVar(&once, "once", false, "Run once and exit, do not run as a daemon")
  	p.FlagSet.BoolVar(&past, "past", false, "Include past trips")

	p.FlagSet.BoolVar(&debug, "d", false, "Enable debug logging")

	// Set the before function.
	p.Before = func(ctx context.Context) error {
		// Set the log level.
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		if len(tripitUsername) < 1 {
			return errors.New("tripit username cannot be empty")
		}

		if len(tripitPassword) < 1 {
			return errors.New("tripit password cannot be empty")
		}

		if _, err := os.Stat(googleCalendarKeyfile); os.IsNotExist(err) {
			return fmt.Errorf("Google Calendar keyfile %q does not exist", googleCalendarKeyfile)
		}

		if len(calendarName) < 1 {
			return errors.New("calendar name cannot be empty")
		}

		return nil
	}

	// Set the main program action.
	p.Action = func(ctx context.Context, args []string) error {
		ticker := time.NewTicker(interval)

		// On ^C, or SIGTERM handle exit.
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		signal.Notify(c, syscall.SIGTERM)
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		go func() {
			for sig := range c {
				cancel()
				ticker.Stop()
				logrus.Infof("Received %s, exiting.", sig.String())
				os.Exit(0)
			}
		}()

		// Create the TripIt API client.
		tripitClient := tripit.New(tripitUsername, tripitPassword)

		// Create the Google calendar API client.
		gcalData, err := ioutil.ReadFile(googleCalendarKeyfile)
		if err != nil {
			logrus.Fatalf("reading file %s failed: %v", googleCalendarKeyfile, err)
		}
		gcalTokenSource, err := google.JWTConfigFromJSON(gcalData, calendar.CalendarScope)
		if err != nil {
			logrus.Fatalf("creating google calendar token source from file %s failed: %v", googleCalendarKeyfile, err)
		}

		// Create the Google calendar client.
		gcalClient, err := calendar.New(gcalTokenSource.Client(ctx))
		if err != nil {
			logrus.Fatalf("creating google calendar client failed: %v", err)
		}

    	pastFilter := fmt.Sprintf("%v", past)

		// If the user passed the once flag, just do the run once and exit.

    	if once {
			run(tripitClient, gcalClient, calendarName, pastFilter)
			logrus.Infof("Updated TripIt calendar entries in Google calendar %s", calendarName)
			os.Exit(0)
		}

		logrus.Infof("Starting bot to update TripIt calendar entries in Google calendar %s every %s", calendarName, interval)
		for range ticker.C {
			run(tripitClient, gcalClient, calendarName, pastFilter)
		}

		return nil
	}

	// Run our program.
	p.Run()
}

func run(tripitClient *tripit.Client, gcalClient *calendar.Service, calendarName string, pastFilter string) {
	// Get a list of events from Google calendar.
	t := time.Now().AddDate(-4, 0, 0).Format(time.RFC3339)
	events, err := gcalClient.Events.List(calendarName).ShowDeleted(false).SingleEvents(true).TimeMin(t).OrderBy("startTime").Q("Flight").MaxResults(2500).Do()
	if err != nil {
		logrus.Fatalf("getting events from google calendar %s failed: %v", calendarName, err)
	}

	trips, err := getTripItEvents(tripitClient, 1, pastFilter)
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
