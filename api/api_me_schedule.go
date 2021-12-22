package api

import (
	"net/http"
	"net/url"
)

var meScheduleURL, _ = url.Parse("me/schedule")

type MeScheduleReq struct {
	Auth
}

func (req MeScheduleReq) Req(c *Client) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(meScheduleURL).String(), nil)
	if err != nil {
		return nil, err
	}
	req.Auth.setAHeader(request.Header)
	return request, nil
}

type MeScheduleResp = []Schedule
