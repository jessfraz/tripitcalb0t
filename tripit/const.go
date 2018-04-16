package tripit

import (
	"fmt"
	"strings"
)

const (
	// APIUri holds the TripIt API uri.
	APIUri = "https://api.tripit.com"
	// APIVersion holds the TripIt API version.
	APIVersion = "v1"

	// ListTripsEndpoint is the API endoint to list trips.
	ListTripsEndpoint = "/list/trip"
	// ListObjectsEndpoint is the API endoint to list objects.
	ListObjectsEndpoint = "/list/object"
	// ListPointsProgramsEndpoint is the API endoint to list points programs.
	ListPointsProgramsEndpoint = "/list/points_program"

	// EndpointFormatGetObject is the endpoint format to get an object.
	EndpointFormatGetObject = "get/%s/id/%s/%s"
	// EndpointFormatDeleteObject is the endpoint format to delete an object.
	EndpointFormatDeleteObject = "delete/%s/id/%s"
	// EndpointFormatReplaceObject is the endpoint format to replace an object.
	EndpointFormatReplaceObject = "replace/%s/id/%s"

	// FilterNone is the filter type representing no filter.
	FilterNone TypeFilter = "" // valid on trip, object, points_program
	// FilterTraveler is the filter type for filtering on travelers.
	FilterTraveler TypeFilter = "traveler" // valid on trip, object. Values: true, false, all
	// FilterPast is the filter type for filtering over past objects.
	FilterPast TypeFilter = "past" // valid on trip, object. Values: true, false
	// FilterModifiedSince is the filter type for filtering objects based on when they were modified.
	FilterModifiedSince TypeFilter = "modified_since" // valid on trip, object. Values: integer
	// FilterIncludeObjects is the filter type for including objects.
	FilterIncludeObjects TypeFilter = "include_objects" // valid on trip. Values: true, false
	// FilterTripID is the filter type for filtering by trip ID.
	FilterTripID TypeFilter = "trip_id" // valid on object. Values: integer trip id
	// FilterType is the filter type for filtering by object type.
	FilterType TypeFilter = "type" // valid on object. Values: all object types
	// FilterPageNum is the filter type for the page number.
	FilterPageNum TypeFilter = "page_num"
	// FilterPageSize is the filter type for the page size.
	FilterPageSize TypeFilter = "page_size"

	// TypeActivity is the activity object type.
	TypeActivity Type = "activity"
	// TypeCar is the car object type.
	TypeCar Type = "car"
	// TypeCruise is the cruise object type.
	TypeCruise Type = "cruise"
	// TypeDirections is the directions object type.
	TypeDirections Type = "directions"
	// TypeFlight is the air object type.
	TypeFlight Type = "air"
	// TypeLodging is the lodging object type.
	TypeLodging Type = "lodging"
	// TypeMap is the map object type.
	TypeMap Type = "map"
	// TypeNote is the note object type.
	TypeNote Type = "note"
	// TypePointsProgram is the points_program object type.
	TypePointsProgram Type = "points_program"
	// TypeProfile is the profile object type.
	TypeProfile Type = "profile"
	// TypeRail is the rail object type.
	TypeRail Type = "rail"
	// TypeRestaurant is the restaurant object type.
	TypeRestaurant Type = "restaurant"
	// TypeSegment is the segment object type.
	TypeSegment Type = "segment"
	// TypeTransport is the transport object type.
	TypeTransport Type = "transport"
	// TypeTrip is the trip object type.
	TypeTrip Type = "trip"
	// TypeWeather is the weather object type.
	TypeWeather Type = "weather"

	// FlightStatusNotMonitorable means an AirSegment in this state is usually in this state because the TripIt platform doesn't have enough information about the flight to monitor it.
	FlightStatusNotMonitorable FlightStatusCode = 100

	// FlightStatusNotMonitored means the flight is either not a real flight or TripIt doesn't have a record of it in its database.
	FlightStatusNotMonitored FlightStatusCode = 200

	// FlightStatusScheduled means the flight is monitored but the TripIt platform hasn't seen any updates to the flight yet.
	FlightStatusScheduled FlightStatusCode = 300
	// FlightStatusOnTime means the flight is currently considered to be on time.
	FlightStatusOnTime FlightStatusCode = 301
	// FlightStatusInFlightOnTime means the flight is currently in the air and is considered to be on time.
	FlightStatusInFlightOnTime FlightStatusCode = 302
	// FlightStatusArrivedOnTime means the flight has arrived and the actual arrival time was within 14 minutes of the ScheduledArrivalDateTime.
	FlightStatusArrivedOnTime FlightStatusCode = 303

	// FlightStatusCancelled means the flight has been cancelled.
	FlightStatusCancelled FlightStatusCode = 400
	// FlightStatusDelayed means the flight is currently considered to be delayed.
	FlightStatusDelayed FlightStatusCode = 401
	// FlightStatusInFlightLate means the flight is currently in the air and is considered to be late.
	FlightStatusInFlightLate FlightStatusCode = 402
	// FlightStatusArrivedLate means the flight has arrived and the actual arrival time was more than 14 minutes after the ScheduledArrivalDateTime.
	FlightStatusArrivedLate FlightStatusCode = 403
	// FlightStatusDiverted means the flight was diverted to an airport other than the one it was originally going to fly into.
	FlightStatusDiverted FlightStatusCode = 404
	// FlightStatusPossiblyDelayed means the flight has possibly been delayed.
	FlightStatusPossiblyDelayed FlightStatusCode = 405
	// FlightStatusInFlightPossiblyLate means the flight is in flight and may possibly be late.
	FlightStatusInFlightPossiblyLate FlightStatusCode = 406
	// FlightStatusArrivedPossiblyLate means the flight has arrived but may possibly be late.
	FlightStatusArrivedPossiblyLate FlightStatusCode = 407
	// FlightStatusUnknown means the flight status is unknown.
	FlightStatusUnknown FlightStatusCode = 408

	// TransportDetailTypeFerry is the detail type code for ferry transport.
	TransportDetailTypeFerry DetailTypeCode = "F"
	// TransportDetailTypeGroundTransportation is the detail type code for ground transport.
	TransportDetailTypeGroundTransportation DetailTypeCode = "G"

	// ActivityDetailTypeConcert is the activity detail type code for a concert.
	ActivityDetailTypeConcert DetailTypeCode = "C"
	// ActivityDetailTypeTheatre is the activity detail type code for the theatre.
	ActivityDetailTypeTheatre DetailTypeCode = "H"
	// ActivityDetailTypeMeeting is the activity detail type code for a meeting.
	ActivityDetailTypeMeeting DetailTypeCode = "M"
	// ActivityDetailTypeTour is the activity detail type code for a tour.
	ActivityDetailTypeTour DetailTypeCode = "T"

	// CruiseDetailTypePortOfCall is the detail type code for a cruise port of call.
	CruiseDetailTypePortOfCall DetailTypeCode = "P"

	// NoteDetailTypeArticle is the note detail type code for an article.
	NoteDetailTypeArticle DetailTypeCode = "A"
)

// Type defines the type for an object.
type Type string

// DetailTypeCode defines the detail type code for an object.
type DetailTypeCode string

// FlightStatusCode defines the type for a flight status code.
type FlightStatusCode int

// TypeFilter defines the type for a filter type.
type TypeFilter string

// Filter holds the information about a filter including the type and value.
type Filter struct {
	Type  TypeFilter
	Value string
}

// String returns the string representation of a filter.
func (f Filter) String() string {
	if f.Type == FilterNone {
		return ""
	}

	return fmt.Sprintf("%s/%s/", f.Type, f.Value)
}

// formatFilters converts an array of Filter objects into the correct format for URL parameters.
func formatFilters(filters []Filter) (s string) {
	for _, filter := range filters {
		s += filter.String()
	}

	// Trim any leading or trailing slashes.
	return strings.Trim(s, "/")
}
