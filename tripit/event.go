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

Airline: %s %s

Departing Terminal %s Gate %s

Arrive -> %s (%s)
%s

Duration: %s

Distance: %s

Check-in URL: %s

View and/or edit details of this flight [%s]: https://www.tripit.com/%s

View and/or edit details of this trip: https://www.tripit.com/trip/show/id/%s`

	eventColorID  = "3"
	bufferColorID = "8"
)

// Event holds the data we will use when creating calendar events for flights, activities, and other
// TripIt API objects.
type Event struct {
	Title              string
	Description        string
	AirportCode        string
	Start              calendar.EventDateTime
	End                calendar.EventDateTime
	ID                 string
	SegmentID          string
	ConfirmationNumber string
	ColorID            string
}

// GetFlightSegmentsAsEvents returns an Event object for each of the
// flight segments in the given flight object.
func (f Flight) GetFlightSegmentsAsEvents() ([]Event, error) {
	// Initialize our events array.
	events := []Event{}

	// Iterate over the flight segments.
	for i := 0; i < len(f.Segments); i++ {
		segment := f.Segments[i]

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

		// Sort out operating versus marketing airline
		var airlineName, airlineCode, flightNumber string
		if segment.OperatingAirline != "" {
			airlineName = segment.OperatingAirline
			airlineCode = segment.OperatingAirlineCode
			flightNumber = segment.OperatingFlightNumber
		} else if segment.MarketingAirline != "" {
			airlineName = segment.MarketingAirline
			airlineCode = segment.MarketingAirlineCode
			flightNumber = segment.MarketingFlightNumber
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
			airlineName,
			flightNumber,
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

		var confirmationNumber string
		if f.SupplierConfNum != "" {
			confirmationNumber = f.SupplierConfNum
		} else if f.BookingSiteConfNum != "" {
			confirmationNumber = f.BookingSiteConfNum
		}

		// Append the event to our events array.
		events = append(events, Event{
			Title:              fmt.Sprintf("Flight to %s (%s %s)", segment.EndCityName, airlineCode, flightNumber),
			Description:        description,
			AirportCode:        segment.StartAirportCode,
			Start:              start,
			End:                end,
			ID:                 f.TripID,
			SegmentID:          segment.ID,
			ConfirmationNumber: confirmationNumber,
			ColorID:            eventColorID,
		})

		// If we have the first item in the segment, create the buffer event
		// for travel time to the airport.
		if i == 0 {
			events = append(events, Event{
				Title:       fmt.Sprintf("Buffer for travel time to %s & security", segment.StartAirportCode),
				Description: description,
				AirportCode: segment.StartAirportCode,
				Start: calendar.EventDateTime{
					DateTime: startDate.Add(-3 * time.Hour).Format(time.RFC3339),
					TimeZone: segment.StartDateTime.Timezone,
				},
				End: calendar.EventDateTime{
					DateTime: startDate.Format(time.RFC3339),
					TimeZone: segment.StartDateTime.Timezone,
				},
				ID:                 f.TripID,
				SegmentID:          segment.ID,
				ConfirmationNumber: confirmationNumber,
				ColorID:            bufferColorID,
			})
		}
		// If we have the last item in the segment, create the buffer event
		// for travel time from the airport.
		if i == len(f.Segments)-1 {
			events = append(events, Event{
				Title:       fmt.Sprintf("Buffer for travel time from %s", segment.EndAirportCode),
				Description: description,
				// TODO: put the hotel as location here.
				AirportCode: "",
				Start: calendar.EventDateTime{
					DateTime: endDate.Format(time.RFC3339),
					TimeZone: segment.EndDateTime.Timezone,
				},
				End: calendar.EventDateTime{
					DateTime: endDate.Add(2 * time.Hour).Format(time.RFC3339),
					TimeZone: segment.EndDateTime.Timezone,
				},
				ID:                 f.TripID,
				SegmentID:          segment.ID,
				ConfirmationNumber: confirmationNumber,
				ColorID:            bufferColorID,
			})
		}
	}

	return events, nil
}
