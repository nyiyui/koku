package api

import (
	"net/http"
	"net/url"
)

var userURL, _ = url.Parse("user/")

type UserReq struct {
	Username string
}

func (req UserReq) Req(c *Client) (*http.Request, error) {
	username, err := url.Parse(url.PathEscape(req.Username))
	if err != nil {
		return nil, err
	}
	return http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(userURL).ResolveReference(username).String(), nil)
}

type UserResp User
