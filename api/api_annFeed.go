package api

import (
	"net/http"
	"net/url"
)

var annFeedURL, _ = url.Parse("announcements/feed")

type AnnFeedReq struct {
	Auth
}

func (req AnnFeedReq) Req(c *Client) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(annFeedURL).String(), nil)
	if err != nil {
		return nil, err
	}
	req.Auth.setAHeader(request.Header)
	return request, nil
}

type AnnFeedResp []Ann
