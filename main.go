package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jessfraz/tripitcalb0t/tripit"
	"github.com/jessfraz/tripitcalb0t/version"
	"github.com/sirupsen/logrus"
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
	tripitUsername string
	tripitToken    string

	interval string
	once     bool

	debug bool
	vrsn  bool
)

func init() {
	// parse flags
	flag.StringVar(&tripitUsername, "tripit-username", os.Getenv("TRIPIT_USERNAME"), "TripIt Username for authentication (or env var TRIPIT_USERNAME)")
	flag.StringVar(&tripitToken, "tripit-token", os.Getenv("TRIPIT_TOKEN"), "TripIt Token for authentication (or env var TRIPIT_TOKEN)")

	flag.StringVar(&interval, "interval", "1m", "update interval (ex. 5ms, 10s, 1m, 3h)")
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

}

func main() {
	var ticker *time.Ticker

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

	// Parse the duration.
	dur, err := time.ParseDuration(interval)
	if err != nil {
		logrus.Fatalf("parsing %s as duration failed: %v", interval, err)
	}
	ticker = time.NewTicker(dur)

	// Create the TripIt API client.
	client := tripit.New(tripitUsername, tripitToken)

	// If the user passed the once flag, just do the run once and exit.
	if once {
		run(client)
		logrus.Info("Updated TripIt calendar entries")
		os.Exit(0)
	}

	logrus.Infof("Starting bot to update TripIt calendar entries every %s", interval)
	for range ticker.C {
		run(client)
	}
}

func run(client *tripit.Client) {
	// Get a list of trips.
	resp, err := client.ListTrips(
		tripit.Filter{
			Type:  tripit.FilterPast,
			Value: "true",
		},
		tripit.Filter{
			Type:  tripit.FilterIncludeObjects,
			Value: "true",
		})
	if err != nil {
		logrus.Fatal(err)
	}

	// Iterate over our flights and create/update calendar entries in Google calendar.
	for _, flight := range resp.Flights {
		// Create the events for the flight.
		events, err := flight.GetFlightSegmentsAsEvents()
		if err != nil {
			// Warn on error and continue iterating through the flights.
			logrus.Warn(err)
			continue
		}

		logrus.Infof("events: %#v", events)

		// Create / Update a Google Calendar entry for each event.
		// TODO(jessfraz): do this.
	}
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
