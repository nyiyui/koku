package api

import (
	"net/http"
	"net/url"
)

var authURL, _ = url.Parse("auth/token")

type AuthReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req AuthReq) Req(c *Client) (*http.Request, error) {
	encoded, err := jsonEncode(req)
	if err != nil {
		return nil, err
	}
	return http.NewRequest(http.MethodPost, c.baseURL.ResolveReference(authURL).String(), encoded)
}

type AuthResp struct {
	Refresh string `json:"refresh"`
	Access  string `json:"access"`
}

func (resp AuthResp) UpdateAuth(a *Auth) {
	a.Access = resp.Access
	a.Refresh = resp.Refresh
}
