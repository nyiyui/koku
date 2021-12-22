package api

import (
	"bytes"
	"net/http"
)

type buffer struct{ bytes.Buffer }

func (b *buffer) Close() error {
	b.Reset()
	return nil
}

type dummyClient struct {
	m map[string]func() *http.Response
}

func (d *dummyClient) Get(url string) (resp *http.Response, err error) {
	newResp, ok := d.m[url]
	if !ok {
		return &http.Response{
			Status:        "404 Not Found",
			StatusCode:    http.StatusNotFound,
			Proto:         "HTTP/2",
			ProtoMajor:    2,
			ProtoMinor:    0,
			Header:        http.Header{},
			Body:          new(buffer),
			ContentLength: 0,
			Close:         false,
			Uncompressed:  false,
			Trailer:       nil,
			Request:       nil,
			TLS:           nil,
		}, nil
	}
	return newResp(), nil
}

func testClient() *Client {
	return DefaultClient()
}
