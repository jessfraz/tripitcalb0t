package tripit

import (
	"fmt"
	"net/http"
)

// ReplaceActivity replaces the activity with the given id.
func (c *Client) ReplaceActivity(id string, activity Activity) (*Response, error) {
	req := Request{
		Activity: activity,
	}
	return c.doRequest(http.MethodPost, fmt.Sprintf(EndpointFormatReplaceObject, TypeActivity, id), req)
}

// ReplaceCar replaces the car with the given id.
func (c *Client) ReplaceCar(id string, car Car) (*Response, error) {
	req := Request{
		Car: car,
	}
	return c.doRequest(http.MethodPost, fmt.Sprintf(EndpointFormatReplaceObject, TypeCar, id), req)
}

// ReplaceCruise replaces the cruise with the given id.
func (c *Client) ReplaceCruise(id string, cruise Cruise) (*Response, error) {
	req := Request{
		Cruise: cruise,
	}
	return c.doRequest(http.MethodPost, fmt.Sprintf(EndpointFormatReplaceObject, TypeCruise, id), req)
}

// ReplaceDirections replaces the directions with the given id.
func (c *Client) ReplaceDirections(id string, directions Directions) (*Response, error) {
	req := Request{
		Directions: directions,
	}
	return c.doRequest(http.MethodPost, fmt.Sprintf(EndpointFormatReplaceObject, TypeDirections, id), req)
}

// ReplaceFlight replaces the flight with the given id.
func (c *Client) ReplaceFlight(id string, flight Flight) (*Response, error) {
	req := Request{
		Flight: flight,
	}
	return c.doRequest(http.MethodPost, fmt.Sprintf(EndpointFormatReplaceObject, TypeFlight, id), req)
}

// ReplaceLodging replaces the lodging with the given id.
func (c *Client) ReplaceLodging(id string, lodging Lodging) (*Response, error) {
	req := Request{
		Lodging: lodging,
	}
	return c.doRequest(http.MethodPost, fmt.Sprintf(EndpointFormatReplaceObject, TypeLodging, id), req)
}

// ReplaceMap replaces the map with the given id.
func (c *Client) ReplaceMap(id string, m Map) (*Response, error) {
	req := Request{
		Map: m,
	}
	return c.doRequest(http.MethodPost, fmt.Sprintf(EndpointFormatReplaceObject, TypeMap, id), req)
}

// ReplaceNote replaces the note with the given id.
func (c *Client) ReplaceNote(id string, note Note) (*Response, error) {
	req := Request{
		Note: note,
	}
	return c.doRequest(http.MethodPost, fmt.Sprintf(EndpointFormatReplaceObject, TypeNote, id), req)
}

// ReplaceRail replaces the rail with the given id.
func (c *Client) ReplaceRail(id string, rail Rail) (*Response, error) {
	req := Request{
		Rail: rail,
	}
	return c.doRequest(http.MethodPost, fmt.Sprintf(EndpointFormatReplaceObject, TypeRail, id), req)
}

// ReplaceRestaurant replaces the restaurant with the given id.
func (c *Client) ReplaceRestaurant(id string, restaurant Restaurant) (*Response, error) {
	req := Request{
		Restaurant: restaurant,
	}
	return c.doRequest(http.MethodPost, fmt.Sprintf(EndpointFormatReplaceObject, TypeRestaurant, id), req)
}

// ReplaceTransport replaces the transport with the given id.
func (c *Client) ReplaceTransport(id string, transport Transport) (*Response, error) {
	req := Request{
		Transport: transport,
	}
	return c.doRequest(http.MethodPost, fmt.Sprintf(EndpointFormatReplaceObject, TypeTransport, id), req)
}

// ReplaceTrip replaces the trip with the given id.
func (c *Client) ReplaceTrip(id string, trip Trip) (*Response, error) {
	req := Request{
		Trip: trip,
	}
	return c.doRequest(http.MethodPost, fmt.Sprintf(EndpointFormatReplaceObject, TypeTrip, id), req)
}
