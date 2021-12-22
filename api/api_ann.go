package api

import (
	"net/http"
	"net/url"
)

var annURL, _ = url.Parse("announcements")

type AnnReq struct{}

func (req AnnReq) Req(c *Client) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(annURL).String(), nil)
}

type AnnResp = []Ann
