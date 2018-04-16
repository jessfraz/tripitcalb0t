package tripit

import (
	"fmt"
	"net/http"
)

// ListTrips returns a list of trips and other object data depending on the filters passed.
func (c *Client) ListTrips(filters ...Filter) (*Response, error) {
	return c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s", ListTripsEndpoint, formatFilters(filters)), nil)
}

// ListObjects returns a list of objects and other data depending on the filters passed.
func (c *Client) ListObjects(filters ...Filter) (*Response, error) {
	return c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s", ListObjectsEndpoint, formatFilters(filters)), nil)
}

// ListPointsPrograms returns a list of points programs depending on the filters passed.
func (c *Client) ListPointsPrograms(filters ...Filter) ([]PointsProgram, error) {
	resp, err := c.doRequest(http.MethodGet, fmt.Sprintf("%s/%s", ListPointsProgramsEndpoint, formatFilters(filters)), nil)
	if err != nil {
		return nil, err
	}

	return resp.PointsPrograms, nil
}
