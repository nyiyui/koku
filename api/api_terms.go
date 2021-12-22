package api

import (
	"net/http"
	"net/url"
)

var termsURL, _ = url.Parse("terms")

type TermsReq struct{}

func (req TermsReq) Req(c *Client) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(termsURL).String(), nil)
}

type TermsResp struct{}
