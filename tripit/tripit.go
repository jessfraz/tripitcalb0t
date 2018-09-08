package tripit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

// Client holds the information needed for TripIt API authentication.
type Client struct {
	username string
	password string
}

// New creates a new TripIt API client.
func New(username, password string) *Client {
	return &Client{
		username: username,
		password: password,
	}
}

func (c *Client) doRequest(method, endpoint string, data interface{}) (*Response, error) {
	client := http.DefaultClient

	// Encode data if we are passed an object.
	b := bytes.NewBuffer(nil)
	if data != nil {
		// Create the encoder.
		enc := json.NewEncoder(b)
		if err := enc.Encode(data); err != nil {
			return nil, fmt.Errorf("json encoding data for doRequest failed: %v", err)
		}
	}

	// Create the request.
	uri := fmt.Sprintf("%s/%s/%s/format/json", APIUri, APIVersion, strings.Trim(endpoint, "/"))
	req, err := http.NewRequest(method, uri, b)
	if err != nil {
		return nil, fmt.Errorf("creating %s request to %s failed: %v", method, uri, err)
	}

	// Set the basic auth credentials.
	req.SetBasicAuth(c.username, c.password)

	// Do the request.
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("performing %s request to %s failed: %v", method, uri, err)
	}
	defer resp.Body.Close()

	// Check that the response status code was OK.
	if resp.StatusCode != http.StatusOK {
		// Read the body of the request, ignore the error since we are already in the error state.
		body, _ := ioutil.ReadAll(resp.Body)

		// Create a friendly error message based off the status code returned.
		// These come from: https://tripit.github.io/api/doc/v1/index.html#http_status_codes_section
		var message string
		switch resp.StatusCode {
		case http.StatusBadRequest: // 400
			message = "The request was either invalid or malformed in some way. For example, a create call with no xml or json request parameter in the POST args would return a 400 Bad Request from the server."
		case http.StatusUnauthorized: // 401
			message = `The OAuth Consumer has errored for one of the following reasons:
1) The authentication credentials passed to the API for the request were somehow invalid. This status code could be caused by an invalid username/password combination used in the web authentication scheme or an invalid OAuth token for the OAuth scheme.
2) The TripIt account for which the consumer was authorized is no longer authorizing it.
3) The OAuth Consumer key has been de-activated.`
		case http.StatusForbidden: // 403
			message = "The OAuth Consumer is not yet confirmed. The most common situation in which this happens is when a new account that authorizes an API client hasn't been confirmed before the API client attempts to execute a read operation on the API (e.g. /v1/list/trip)."
		case http.StatusNotFound: // 404
			message = "Either the resource URL or the object the client was requesting either does not exist or the user the client was authenticated and does not have permission to operate on the object."
		case http.StatusInternalServerError: // 500
			message = "Something catastrophic happened while the TripIt platform was trying to complete the request. A 500 error is a pretty serious and catastrophic problem that should be reported to the TripIt engineering team through support@tripit.com."
		case http.StatusServiceUnavailable: // 503
			message = "The TripIt API is currently undergoing maintenance and is not available."
		}

		return nil, fmt.Errorf("%s request to %s returned status code %d: message -> %s\nbody -> %s", method, uri, resp.StatusCode, message, string(body))
	}
	/*body, _ := ioutil.ReadAll(resp.Body)
	var out bytes.Buffer
	json.Indent(&out, body, "", "  ")
	logrus.Fatalf("body: %s", string(body))*/

	// Decode the response into a TripIt Response object.
	var r Response
	if err := decodeResponse(resp, &r); err != nil {
		// Read the body of the request, ignore the error since we are already in the error state.
		body, _ := ioutil.ReadAll(resp.Body)

		return nil, fmt.Errorf("decoding response from %s request to %s failed: body -> %s\nerr -> %v", method, uri, string(body), err)
	}

	// Log warnings on the API warnings.
	for _, warning := range r.Warnings {
		logrus.Warnf("[%s] %s: %s", warning.Timestamp, warning.EntityType, warning.Description)
	}

	// Log errors on the API errors.
	for _, e := range r.Errors {
		logrus.Errorf("[%s] %s code -> %d, detailed code -> %f: %s", e.Timestamp, e.EntityType, e.Code, e.DetailedErrorCode, e.Description)
	}

	return &r, nil
}

func decodeResponse(resp *http.Response, v interface{}) error {
	// Copy buffer and change "@attributes" to "_attributes" since the json package doesn't support "@".
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return err
	}
	b := bytes.Replace(buf.Bytes(), []byte(`"@attributes"`), []byte(`"_attributes"`), -1)

	return json.Unmarshal(b, v)
}
