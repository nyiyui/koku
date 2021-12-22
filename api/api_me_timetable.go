package api

import (
	"net/http"
	"net/url"
)

var meTimetableURL, _ = url.Parse("me/timetable")

type MeTimetableReq struct {
	Auth
}

func (req MeTimetableReq) Req(c *Client) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(meTimetableURL).String(), nil)
	if err != nil {
		return nil, err
	}
	req.Auth.setAHeader(request.Header)
	return request, nil
}

type MeTimetableResp struct {
	Id      int      `json:"id"`
	Owner   User     `json:"owner"`
	Term    Term     `json:"term"`
	Courses []Course `json:"courses"`
}
