package tripit

import (
	"encoding/json"
	"fmt"
	"time"
)

// Request contains the objects that can be sent to the TripIt API in a request.
type Request struct {
	Invitations []Invitation `json:"Invitation,omitempty"`       // optional
	Trip        Trip         `json:"Trip,omitempty"`             // optional
	Activity    Activity     `json:"ActivityObject,omitempty"`   // optional
	Car         Car          `json:"CarObject,omitempty"`        // optional
	Cruise      Cruise       `json:"CruiseObject,omitempty"`     // optional
	Directions  Direction    `json:"DirectionsObject,omitempty"` // optional
	Flight      Flight       `json:"AirObject,omitempty"`        // optional
	Lodging     Lodging      `json:"LodgingObject,omitempty"`    // optional
	Map         Map          `json:"MapObject,omitempty"`        // optional
	Note        Note         `json:"NoteObject,omitempty"`       // optional
	Rail        Rail         `json:"RailObject,omitempty"`       // optional
	Restaurant  Restaurant   `json:"RestaurantObject,omitempty"` // optional
	Transport   Transport    `json:"TransportObject,omitempty"`  // optional
}

// Response represents the TripIt API Response.
type Response struct {
	Timestamp string `json:"timestamp,omitempty" xml:"timestamp"`
	NumBytes  int    `json:"num_bytes,string,omitempty" xml:"num_bytes"`

	Errors   []Error   `json:"Error,omitempty" xml:"Error"`     // optional
	Warnings []Warning `json:"Warning,omitempty" xml:"Warning"` // optional

	Activities  Activities  `json:"ActivityObject,omitempty" xml:"ActivityObject"`     // optional
	Flights     Flights     `json:"AirObject,omitempty" xml:"AirObject"`               // optional
	Cars        Cars        `json:"CarObject,omitempty" xml:"CarObject"`               // optional
	Cruises     Cruises     `json:"CruiseObject,omitempty" xml:"CruiseObject"`         // optional
	Directions  Directions  `json:"DirectionsObject,omitempty" xml:"DirectionsObject"` // optional
	Lodging     Lodges      `json:"LodgingObject,omitempty" xml:"LodgingObject"`       // optional
	Maps        Maps        `json:"MapObject,omitempty" xml:"MapObject"`               // optional
	Notes       Notes       `json:"NoteObject,omitempty" xml:"NoteObject"`             // optional
	Rails       Rails       `json:"RailObject,omitempty" xml:"RailObject"`             // optional
	Restaurants Restaurants `json:"RestaurantObject,omitempty" xml:"RestaurantObject"` // optional

	Transports     Transports      `json:"TransportObject,omitempty" xml:"TransportObject"` // optional
	Trips          Trips           `json:"Trip,omitempty" xml:"Trip"`                       // optional
	Weather        []Weather       `json:"WeatherObject,omitempty" xml:"WeatherObject"`     // optional
	PointsPrograms []PointsProgram `json:"PointsProgram,omitempty" xml:"PointsProgram"`     // optional
	Profiles       []Profile       `json:"Profile,omitempty" xml:"Profile"`                 // optional

	PageNum  string `json:"page_num,omitempty"`
	PageSize string `json:"page_size,omitempty"`
	MaxPage  string `json:"max_page,omitempty"`
}

// Error is returned from TripIt on error conditions.
type Error struct {
	Code              int     `json:"code,string,omitempty" xml:"code"`                               // read-only
	DetailedErrorCode float64 `json:"detailed_error_code,string,omitempty" xml:"detailed_error_code"` // optional, read-only
	Description       string  `json:"description,omitempty" xml:"description"`                        // read-only
	EntityType        string  `json:"entity_type,omitempty" xml:"entity_type"`                        // read-only
	Timestamp         string  `json:"timestamp,omitempty" xml:"timestamp"`                            // read-only, xs:datetime
}

// Warning is returned from TripIt to indicate warning conditions.
type Warning struct {
	Description string `json:"description,omitempty"` // read-only
	EntityType  string `json:"entity_type,omitempty"` // read-only
	Timestamp   string `json:"timestamp,omitempty"`   // read-only, xs:datetime
}

// Activities is a group of Activity objects.
type Activities []Activity

