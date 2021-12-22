package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
)

// Client abstracts requesting the server.
type Client struct {
	client  HTTPClient
	baseURL *url.URL
	noCopy  [0]sync.Mutex
	auth    Auth
}

// DefaultBaseURL is the default base URL (The Lyon's Den)
var DefaultBaseURL, _ = url.Parse("https://maclyonsden.com/api/")

// DefaultClient returns a new client using the default URL.
//
// Deprecated: use NewClient instead.
func DefaultClient() *Client {
	return &Client{client: http.DefaultClient, baseURL: DefaultBaseURL}
}

// NewClient returns a new client.
func NewClient(client HTTPClient, baseURL *url.URL) *Client {
	return &Client{client: client, baseURL: baseURL}
}

// Rel resolves a relative URL to the base URL.
func (c *Client) Rel(u *url.URL) *url.URL {
	if u.IsAbs() {
		return u
	}
	return c.baseURL.ResolveReference(u)
}

// HTTPClient returns the HTTPClient this Client is using.
func (c *Client) HTTPClient() HTTPClient {
	return c.client
}

// BaseURL returns the base URL this Client is using.
func (c *Client) BaseURL() *url.URL {
	return c.baseURL
}

// Do performs a Req and unmarshals its result to v.
func (c *Client) Do(req Req, v interface{}) (err error) {
	request, err := req.Req(c)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err2 := Body.Close()
		if err2 != nil {
			err = err2
		}
	}(resp.Body)
	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("%s: %s", resp.Status, body)
	}
	err = json.NewDecoder(resp.Body).Decode(v)
	if err != nil {
		return fmt.Errorf("json: %s", err)
	}
	log.Println(request.Header, "\n\n", request, "\n", resp, "\n\n", v)
	return
}
