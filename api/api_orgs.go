package api

import (
	"net/http"
	"net/url"
)

var orgsURL, _ = url.Parse("organizations")

type OrgsReq struct{}

func (req OrgsReq) Req(c *Client) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(orgsURL).String(), nil)
}

type OrgsResp = []Org
