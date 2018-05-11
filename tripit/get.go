package tripit

import (
	"fmt"
	"net/http"
)

// GetActivity returns the specific activity for the given id.
func (c *Client) GetActivity(id string, filters ...Filter) (Activity, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatGetObject, TypeActivity, id, formatFilters(filters)), nil)
	if err != nil {
		return Activity{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if len(resp.Activities) <= 0 {
		return Activity{}, fmt.Errorf("get activity id %s returned an empty result", id)
	}

	// Return the object.
	return resp.Activities[0], nil
}

// GetCar returns the specific car for the given id.
func (c *Client) GetCar(id string, filters ...Filter) (Car, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatGetObject, TypeCar, id, formatFilters(filters)), nil)
	if err != nil {
		return Car{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if len(resp.Cars) <= 0 {
		return Car{}, fmt.Errorf("get car id %s returned an empty result", id)
	}

	// Return the object.
	return resp.Cars[0], nil
}

// GetCruise returns the specific cruise for the given id.
func (c *Client) GetCruise(id string, filters ...Filter) (Cruise, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatGetObject, TypeCruise, id, formatFilters(filters)), nil)
	if err != nil {
		return Cruise{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if len(resp.Cruises) <= 0 {
		return Cruise{}, fmt.Errorf("get cruise id %s returned an empty result", id)
	}

	// Return the object.
	return resp.Cruises[0], nil
}

// GetDirections returns the specific directions for the given id.
func (c *Client) GetDirections(id string, filters ...Filter) (Direction, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatGetObject, TypeDirections, id, formatFilters(filters)), nil)
	if err != nil {
		return Direction{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if len(resp.Directions) <= 0 {
		return Direction{}, fmt.Errorf("get directions id %s returned an empty result", id)
	}

	// Return the object.
	return resp.Directions[0], nil
}

// GetFlight returns the specific flight for the given id.
func (c *Client) GetFlight(id string, filters ...Filter) (Flight, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatGetObject, TypeFlight, id, formatFilters(filters)), nil)
	if err != nil {
		return Flight{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if len(resp.Flights) <= 0 {
		return Flight{}, fmt.Errorf("get flight id %s returned an empty result", id)
	}

	// Return the object.
	return resp.Flights[0], nil
}

// GetLodging returns the specific lodging for the given id.
func (c *Client) GetLodging(id string, filters ...Filter) (Lodging, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatGetObject, TypeLodging, id, formatFilters(filters)), nil)
	if err != nil {
		return Lodging{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if len(resp.Lodging) <= 0 {
		return Lodging{}, fmt.Errorf("get lodging id %s returned an empty result", id)
	}

	// Return the object.
	return resp.Lodging[0], nil
}

// GetMap returns the specific map for the given id.
func (c *Client) GetMap(id string, filters ...Filter) (Map, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatGetObject, TypeMap, id, formatFilters(filters)), nil)
	if err != nil {
		return Map{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if len(resp.Maps) <= 0 {
		return Map{}, fmt.Errorf("get map id %s returned an empty result", id)
	}

	// Return the object.
	return resp.Maps[0], nil
}

// GetNote returns the specific note for the given id.
func (c *Client) GetNote(id string, filters ...Filter) (Note, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatGetObject, TypeNote, id, formatFilters(filters)), nil)
	if err != nil {
		return Note{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if len(resp.Notes) <= 0 {
		return Note{}, fmt.Errorf("get note id %s returned an empty result", id)
	}

	// Return the object.
	return resp.Notes[0], nil
}

// GetPointsProgram returns the specific points program for the given id.
func (c *Client) GetPointsProgram(id string, filters ...Filter) (PointsProgram, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatGetObject, TypePointsProgram, id, formatFilters(filters)), nil)
	if err != nil {
		return PointsProgram{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if len(resp.PointsPrograms) <= 0 {
		return PointsProgram{}, fmt.Errorf("get points program id %s returned an empty result", id)
	}

	// Return the object.
	return resp.PointsPrograms[0], nil
}

// GetProfile returns the specific profile for the given id.
func (c *Client) GetProfile(id string, filters ...Filter) (Profile, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatGetObject, TypeProfile, id, formatFilters(filters)), nil)
	if err != nil {
		return Profile{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if len(resp.Profiles) <= 0 {
		return Profile{}, fmt.Errorf("get profile id %s returned an empty result", id)
	}

	// Return the object.
	return resp.Profiles[0], nil
}

// GetRail returns the specific rail for the given id.
func (c *Client) GetRail(id string, filters ...Filter) (Rail, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatGetObject, TypeRail, id, formatFilters(filters)), nil)
	if err != nil {
		return Rail{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if len(resp.Rails) <= 0 {
		return Rail{}, fmt.Errorf("get rail id %s returned an empty result", id)
	}

	// Return the object.
	return resp.Rails[0], nil
}

// GetRestaurant returns the specific restaurant for the given id.
func (c *Client) GetRestaurant(id string, filters ...Filter) (Restaurant, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatGetObject, TypeRestaurant, id, formatFilters(filters)), nil)
	if err != nil {
		return Restaurant{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if len(resp.Restaurants) <= 0 {
		return Restaurant{}, fmt.Errorf("get restaurant id %s returned an empty result", id)
	}

	// Return the object.
	return resp.Restaurants[0], nil
}

// GetTransport returns the specific transport for the given id.
func (c *Client) GetTransport(id string, filters ...Filter) (Transport, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatGetObject, TypeTransport, id, formatFilters(filters)), nil)
	if err != nil {
		return Transport{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if len(resp.Transports) <= 0 {
		return Transport{}, fmt.Errorf("get transport id %s returned an empty result", id)
	}

	// Return the object.
	return resp.Transports[0], nil
}

// GetTrip returns the specific trip for the given id.
func (c *Client) GetTrip(id string, filters ...Filter) (Trip, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatGetObject, TypeTrip, id, formatFilters(filters)), nil)
	if err != nil {
		return Trip{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if len(resp.Trips) <= 0 {
		return Trip{}, fmt.Errorf("get trip id %s returned an empty result", id)
	}

	// Return the object.
	return resp.Trips[0], nil
}

// GetWeather returns the specific weather information for the given id.
func (c *Client) GetWeather(id string, filters ...Filter) (Weather, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatGetObject, TypeWeather, id, formatFilters(filters)), nil)
	if err != nil {
		return Weather{}, err
	}

	// Check if we didn't get a result and return an error if true.
	if len(resp.Weather) <= 0 {
		return Weather{}, fmt.Errorf("get weather id %s returned an empty result", id)
	}

	// Return the object.
	return resp.Weather[0], nil
}
