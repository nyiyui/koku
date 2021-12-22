package api

import (
	"net/http"
	"net/url"
)

var authReURL, _ = url.Parse("auth/token/refresh")

type AuthReReq struct {
	Auth
}

func (req AuthReReq) Req(c *Client) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPost, c.baseURL.ResolveReference(authReURL).String(), nil)
	if err != nil {
		return nil, err
	}
	req.Auth.setRHeader(request.Header)
	return request, nil
}

type AuthReResp struct {
	Access string `json:"access"`
}

func (resp AuthReResp) UpdateAuth(a *Auth) {
	a.Access = resp.Access
}
