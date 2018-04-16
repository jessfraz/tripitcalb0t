package tripit

import (
	"fmt"
	"net/http"
)

// DeleteActivity deletes the specific activity with the given id.
func (c *Client) DeleteActivity(id string) error {
	_, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatDeleteObject, TypeActivity, id), nil)
	return err
}

// DeleteCar deletes the specific car with the given id.
func (c *Client) DeleteCar(id string) error {
	_, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatDeleteObject, TypeCar, id), nil)
	return err
}

// DeleteCruise deletes the specific cruise with the given id.
func (c *Client) DeleteCruise(id string) error {
	_, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatDeleteObject, TypeCruise, id), nil)
	return err
}

// DeleteDirections deletes the specific directions with the given id.
func (c *Client) DeleteDirections(id string) error {
	_, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatDeleteObject, TypeDirections, id), nil)
	return err
}

// DeleteFlight deletes the specific flight with the given id.
func (c *Client) DeleteFlight(id string) error {
	_, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatDeleteObject, TypeFlight, id), nil)
	return err
}

// DeleteLodging deletes the specific lodging with the given id.
func (c *Client) DeleteLodging(id string) error {
	_, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatDeleteObject, TypeLodging, id), nil)
	return err
}

// DeleteMap deletes the specific map with the given id.
func (c *Client) DeleteMap(id string) error {
	_, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatDeleteObject, TypeMap, id), nil)
	return err
}

// DeleteNote deletes the specific note with the given id.
func (c *Client) DeleteNote(id string) error {
	_, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatDeleteObject, TypeNote, id), nil)
	return err
}

// DeleteRail deletes the specific rail with the given id.
func (c *Client) DeleteRail(id string) error {
	_, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatDeleteObject, TypeRail, id), nil)
	return err
}

// DeleteRestaurant deletes the specific restaurant with the given id.
func (c *Client) DeleteRestaurant(id string) error {
	_, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatDeleteObject, TypeRestaurant, id), nil)
	return err
}

// DeleteSegment deletes the specific segment with the given id.
func (c *Client) DeleteSegment(id string) error {
	_, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatDeleteObject, TypeSegment, id), nil)
	return err
}

// DeleteTransport deletes the specific transport with the given id.
func (c *Client) DeleteTransport(id string) error {
	_, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatDeleteObject, TypeTransport, id), nil)
	return err
}

// DeleteTrip deletes the specific trip with the given id.
func (c *Client) DeleteTrip(id string) error {
	_, err := c.doRequest(http.MethodGet, fmt.Sprintf(EndpointFormatDeleteObject, TypeTrip, id), nil)
	return err
}

// DeleteTripParticipant deletes the specific participant from the trip with the given id.
func (c *Client) DeleteTripParticipant(tripID, profileRef string) error {
	_, err := c.doRequest(http.MethodGet, fmt.Sprintf("delete/trip_participant/trip_id/%s/profile_ref/%s", tripID, profileRef), nil)
	return err
}
