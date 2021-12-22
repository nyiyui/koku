package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type query interface { // TODO: document this
	Query() url.Values
}

// queryReq returns the *http.Request for a request.
func queryReq(u2 *url.URL, q query, c *Client) (*http.Request, error) {
	var u *url.URL
	{ // copy u
		tmp := *u2
		u = &tmp
	}
	u.RawQuery = q.Query().Encode()
	return http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(u).String(), nil)
}

// jsonEncode encodes v to a buffer.
func jsonEncode(v interface{}) (io.Reader, error) {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(v)
	if err != nil {
		return nil, err
	}
	return body, nil
}
