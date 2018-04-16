package tripit

import (
	"net/http"
)

// Create takes a Request object and creates it.
func (c *Client) Create(req Request) (*Response, error) {
	return c.doRequest(http.MethodPost, "v1/create", req)
}