// UnmarshalJSON builds the vector from the JSON in b.
func (p *Activities) UnmarshalJSON(b []byte) error {
	var arr *[]Activity
	arr = (*[]Activity)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Activity, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// Activity contains details about activities like museum, theatre, and other events.
type Activity struct {
	ID                   string         `json:"id,omitempty"`                        // optional, read-only
	TripID               string         `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler     bool           `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeURL          string         `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName          string         `json:"display_name,omitempty"`              // optional
	Images               []Image        `json:"Image,omitempty"`                     // optional
	CancellationDateTime DateTime       `json:"CancellationDateTime,omitempty"`      // optional
	BookingDate          string         `json:"booking_date,omitempty"`              // optional, xs:date
	BookingRate          string         `json:"booking_rate,omitempty"`              // optional
	BookingSiteConfNum   string         `json:"booking_site_conf_num,omitempty"`     // optional
	BookingSiteName      string         `json:"booking_site_name,omitempty"`         // optional
	BookingSitePhone     string         `json:"booking_site_phone,omitempty"`        // optional
	BookingSiteURL       string         `json:"booking_site_url,omitempty"`          // optional
	RecordLocator        string         `json:"record_locator,omitempty"`            // optional
	SupplierConfNum      string         `json:"supplier_conf_num,omitempty"`         // optional
	SupplierContact      string         `json:"supplier_contact,omitempty"`          // optional
	SupplierEmailAddress string         `json:"supplier_email_address,omitempty"`    // optional
	SupplierName         string         `json:"supplier_name,omitempty"`             // optional
	SupplierPhone        string         `json:"supplier_phone,omitempty"`            // optional
	SupplierURL          string         `json:"supplier_url,omitempty"`              // optional
	IsPurchased          bool           `json:"is_purchased,string,omitempty"`       // optional
	Notes                string         `json:"notes,omitempty"`                     // optional
	Restrictions         string         `json:"restrictions,omitempty"`              // optional
	TotalCost            string         `json:"total_cost,omitempty"`                // optional
	StartDateTime        DateTime       `json:"StartDateTime,omitempty"`             // optional
	EndTime              string         `json:"end_time,omitempty"`                  // optional, xs:time
	Address              Address        `json:"Address,omitempty"`                   // optional
	Participants         Travelers      `json:"Participant,omitempty"`               // optional
	DetailTypeCode       DetailTypeCode `json:"detail_type_code,omitempty"`          // optional
	LocationName         string         `json:"location_name,omitempty"`             // optional
}

// Cars is a group of Car objects.
type Cars []Car

// UnmarshalJSON builds the vector from the JSON in b.
func (p *Cars) UnmarshalJSON(b []byte) error {
	var arr *[]Car
	arr = (*[]Car)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Car, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// Car contains information about rental cars. car cancellation remarks should be in restrictions. car pickup instructions should be in notes. car daily rate should be in booking_rate.
type Car struct {
	ID                   string    `json:"id,omitempty" xml:"id"`                                         // optional, read-only
	TripID               string    `json:"trip_id,omitempty" xml:"trip_id"`                               // optional
	IsClientTraveler     bool      `json:"is_client_traveler,string,omitempty" xml:"is_client_traveler"`  // optional, read-only
	RelativeURL          string    `json:"relative_url,omitempty" xml:"relative_url"`                     // optional, read-only
	DisplayName          string    `json:"display_name,omitempty" xml:"display_name"`                     // optional
	Images               []Image   `json:"Image,omitempty" xml:"Image"`                                   // optional
	CancellationDateTime DateTime  `json:"CancellationDateTime,omitempty" xml:"CancellationDateTime"`     // optional
	BookingDate          string    `json:"booking_date,omitempty" xml:"booking_date"`                     // optional, xs:date
	BookingRate          string    `json:"booking_rate,omitempty" xml:"booking_rate"`                     // optional
	BookingSiteConfNum   string    `json:"booking_site_conf_num,omitempty" xml:"booking_site_conf_num"`   // optional
	BookingSiteName      string    `json:"booking_site_name,omitempty" xml:"booking_site_name"`           // optional
	BookingSitePhone     string    `json:"booking_site_phone,omitempty" xml:"booking_site_phone"`         // optional
	BookingSiteURL       string    `json:"booking_site_url,omitempty" xml:"booking_site_url"`             // optional
	RecordLocator        string    `json:"record_locator,omitempty" xml:"record_locator"`                 // optional
	SupplierConfNum      string    `json:"supplier_conf_num,omitempty" xml:"supplier_conf_num"`           // optional
	SupplierContact      string    `json:"supplier_contact,omitempty" xml:"supplier_contact"`             // optional
	SupplierEmailAddress string    `json:"supplier_email_address,omitempty" xml:"supplier_email_address"` // optional
	SupplierName         string    `json:"supplier_name,omitempty" xml:"supplier_name"`                   // optional
	SupplierPhone        string    `json:"supplier_phone,omitempty" xml:"supplier_phone"`                 // optional
	SupplierURL          string    `json:"supplier_url,omitempty" xml:"supplier_url"`                     // optional
	IsPurchased          bool      `json:"is_purchased,string,omitempty" xml:"is_purchased"`              // optional
	Notes                string    `json:"notes,omitempty" xml:"notes"`                                   // optional
	Restrictions         string    `json:"restrictions,omitempty" xml:"restrictions"`                     // optional
	TotalCost            string    `json:"total_cost,omitempty" xml:"total_cost"`                         // optional
	StartDateTime        DateTime  `json:"StartDateTime,omitempty" xml:"StartDateTime"`                   // optional
	EndDateTime          DateTime  `json:"EndDateTime,omitempty" xml:"EndDateTime"`                       // optional
	StartLocationAddress Address   `json:"StartLocationAddress,omitempty" xml:"StartLocationAddress"`     // optional
	EndLocationAddress   Address   `json:"EndLocationAddress,omitempty" xml:"EndLocationAddress"`         // optional
	Drivers              Travelers `json:"Driver,omitempty" xml:"Driver"`                                 // optional
	StartLocationHours   string    `json:"start_location_hours,omitempty" xml:"start_location_hours"`     // optional
	StartLocationName    string    `json:"start_location_name,omitempty" xml:"start_location_name"`       // optional
	StartLocationPhone   string    `json:"start_location_phone,omitempty" xml:"start_location_phone"`     // optional
	EndLocationHours     string    `json:"end_location_hours,omitempty" xml:"end_location_hours"`         // optional
	EndLocationName      string    `json:"end_location_name,omitempty" xml:"end_location_name"`           // optional
	EndLocationPhone     string    `json:"end_location_phone,omitempty" xml:"end_location_phone"`         // optional
	CarDescription       string    `json:"car_description,omitempty" xml:"car_description"`               // optional
	CarType              string    `json:"car_type,omitempty" xml:"car_type"`                             // optional
	MileageCharges       string    `json:"mileage_charges,omitempty" xml:"mileage_charges"`               // optional
}

// Cruises is a group of Cruise objects.
type Cruises []Cruise

// UnmarshalJSON builds the vector from the JSON in b.
func (p *Cruises) UnmarshalJSON(b []byte) error {
	var arr *[]Cruise
	arr = (*[]Cruise)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Cruise, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// Cruise contains information about cruises.
type Cruise struct {
	ID                   string          `json:"id,omitempty"`                        // optional, read-only
	TripID               string          `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler     bool            `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeURL          string          `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName          string          `json:"display_name,omitempty"`              // optional
	Images               []Image         `json:"Image,omitempty"`                     // optional
	CancellationDateTime DateTime        `json:"CancellationDateTime,omitempty"`      // optional
	BookingDate          string          `json:"booking_date,omitempty"`              // optional, xs:date
	BookingRate          string          `json:"booking_rate,omitempty"`              // optional
	BookingSiteConfNum   string          `json:"booking_site_conf_num,omitempty"`     // optional
	BookingSiteName      string          `json:"booking_site_name,omitempty"`         // optional
	BookingSitePhone     string          `json:"booking_site_phone,omitempty"`        // optional
	BookingSiteURL       string          `json:"booking_site_url,omitempty"`          // optional
	RecordLocator        string          `json:"record_locator,omitempty"`            // optional
	SupplierConfNum      string          `json:"supplier_conf_num,omitempty"`         // optional
	SupplierContact      string          `json:"supplier_contact,omitempty"`          // optional
	SupplierEmailAddress string          `json:"supplier_email_address,omitempty"`    // optional
	SupplierName         string          `json:"supplier_name,omitempty"`             // optional
	SupplierPhone        string          `json:"supplier_phone,omitempty"`            // optional
	SupplierURL          string          `json:"supplier_url,omitempty"`              // optional
	IsPurchased          bool            `json:"is_purchased,string,omitempty"`       // optional
	Notes                string          `json:"notes,omitempty"`                     // optional
	Restrictions         string          `json:"restrictions,omitempty"`              // optional
	TotalCost            string          `json:"total_cost,omitempty"`                // optional
	Segments             []CruiseSegment `json:"Segment,omitempty"`
	Travelers            Travelers       `json:"Traveler,omitempty"`     // optional
	CabinNumber          string          `json:"cabin_number,omitempty"` // optional
	CabinType            string          `json:"cabin_type,omitempty"`   // optional
	Dining               string          `json:"dining,omitempty"`       // optional
	ShipName             string          `json:"ship_name,omitempty"`    // optional
}

// CruiseSegment contains details about indivual cruise segments.
type CruiseSegment struct {
	StartDateTime   DateTime       `json:"StartDateTime,omitempty"`    // optional
	EndDateTime     DateTime       `json:"EndDateTime,omitempty"`      // optional
	LocationAddress Address        `json:"LocationAddress,omitempty"`  // optional
	LocationName    string         `json:"location_name,omitempty"`    // optional
	DetailTypeCode  DetailTypeCode `json:"detail_type_code,omitempty"` // optional
	ID              string         `json:"id,omitempty"`               // optional, read-only
}

// Directions is a group of Direction objects.
type Directions []Direction

// UnmarshalJSON builds the vector from the JSON in b.
func (p *Directions) UnmarshalJSON(b []byte) error {
	var arr *[]Direction
	arr = (*[]Direction)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Direction, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// Direction contains addresses to show directions for on the trip.
type Direction struct {
	ID               string   `json:"id,omitempty"`                        // optional, read-only
	TripID           string   `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler bool     `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeURL      string   `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName      string   `json:"display_name,omitempty"`              // optional
	Images           []Image  `json:"Image,omitempty"`                     // optional
	DateTime         DateTime `json:"DateTime,omitempty"`                  // optional
	StartAddress     Address  `json:"StartAddress,omitempty"`              // optional
	EndAddress       Address  `json:"EndAddress,omitempty"`                // optional
}

// Flights is a group of Flight objects.
type Flights []Flight

// UnmarshalJSON builds the vector from the JSON in b.
func (p *Flights) UnmarshalJSON(b []byte) error {
	var arr *[]Flight
	arr = (*[]Flight)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Flight, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// Flight contains data about a flight.
type Flight struct {
	ID                   string         `json:"id,omitempty" xml:"id"`                                         // optional, read-only
	TripID               string         `json:"trip_id,omitempty" xml:"trip_id"`                               // optional
	IsClientTraveler     bool           `json:"is_client_traveler,string,omitempty" xml:"is_client_traveler"`  // optional, read-only
	RelativeURL          string         `json:"relative_url,omitempty" xml:"relative_url"`                     // optional, read-only
	DisplayName          string         `json:"display_name,omitempty" xml:"display_name"`                     // optional
	Images               []Image        `json:"Image,omitempty" xml:"Image"`                                   // optional
	CancellationDateTime DateTime       `json:"CancellationDateTime,omitempty" xml:"CancellationDateTime"`     // optional
	BookingDate          string         `json:"booking_date,omitempty" xml:"booking_date"`                     // optional, xs:date
	BookingRate          string         `json:"booking_rate,omitempty" xml:"booking_rate"`                     // optional
	BookingSiteConfNum   string         `json:"booking_site_conf_num,omitempty" xml:"booking_site_conf_num"`   // optional
	BookingSiteName      string         `json:"booking_site_name,omitempty" xml:"booking_site_name"`           // optional
	BookingSitePhone     string         `json:"booking_site_phone,omitempty" xml:"booking_site_phone"`         // optional
	BookingSiteURL       string         `json:"booking_site_url,omitempty" xml:"booking_site_url"`             // optional
	RecordLocator        string         `json:"record_locator,omitempty" xml:"record_locator"`                 // optional
	SupplierConfNum      string         `json:"supplier_conf_num,omitempty" xml:"supplier_conf_num"`           // optional
	SupplierContact      string         `json:"supplier_contact,omitempty" xml:"supplier_contact"`             // optional
	SupplierEmailAddress string         `json:"supplier_email_address,omitempty" xml:"supplier_email_address"` // optional
	SupplierName         string         `json:"supplier_name,omitempty" xml:"supplier_name"`                   // optional
	SupplierPhone        string         `json:"supplier_phone,omitempty" xml:"supplier_phone"`                 // optional
	SupplierURL          string         `json:"supplier_url,omitempty" xml:"supplier_url"`                     // optional
	IsPurchased          bool           `json:"is_purchased,string,omitempty" xml:"is_purchased"`              // optional
	Notes                string         `json:"notes,omitempty" xml:"notes"`                                   // optional
	Restrictions         string         `json:"restrictions,omitempty" xml:"restrictions"`                     // optional
	TotalCost            string         `json:"total_cost,omitempty" xml:"total_cost"`                         // optional
	Segments             FlightSegments `json:"Segment,omitempty" xml:"Segment"`
	Travelers            Travelers      `json:"Traveler,omitempty" xml:"Traveler"` // optional
}

// FlightSegments is a group of FlightSegment objects.
type FlightSegments []FlightSegment

// UnmarshalJSON builds the vector from the JSON in b.
func (p *FlightSegments) UnmarshalJSON(b []byte) error {
	var arr *[]FlightSegment
	arr = (*[]FlightSegment)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]FlightSegment, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// FlightSegment contains details about individual flights.
type FlightSegment struct {
	ID                    string       `json:"id,omitempty" xml:"id"`                                                  // optional, read-only
	Status                FlightStatus `json:"Status,omitempty" xml:"Status"`                                          // optional
	StartDateTime         DateTime     `json:"StartDateTime,omitempty" xml:"StartDateTime"`                            // optional
	EndDateTime           DateTime     `json:"EndDateTime,omitempty" xml:"EndDateTime"`                                // optional
	StartAirportCode      string       `json:"start_airport_code,omitempty" xml:"start_airport_code"`                  // optional
	StartAirportLatitude  float64      `json:"start_airport_latitude,string,omitempty" xml:"start_airport_latitude"`   // optional, read-only
	StartAirportLongitude float64      `json:"start_airport_longitude,string,omitempty" xml:"start_airport_longitude"` // optional, read-only
	StartCityName         string       `json:"start_city_name,omitempty" xml:"start_city_name"`                        // optional
	StartGate             string       `json:"start_gate,omitempty" xml:"start_gate"`                                  // optional
	StartTerminal         string       `json:"start_terminal,omitempty" xml:"start_terminal"`                          // optional
	EndAirportCode        string       `json:"end_airport_code,omitempty" xml:"end_airport_code"`                      // optional
	EndAirportLatitude    float64      `json:"end_airport_latitude,string,omitempty" xml:"end_airport_latitude"`       // optional, read-only
	EndAirportLongitude   float64      `json:"end_airport_longitude,string,omitempty" xml:"end_airport_longitude"`     // optional, read-only
	EndCityName           string       `json:"end_city_name,omitempty" xml:"end_city_name"`                            // optional
	EndGate               string       `json:"end_gate,omitempty" xml:"end_gate"`                                      // optional
	EndTerminal           string       `json:"end_terminal,omitempty" xml:"end_terminal"`                              // optional
	MarketingAirline      string       `json:"marketing_airline,omitempty" xml:"marketing_airline"`                    // optional
	MarketingAirlineCode  string       `json:"marketing_airline_code,omitempty" xml:"marketing_airline_code"`          // optional, read-only
	MarketingFlightNumber string       `json:"marketing_flight_number,omitempty" xml:"marketing_flight_number"`        // optional
	OperatingAirline      string       `json:"operating_airline,omitempty" xml:"operating_airline"`                    // optional
	OperatingAirlineCode  string       `json:"operating_airline_code,omitempty" xml:"operating_airline_code"`          // optional, read-only
	OperatingFlightNumber string       `json:"operating_flight_number,omitempty" xml:"operating_flight_number"`        // optional
	AlternativeFlightsURL string       `json:"alternate_flights_url,omitempty" xml:"alternate_flights_url"`            // optional, read-only
	Aircraft              string       `json:"aircraft,omitempty" xml:"aircraft"`                                      // optional
	AircraftDisplayName   string       `json:"aircraft_display_name,omitempty" xml:"aircraft_display_name"`            // optional, read-only
	Distance              string       `json:"distance,omitempty" xml:"distance"`                                      // optional
	Duration              string       `json:"duration,omitempty" xml:"duration"`                                      // optional
	Entertainment         string       `json:"entertainment,omitempty" xml:"entertainment"`                            // optional
	Meal                  string       `json:"meal,omitempty" xml:"meal"`                                              // optional
	Notes                 string       `json:"notes,omitempty" xml:"notes"`                                            // optional
	OntimePerc            string       `json:"ontime_perc,omitempty" xml:"ontime_perc"`                                // optional
	Seats                 string       `json:"seats,omitempty" xml:"seats"`                                            // optional
	ServiceClass          string       `json:"service_class,omitempty" xml:"service_class"`                            // optional
	Stops                 string       `json:"stops,omitempty" xml:"stops"`                                            // optional
	BaggageClaim          string       `json:"baggage_claim,omitempty" xml:"baggage_claim"`                            // optional
	CheckInURL            string       `json:"check_in_url,omitempty" xml:"check_in_url"`                              // optional
	ConflictResolutionURL string       `json:"conflict_resolution_url,omitempty" xml:"conflict_resolution_url"`        // optional, read-only
	IsHidden              bool         `json:"is_hidden,string,omitempty" xml:"is_hidden"`                             // optional, read-only
}

// FlightStatus fields are read-only and only available for monitored TripIt Pro AirSegments.
type FlightStatus struct {
	ScheduledDepartureDateTime DateTime         `json:"ScheduledDepartureDateTime,omitempty" xml:"ScheduledDepartureDateTime"` // optional, read-only
	EstimatedDepartureDateTime DateTime         `json:"EstimatedDepartureDateTime,omitempty" xml:"EstimatedDepartureDateTime"` // optional, read-only
	ScheduledArrivalDateTime   DateTime         `json:"ScheduledArrivalDateTime,omitempty" xml:"ScheduledArrivalDateTime"`     // optional, read-only
	EstimatedArrivalDateTime   DateTime         `json:"EstimatedArrivalDateTime,omitempty" xml:"EstimatedArrivalDateTime"`     // optional, read-only
	FlightStatus               FlightStatusCode `json:"flight_status,string,omitempty" xml:"flight_status"`                    // optional, read-only
	IsConnectionAtRisk         bool             `json:"is_connection_at_risk,string,omitempty" xml:"is_connection_at_risk"`    // optional, read-only
	DepartureTerminal          string           `json:"departure_terminal,omitempty" xml:"departure_terminal"`                 // optional, read-only
	DepartureGate              string           `json:"departure_gate,omitempty" xml:"departure_gate"`                         // optional, read-only
	ArrivalTerminal            string           `json:"arrival_terminal,omitempty" xml:"arrival_terminal"`                     // optional, read-only
	ArrivalGate                string           `json:"arrival_gate,omitempty" xml:"arrival_gate"`                             // optional, read-only
	LayoverMinutes             string           `json:"layover_minutes,omitempty" xml:"layover_minutes"`                       // optional, read-only
	BaggageClaim               string           `json:"baggage_claim,omitempty" xml:"baggage_claim"`                           // optional, read-only
	DivertedAirportCode        string           `json:"diverted_airport_code,omitempty" xml:"diverted_airport_code"`           // optional, read-only
	LastModified               string           `json:"last_modified,omitempty" xml:"last_modified"`                           // read-only
}

// Lodges is a group of Lodging objects.
type Lodges []Lodging

// UnmarshalJSON builds the vector from the JSON in b.
func (p *Lodges) UnmarshalJSON(b []byte) error {
	var arr *[]Lodging
	arr = (*[]Lodging)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Lodging, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// Lodging contains information about hotels or other lodging. hotel cancellation remarks should be in restrictions. hotel room description should be in notes. hotel average daily rate should be in booking_rate.
type Lodging struct {
	ID                   string    `json:"id,omitempty" xml:"id"`                                         // optional, read-only
	TripID               string    `json:"trip_id,omitempty" xml:"trip_id"`                               // optional
	IsClientTraveler     bool      `json:"is_client_traveler,string,omitempty" xml:"is_client_traveler"`  // optional, read-only
	RelativeURL          string    `json:"relative_url,omitempty" xml:"relative_url"`                     // optional, read-only
	DisplayName          string    `json:"display_name,omitempty" xml:"display_name"`                     // optional
	Images               []Image   `json:"Image,omitempty" xml:"Image"`                                   // optional
	CancellationDateTime DateTime  `json:"CancellationDateTime,omitempty" xml:"CancellationDateTime"`     // optional
	BookingDate          string    `json:"booking_date,omitempty" xml:"booking_date"`                     // optional, xs:date
	BookingRate          string    `json:"booking_rate,omitempty" xml:"booking_rate"`                     // optional
	BookingSiteConfNum   string    `json:"booking_site_conf_num,omitempty" xml:"booking_site_conf_num"`   // optional
	BookingSiteName      string    `json:"booking_site_name,omitempty" xml:"booking_site_name"`           // optional
	BookingSitePhone     string    `json:"booking_site_phone,omitempty" xml:"booking_site_phone"`         // optional
	BookingSiteURL       string    `json:"booking_site_url,omitempty" xml:"booking_site_url"`             // optional
	RecordLocator        string    `json:"record_locator,omitempty" xml:"record_locator"`                 // optional
	SupplierConfNum      string    `json:"supplier_conf_num,omitempty" xml:"supplier_conf_num"`           // optional
	SupplierContact      string    `json:"supplier_contact,omitempty" xml:"supplier_contact"`             // optional
	SupplierEmailAddress string    `json:"supplier_email_address,omitempty" xml:"supplier_email_address"` // optional
	SupplierName         string    `json:"supplier_name,omitempty" xml:"supplier_name"`                   // optional
	SupplierPhone        string    `json:"supplier_phone,omitempty" xml:"supplier_phone"`                 // optional
	SupplierURL          string    `json:"supplier_url,omitempty" xml:"supplier_url"`                     // optional
	IsPurchased          bool      `json:"is_purchased,string,omitempty" xml:"is_purchased"`              // optional
	Notes                string    `json:"notes,omitempty" xml:"notes"`                                   // optional
	Restrictions         string    `json:"restrictions,omitempty" xml:"restrictions"`                     // optional
	TotalCost            string    `json:"total_cost,omitempty" xml:"total_cost"`                         // optional
	StartDateTime        DateTime  `json:"StartDateTime,omitempty" xml:"StartDateTime"`                   // optional
	EndDateTime          DateTime  `json:"EndDateTime,omitempty" xml:"EndDateTime"`                       // optional
	Address              Address   `json:"Address,omitempty" xml:"Address"`                               // optional
	Guests               Travelers `json:"Guest,omitempty" xml:"Guest"`                                   // optional
	NumberGuests         string    `json:"number_guests,omitempty" xml:"number_guests"`                   // optional
	NumberRooms          string    `json:"number_rooms,omitempty" xml:"number_rooms"`                     // optional
	RoomType             string    `json:"room_type,omitempty" xml:"room_type"`                           // optional
}

// Maps is a group of Map objects.
type Maps []Map

// UnmarshalJSON builds the vector from the JSON in b.
func (p *Maps) UnmarshalJSON(b []byte) error {
	var arr *[]Map
	arr = (*[]Map)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Map, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// Map contains addresses to show on a map.
type Map struct {
	ID               string   `json:"id,omitempty"`                        // optional, read-only
	TripID           string   `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler bool     `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeURL      string   `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName      string   `json:"display_name,omitempty"`              // optional
	Images           []Image  `json:"Image,omitempty"`                     // optional
	DateTime         DateTime `json:"DateTime,omitempty"`                  // optional
	Address          Address  `json:"Address,omitempty"`                   // optional
}

// Notes is a group of Note objects.
type Notes []Note

// UnmarshalJSON builds the vector from the JSON in b.
func (p *Notes) UnmarshalJSON(b []byte) error {
	var arr *[]Note
	arr = (*[]Note)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Note, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// Note contains information about notes added by the traveler.
type Note struct {
	ID               string         `json:"id,omitempty"`                        // optional, read-only
	TripID           string         `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler bool           `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeURL      string         `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName      string         `json:"display_name,omitempty"`              // optional
	Images           []Image        `json:"Image,omitempty"`                     // optional
	DateTime         DateTime       `json:"DateTime,omitempty"`                  // optional
	Address          Address        `json:"Address,omitempty"`                   // optional
	DetailTypeCode   DetailTypeCode `json:"detail_type_code,omitempty"`          // optional
	Source           string         `json:"source,omitempty"`                    // optional
	Text             string         `json:"text,omitempty"`                      // optional
	URL              string         `json:"url,omitempty"`                       // optional
	Notes            string         `json:"notes,omitempty"`                     // optional
}

// Rails is a group of Rail objects.
type Rails []Rail

// UnmarshalJSON builds the vector from the JSON in b.
func (p *Rails) UnmarshalJSON(b []byte) error {
	var arr *[]Rail
	arr = (*[]Rail)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Rail, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// Rail contains information about trains.
type Rail struct {
	ID                   string       `json:"id,omitempty"`                        // optional, read-only
	TripID               string       `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler     bool         `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeURL          string       `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName          string       `json:"display_name,omitempty"`              // optional
	Images               []Image      `json:"Image,omitempty"`                     // optional
	CancellationDateTime DateTime     `json:"CancellationDateTime,omitempty"`      // optional
	BookingDate          string       `json:"booking_date,omitempty"`              // optional, xs:date
	BookingRate          string       `json:"booking_rate,omitempty"`              // optional
	BookingSiteConfNum   string       `json:"booking_site_conf_num,omitempty"`     // optional
	BookingSiteName      string       `json:"booking_site_name,omitempty"`         // optional
	BookingSitePhone     string       `json:"booking_site_phone,omitempty"`        // optional
	BookingSiteURL       string       `json:"booking_site_url,omitempty"`          // optional
	RecordLocator        string       `json:"record_locator,omitempty"`            // optional
	SupplierConfNum      string       `json:"supplier_conf_num,omitempty"`         // optional
	SupplierContact      string       `json:"supplier_contact,omitempty"`          // optional
	SupplierEmailAddress string       `json:"supplier_email_address,omitempty"`    // optional
	SupplierName         string       `json:"supplier_name,omitempty"`             // optional
	SupplierPhone        string       `json:"supplier_phone,omitempty"`            // optional
	SupplierURL          string       `json:"supplier_url,omitempty"`              // optional
	IsPurchased          bool         `json:"is_purchased,string,omitempty"`       // optional
	Notes                string       `json:"notes,omitempty"`                     // optional
	Restrictions         string       `json:"restrictions,omitempty"`              // optional
	TotalCost            string       `json:"total_cost,omitempty"`                // optional
	Segments             RailSegments `json:"Segment,omitempty"`
	Travelers            Travelers    `json:"Traveler,omitempty"` // optional
}

// RailSegments is a group of RailSegment objects.
type RailSegments []RailSegment

// UnmarshalJSON builds the vector from the JSON in b.
func (p *RailSegments) UnmarshalJSON(b []byte) error {
	var arr *[]RailSegment
	arr = (*[]RailSegment)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]RailSegment, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// RailSegment contains details about an individual train ride.
type RailSegment struct {
	ID                  string   `json:"id,omitempty"`                  // optional, read-only
	StartDateTime       DateTime `json:"StartDateTime,omitempty"`       // optional
	EndDateTime         DateTime `json:"EndDateTime,omitempty"`         // optional
	StartStationAddress Address  `json:"StartStationAddress,omitempty"` // optional
	EndStationAddress   Address  `json:"EndStationAddress,omitempty"`   // optional
	StartStationName    string   `json:"start_station_name,omitempty"`  // optional
	EndStationName      string   `json:"end_station_name,omitempty"`    // optional
	CarrierName         string   `json:"carrier_name,omitempty"`        // optional
	CoachNumber         string   `json:"coach_number,omitempty"`        // optional
	ConfirmationNum     string   `json:"confirmation_num,omitempty"`    // optional
	Seats               string   `json:"seats,omitempty"`               // optional
	ServiceClass        string   `json:"service_class,omitempty"`       // optional
	TrainNumber         string   `json:"train_number,omitempty"`        // optional
	TrainType           string   `json:"train_type,omitempty"`          // optional
}

// Restaurants is a group of Restaurant objects.
type Restaurants []Restaurant

// UnmarshalJSON builds the vector from the JSON in b.
func (p *Restaurants) UnmarshalJSON(b []byte) error {
	var arr *[]Restaurant
	arr = (*[]Restaurant)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Restaurant, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// Restaurant contains details about dining reservations. restaurant name should be in supplier_name. restaurant notes should be in notes.
type Restaurant struct {
	ID                   string   `json:"id,omitempty"`                        // optional, read-only
	TripID               string   `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler     bool     `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeURL          string   `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName          string   `json:"display_name,omitempty"`              // optional
	Images               []Image  `json:"Image,omitempty"`                     // optional
	CancellationDateTime DateTime `json:"CancellationDateTime,omitempty"`      // optional
	BookingDate          string   `json:"booking_date,omitempty"`              // optional, xs:date
	BookingRate          string   `json:"booking_rate,omitempty"`              // optional
	BookingSiteConfNum   string   `json:"booking_site_conf_num,omitempty"`     // optional
	BookingSiteName      string   `json:"booking_site_name,omitempty"`         // optional
	BookingSitePhone     string   `json:"booking_site_phone,omitempty"`        // optional
	BookingSiteURL       string   `json:"booking_site_url,omitempty"`          // optional
	RecordLocator        string   `json:"record_locator,omitempty"`            // optional
	SupplierConfNum      string   `json:"supplier_conf_num,omitempty"`         // optional
	SupplierContact      string   `json:"supplier_contact,omitempty"`          // optional
	SupplierEmailAddress string   `json:"supplier_email_address,omitempty"`    // optional
	SupplierName         string   `json:"supplier_name,omitempty"`             // optional
	SupplierPhone        string   `json:"supplier_phone,omitempty"`            // optional
	SupplierURL          string   `json:"supplier_url,omitempty"`              // optional
	IsPurchased          bool     `json:"is_purchased,string,omitempty"`       // optional
	Notes                string   `json:"notes,omitempty"`                     // optional
	Restrictions         string   `json:"restrictions,omitempty"`              // optional
	TotalCost            string   `json:"total_cost,omitempty"`                // optional
	DateTime             DateTime `json:"DateTime,omitempty"`                  // optional
	Address              Address  `json:"Address,omitempty"`                   // optional
	ReservationHolder    Traveler `json:"ReservationHolder,omitempty"`         // optional
	Cuisine              string   `json:"cuisine,omitempty"`                   // optional
	DressCode            string   `json:"dress_code,omitempty"`                // optional
	Hours                string   `json:"hours,omitempty"`                     // optional
	NumberPatrons        string   `json:"number_patrons,omitempty"`            // optional
	PriceRange           string   `json:"price_range,omitempty"`               // optional
}

// Transports is a group of Transport objects.
type Transports []Transport

// UnmarshalJSON builds the vector from the JSON in b.
func (p *Transports) UnmarshalJSON(b []byte) error {
	var arr *[]Transport
	arr = (*[]Transport)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Transport, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// Transport contains details about other forms of transport like bus rides.
type Transport struct {
	ID                   string            `json:"id,omitempty"`                        // optional, read-only
	TripID               string            `json:"trip_id,omitempty"`                   // optional
	IsClientTraveler     bool              `json:"is_client_traveler,string,omitempty"` // optional, read-only
	RelativeURL          string            `json:"relative_url,omitempty"`              // optional, read-only
	DisplayName          string            `json:"display_name,omitempty"`              // optional
	Images               []Image           `json:"Image,omitempty"`                     // optional
	CancellationDateTime DateTime          `json:"CancellationDateTime,omitempty"`      // optional
	BookingDate          string            `json:"booking_date,omitempty"`              // optional, xs:date
	BookingRate          string            `json:"booking_rate,omitempty"`              // optional
	BookingSiteConfNum   string            `json:"booking_site_conf_num,omitempty"`     // optional
	BookingSiteName      string            `json:"booking_site_name,omitempty"`         // optional
	BookingSitePhone     string            `json:"booking_site_phone,omitempty"`        // optional
	BookingSiteURL       string            `json:"booking_site_url,omitempty"`          // optional
	RecordLocator        string            `json:"record_locator,omitempty"`            // optional
	SupplierConfNum      string            `json:"supplier_conf_num,omitempty"`         // optional
	SupplierContact      string            `json:"supplier_contact,omitempty"`          // optional
	SupplierEmailAddress string            `json:"supplier_email_address,omitempty"`    // optional
	SupplierName         string            `json:"supplier_name,omitempty"`             // optional
	SupplierPhone        string            `json:"supplier_phone,omitempty"`            // optional
	SupplierURL          string            `json:"supplier_url,omitempty"`              // optional
	IsPurchased          bool              `json:"is_purchased,string,omitempty"`       // optional
	Notes                string            `json:"notes,omitempty"`                     // optional
	Restrictions         string            `json:"restrictions,omitempty"`              // optional
	TotalCost            string            `json:"total_cost,omitempty"`                // optional
	Segments             TransportSegments `json:"Segment,omitempty"`
	Travelers            Travelers         `json:"Traveler,omitempty"` // optional
}

// TransportSegments is a group of TransportSegment objects.
type TransportSegments []TransportSegment

// UnmarshalJSON builds the vector from the JSON in b.
func (p *TransportSegments) UnmarshalJSON(b []byte) error {
	var arr *[]TransportSegment
	arr = (*[]TransportSegment)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]TransportSegment, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// TransportSegment contains details about indivual transport rides.
type TransportSegment struct {
	ID                   string         `json:"id,omitempty"`                   // optional, read-only
	StartDateTime        DateTime       `json:"StartDateTime,omitempty"`        // optional
	EndDateTime          DateTime       `json:"EndDateTime,omitempty"`          // optional
	StartLocationAddress Address        `json:"StartLocationAddress,omitempty"` // optional
	EndLocationAddress   Address        `json:"EndLocationAddress,omitempty"`   // optional
	StartLocationName    string         `json:"start_location_name,omitempty"`  // optional
	EndLocationName      string         `json:"end_location_name,omitempty"`    // optional
	DetailTypeCode       DetailTypeCode `json:"detail_type_code,omitempty"`     // optional
	CarrierName          string         `json:"carrier_name,omitempty"`         // optional
	ConfirmationNum      string         `json:"confirmation_num,omitempty"`     // optional
	NumberPassengers     string         `json:"number_passengers,omitempty"`    // optional
	VehicleDescription   string         `json:"vehicle_description,omitempty"`  // optional
}

// Trips is a group of Trip objects.
type Trips []Trip

// UnmarshalJSON builds the vector from the JSON in b.
func (p *Trips) UnmarshalJSON(b []byte) error {
	var arr *[]Trip
	arr = (*[]Trip)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Trip, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// Trip represents a trip in the TripIt model.
type Trip struct {
	ID                     string           `json:"id,omitempty" xml:"id"`                                             // optional, id is a read-only field
	RelativeURL            string           `json:"relative_url,omitempty" xml:"relative_url"`                         // optional, relative_url is a read-only field
	StartDate              string           `json:"start_date,omitempty" xml:"start_date"`                             // optional, xs:date
	EndDate                string           `json:"end_date,omitempty" xml:"end_date"`                                 // optional, xs:date
	Description            string           `json:"description,omitempty" xml:"description"`                           // optional
	DisplayName            string           `json:"display_name,omitempty" xml:"display_name"`                         // optional
	ImageURL               string           `json:"image_url,omitempty" xml:"image_url"`                               // optional
	IsPrivate              bool             `json:"is_private,string,omitempty" xml:"is_private"`                      // optional
	PrimaryLocation        string           `json:"primary_location,omitempty" xml:"primary_location"`                 // optional
	PrimaryLocationAddress Address          `json:"primary_location_address,omitempty" xml:"primary_location_address"` // optional, PrimaryLocationAddress is a read-only field
	ClosenessMatches       ClosenessMatches `json:"ClosenessMatches,omitempty" xml:"ClosenessMatches"`                 // optional, ClosenessMatches are read-only
	Invitees               Invitees         `json:"TripInvitees,omitempty" xml:"TripInvitees"`                         // optional, Invitees are read-only
	Remarks                Remarks          `json:"TripCrsRemarks,omitempty" xml:"TripCrsRemarks"`                     // optional, Remarks are read-only
}

// Weather contains information about the weather at a particular destination. Weather is read-only.
type Weather struct {
	ID                 string  `json:"id,omitempty"`                          // optional, read-only
	TripID             string  `json:"trip_id,omitempty"`                     // optional
	IsClientTraveler   bool    `json:"is_client_traveler,string,omitempty"`   // optional, read-only
	RelativeURL        string  `json:"relative_url,omitempty"`                // optional, read-only
	DisplayName        string  `json:"display_name,omitempty"`                // optional
	Images             []Image `json:"Image,omitempty"`                       // optional
	Date               string  `json:"date,omitempty"`                        // optional, read-only, xs:date
	Location           string  `json:"location,omitempty"`                    // optional, read-only
	AvgHighTempC       float64 `json:"avg_high_temp_c,string,omitempty"`      // optional, read-only
	AvgLowTempC        float64 `json:"avg_low_temp_c,string,omitempty"`       // optional, read-only
	AvgWindSpeedKn     float64 `json:"avg_wind_speed_kn,string,omitempty"`    // optional, read-only
	AvgPrecipitationCm float64 `json:"avg_precipitation_cm,string,omitempty"` // optional, read-only
	AvgSnowDepthCm     float64 `json:"avg_snow_depth_cm,string,omitempty"`    // optional, read-only
}

// PointsProgram contains information about tracked travel programs for TripIt Pro users. All PointsProgram elements are read-only.
type PointsProgram struct {
	ID                  uint                      `json:"id,string,omitempty" xml:"id"`                                       // read-only
	Name                string                    `json:"name,omitempty" xml:"name"`                                          // optional, read-only
	AccountNumber       string                    `json:"account_number,omitempty" xml:"account_number"`                      // optional, read-only
	AccountLogin        string                    `json:"account_login,omitempty" xml:"account_login"`                        // optional, read-only
	Balance             string                    `json:"balance,omitempty" xml:"balance"`                                    // optional, read-only
	EliteStatus         string                    `json:"elite_status,omitempty" xml:"elite_status"`                          // optional, read-only
	EliteNextStatus     string                    `json:"elite_next_status,omitempty" xml:"elite_next_status"`                // optional, read-only
	EliteYtdQualify     string                    `json:"elite_ytd_qualify,omitempty" xml:"elite_ytd_qualify"`                // optional, read-only
	EliteNeedToEarn     string                    `json:"elite_need_to_earn,omitempty" xml:"elite_need_to_earn"`              // optional, read-only
	LastModified        string                    `json:"last_modified,omitempty" xml:"last_modified"`                        // read-only
	TotalNumActivities  int                       `json:"total_num_activities,string,omitempty" xml:"total_num_activities"`   // read-only
	TotalNumExpirations int                       `json:"total_num_expirations,string,omitempty" xml:"total_num_expirations"` // read-only
	ErrorMessage        string                    `json:"error_message,omitempty" xml:"error_message"`                        // optional, read-only
	Activities          []PointsProgramActivity   `json:"Activity,omitempty" xml:"Activity"`                                  // optional, read-only
	Expirations         []PointsProgramExpiration `json:"Expiration,omitempty" xml:"Expiration"`                              // optional, read-only
}

// PointsProgramActivity contains program transactions All PointsProgramActivity elements are read-only
type PointsProgramActivity struct {
	Date        string `json:"date,omitempty" xml:"date"`               // read-only, xs:date
	Description string `json:"description,omitempty" xml:"description"` // optional, read-only
	Base        string `json:"base,omitempty" xml:"base"`               // optional, read-only
	Bonus       string `json:"bonus,omitempty" xml:"bonus"`             // optional, read-only
	Total       string `json:"total,omitempty" xml:"total"`             // optional, read-only
}

// PointsProgramExpiration elements are read-only.
type PointsProgramExpiration struct {
	Date   string `json:"date,omitempty" xml:"date"`     // read-only, xs:date
	Amount string `json:"amount,omitempty" xml:"amount"` // optional, read-only
}

// Profile contains user information. All Profile elements are read-only.
type Profile struct {
	Attributes            ProfileAttributes     `json:"_attributes" xml:"attributes"`                                // read-only
	ProfileEmailAddresses ProfileEmailAddresses `json:"ProfileEmailAddresses,omitempty" xml:"ProfileEmailAddresses"` // optional, read-only
	GroupMemberships      GroupMemberships      `json:"GroupMemberships,omitempty" xml:"GroupMemberships"`           // optional, read-only
	IsClient              bool                  `json:"is_client,string,omitempty" xml:"is_client"`                  // read-only
	IsPro                 bool                  `json:"is_pro,string,omitempty" xml:"is_pro"`                        // read-only
	ScreenName            string                `json:"screen_name,omitempty" xml:"screen_name"`                     // read-only
	PublicDisplayName     string                `json:"public_display_name,omitempty" xml:"public_display_name"`     // read-only
	ProfileURL            string                `json:"profile_url,omitempty" xml:"profile_url"`                     // read-only
	HomeCity              string                `json:"home_city,omitempty" xml:"home_city"`                         // optional, read-only
	Company               string                `json:"company,omitempty" xml:"company"`                             // optional, read-only
	AboutMeInfo           string                `json:"about_me_info,omitempty" xml:"about_me_info"`                 // optional, read-only
	PhotoURL              string                `json:"photo_url,omitempty" xml:"photo_url"`                         // optional, read-only
	ActivityFeedURL       string                `json:"activity_feed_url,omitempty" xml:"activity_feed_url"`         // optional, read-only
	AlertsFeedURL         string                `json:"alerts_feed_url,omitempty" xml:"alerts_feed_url"`             // optional, read-only
	IcalURL               string                `json:"ical_url,omitempty" xml:"ical_url"`                           // optional, read-only
}

// ProfileAttributes represent links to profiles.
type ProfileAttributes struct {
	Ref string `json:"ref,omitempty" xml:"ref"` // read-only
}

// ProfileEmailAddresses contains the list of email addresses for a user.
type ProfileEmailAddresses struct {
	ProfileEmailAddresses ProfileEmailAddress `json:"ProfileEmailAddress,omitempty" xml:"ProfileEmailAddress"`
}

// ProfileEmailAddress contains an email address and its properties. All ProfileEmailAddress elements are read-only.
type ProfileEmailAddress struct {
	EmailRef     string `json:"email_ref" xml:"email_ref"`                                            // read-only
	Address      string `json:"address" xml:"address"`                                                // read-only
	IsAutoImport bool   `json:"is_auto_import,string,omitempty" xml:"is_auto_import"`                 // read-only
	IsConfirmed  bool   `json:"is_confirmed,string,omitempty" xml:"is_confirmed"`                     // read-only
	IsPrimary    bool   `json:"is_primary,string,omitempty" xml:"is_primary"`                         // read-only
	IsAutoInbox  bool   `json:"is_auto_inbox_eligible,string,omitempty" xml:"is_auto_inbox_eligible"` // read-only
}

// GroupMemberships contains a list of groups that the user is a member of.
type GroupMemberships []Group

// UnmarshalJSON builds the vector from the JSON in b.
func (p *GroupMemberships) UnmarshalJSON(b []byte) error {
	var arr *[]Group
	arr = (*[]Group)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Group, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// Group contains data about a group in TripIt. All Group elements are read-only.
type Group struct {
	DisplayName string `json:"display_name,omitempty" xml:"display_name"` // read-only
	URL         string `json:"url" xml:"url"`                             // read-only
}

// Address represents the address of a location. For create, use either: - address for single-line addresses. - addr1, addr2, city, state, zip, and country for multi-line addresses. Multi-line address will be ignored if single-line address is present. See documentation for more information.
type Address struct {
	Address   string  `json:"address,omitempty" xml:"address"`            // optional
	Addr1     string  `json:"addr1,omitempty" xml:"addr1"`                // optional
	Addr2     string  `json:"addr2,omitempty" xml:"addr2"`                // optional
	City      string  `json:"city,omitempty" xml:"city"`                  // optional
	State     string  `json:"state,omitempty" xml:"state"`                // optional
	Zip       string  `json:"zip,omitempty" xml:"zip"`                    // optional
	Country   string  `json:"country,omitempty" xml:"country"`            // optional
	Latitude  float64 `json:"latitude,string,omitempty" xml:"latitude"`   // optional, read-only
	Longitude float64 `json:"longitude,string,omitempty" xml:"longitude"` // optional, read-only
}

// DateTime stores date and time zone information.
type DateTime struct {
	Date      string `json:"date,omitempty" xml:"date"`             // optional, xs:date
	Time      string `json:"time,omitempty" xml:"time"`             // optional, xs:time
	Timezone  string `json:"timezone,omitempty" xml:"timezone"`     // optional, read-only
	UTCOffset string `json:"utc_offset,omitempty" xml:"utc_offset"` // optional, read-only
}

// Parse converts the DateTime to a time.Time with the respective Date, Time, and Timezone information from the DateTime object.
func (d *DateTime) Parse() (time.Time, error) {
	if d.UTCOffset == "" {
		return time.Parse(time.RFC3339, fmt.Sprintf("%sT%sZ", d.Date, d.Time))
	}
	return time.Parse(time.RFC3339, fmt.Sprintf("%sT%s%s", d.Date, d.Time, d.UTCOffset))
}

// Image stores information about images.
type Image struct {
	Caption string `json:"caption,omitempty" xml:"caption"` // optional
	URL     string `json:"url" xml:"url"`
}

// Travelers is a group of Traveler objects.
type Travelers []Traveler

// UnmarshalJSON builds the vector from the JSON in b.
func (p *Travelers) UnmarshalJSON(b []byte) error {
	var arr *[]Traveler
	arr = (*[]Traveler)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Traveler, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// Traveler contains information about a traveler.
type Traveler struct {
	FirstName                string `json:"first_name,omitempty" xml:"first_name"`                                 // optional
	MiddleName               string `json:"middle_name,omitempty" xml:"middle_name"`                               // optional
	LastName                 string `json:"last_name,omitempty" xml:"last_name"`                                   // optional
	FrequentTravelerNum      string `json:"frequent_traveler_num,omitempty" xml:"frequent_traveler_num"`           // optional
	FrequentTravelerSupplier string `json:"frequent_traveler_supplier,omitempty" xml:"frequent_traveler_supplier"` // optional
	MealPreference           string `json:"meal_preference,omitempty" xml:"meal_preference"`                       // optional
	SeatPreference           string `json:"seat_preference,omitempty" xml:"seat_preference"`                       // optional
	TicketNum                string `json:"ticket_num,omitempty" xml:"ticket_num"`                                 // optional
}

// ClosenessMatches are TripIt users who are near this trip.
type ClosenessMatches struct {
	ClosenessMatches []ClosenessMatch `json:"Match,omitempty" xml:"Match"` // optional, ClosenessMatches are read-only
}

// ClosenessMatch refers to nearby users. All ClosenessMatch elements are read-only.
type ClosenessMatch struct {
	Attributes ClosenessMatchAttributes `json:"_attributes" xml:"attributes"` // read-only, Use the profile_ref attribute to reference a Profile
}

// ClosenessMatchAttributes links to profiles of nearby users.
type ClosenessMatchAttributes struct {
	ProfileRef string `json:"profile_ref" xml:"profile_ref"` // read-only, Use the profile_ref attribute to reference a Profile
}

// Invitees are people invited to view a trip.
type Invitees []Invitee

// UnmarshalJSON builds the vector from the JSON in b.
func (p *Invitees) UnmarshalJSON(b []byte) error {
	var arr *[]Invitee
	arr = (*[]Invitee)(p)
	*arr = nil
	err := json.Unmarshal(b, arr)
	if err != nil {
		*arr = make([]Invitee, 1)
		err := json.Unmarshal(b, &(*arr)[0])
		if err != nil {
			if err2, ok := err.(*json.UnmarshalTypeError); ok && err2.Value == "null" {
				*arr = (*arr)[0:0]
			} else {
				return err
			}
		}

	}
	return nil
}

// Invitee stores attributes about invitees to a trip. All Invitee elements are read-only.
type Invitee struct {
	IsReadOnly bool              `json:"is_read_only,string,omitempty" xml:"is_read_only"` // read-only
	IsTraveler bool              `json:"is_traveler,string,omitempty" xml:"is_traveler"`   // read-only
	Attributes InviteeAttributes `json:"_attributes" xml:"attributes"`                     // read-only, Use the profile_ref attribute to reference a Profile
}

// InviteeAttributes are used to link to user profiles.
type InviteeAttributes struct {
	ProfileRef string `json:"profile_ref" xml:"profile_ref"` // read-only, used to reference a profile
}

// Remarks are remarks from a reservation system.
type Remarks struct {
	Remarks []Remark `json:"TripCrsRemark,omitempty" xml:"TripCrsRemark"` // optional, TripCrsRemarks are read-only
}

// Remark is a reservation system remark. All TripCrsRemark elements are read-only.
type Remark struct {
	RecordLocator string `json:"record_locator,omitempty" xml:"record_locator"` // read-only
	Notes         string `json:"notes,omitempty" xml:"notes"`                   // read-only
}

// Invitation contains a list of users invited to see the trip.
type Invitation struct {
	EmailAddresses    []string          `json:"EmailAddresses,omitempty" xml:"EmailAddresses"`
	TripShare         TripShare         `json:"TripShare,omitempty" xml:"TripShare"`                 // optional
	ConnectionRequest ConnectionRequest `json:"ConnectionRequest,omitempty" xml:"ConnectionRequest"` // optional
	Message           string            `json:"message,omitempty" xml:"message"`                     // optional
}

// TripShare contains information about which users a trip is shared with.
type TripShare struct {
	TripID            uint `json:"trip_id,string,omitempty" xml:"trip_id"`
	IsTraveler        bool `json:"is_traveler,string,omitempty" xml:"is_traveler"`
	IsReadOnly        bool `json:"is_read_only,string,omitempty" xml:"is_read_only"`
	IsSentWithDetails bool `json:"is_sent_with_details,string,omitempty" xml:"is_sent_with_details"`
}

// ConnectionRequest stores connection request data.
type ConnectionRequest struct{}
