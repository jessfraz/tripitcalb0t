package tripit

import (
	"fmt"
	"strings"
	"time"

	calendar "google.golang.org/api/calendar/v3"
)

const (
	eventDescriptionFormat = `[Flight] %s to %s
%s

Booking Site (%s) Confirmation # %s
Supplier (%s) Confirmation # %s
Record Locator # %s

Operating Airline: %s %s

Departing Terminal %s Gate %s

Arrive -> %s (%s)
%s

Duration: %s

Distance: %s

Check-in URL: %s

View and/or edit details of this flight [%s]: https://www.tripit.com/%s

View and/or edit details of this trip: https://www.tripit.com/trip/show/id/%s`
)

// Event holds the data we will use when creating calendar events for flights, activities, and other
// TripIt API objects.
type Event struct {
	Title             string
	Description       string
	LocationLatitude  float64
	LocationLongitude float64
	Start             calendar.EventDateTime
	End               calendar.EventDateTime
	ID                string
	SegmentID         string
}

// GetFlightSegmentsAsEvents returns an Event object for each of the
// flight segments in the given flight object.
func (f Flight) GetFlightSegmentsAsEvents() ([]Event, error) {
	// Initialize our events array.
	events := []Event{}

	// Iterate over the flight segments.
	for _, segment := range f.Segments {
		// Get the flight start time.
		startDate, err := segment.StartDateTime.Parse()
		if err != nil {
			return nil, fmt.Errorf("parsing StartDateTime for tripID -> %s, segment -> %s, from %s -> %s failed: %v", f.TripID, segment.ID, segment.StartAirportCode, segment.EndAirportCode, err)
		}
		start := calendar.EventDateTime{
			DateTime: startDate.Format(time.RFC3339),
			TimeZone: segment.StartDateTime.Timezone,
		}

		// Get the flight end time.
		endDate, err := segment.EndDateTime.Parse()
		if err != nil {
			return nil, fmt.Errorf("parsing EndDateTime for tripID -> %s, segment -> %s, from %s -> %s failed: %v", f.TripID, segment.ID, segment.StartAirportCode, segment.EndAirportCode, err)
		}
		end := calendar.EventDateTime{
			DateTime: endDate.Format(time.RFC3339),
			TimeZone: segment.EndDateTime.Timezone,
		}

		// Create a description for the flight segment.
		description := fmt.Sprintf(eventDescriptionFormat,
			segment.StartAirportCode,
			segment.EndAirportCode,
			startDate.Format(time.RFC1123Z),
			f.BookingSiteName,
			f.BookingSiteConfNum,
			f.SupplierName,
			f.SupplierConfNum,
			f.RecordLocator,
			segment.OperatingAirline,
			segment.OperatingFlightNumber,
			segment.StartTerminal,
			segment.StartGate,
			segment.EndCityName,
			segment.EndAirportCode,
			endDate.Format(time.RFC1123Z),
			segment.Duration,
			segment.Distance,
			segment.CheckInURL,
			segment.ID,
			strings.TrimPrefix(f.RelativeURL, "/"),
			f.TripID)

		// Append the event to our events array.
		events = append(events, Event{
			Title:             fmt.Sprintf("Flight to %s", segment.EndCityName),
			Description:       description,
			LocationLatitude:  segment.StartAirportLatitude,
			LocationLongitude: segment.StartAirportLongitude,
			Start:             start,
			End:               end,
			ID:                f.TripID,
			SegmentID:         segment.ID,
		})
	}

	return events, nil
}
