package api

import (
	"net/http"
	"net/url"
)

var versionURL, _ = url.Parse("version")

type VersionReq struct{}

func (req VersionReq) Req(c *Client) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(versionURL).String(), nil)
}

type VersionResp struct {
	Version string `json:"version"`
}
