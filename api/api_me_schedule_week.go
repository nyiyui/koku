package api

import (
	"net/http"
	"net/url"
)

var meScheduleWeekURL, _ = url.Parse("me/schedule/week")

type MeScheduleWeekReq struct {
	Auth
}

func (req MeScheduleWeekReq) Req(c *Client) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(meScheduleWeekURL).String(), nil)
	if err != nil {
		return nil, err
	}
	req.Auth.setAHeader(request.Header)
	return request, nil
}

type MeScheduleWeekResp map[string][]Schedule
