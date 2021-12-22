package api

import (
	"net/http"

	"golang.org/x/oauth2"
)

// OauthReq adds oauth authentication to a Req.
type OauthReq struct {
	Token *oauth2.Token
	Inner Req
}

// Req implements Req.
func (req OauthReq) Req(c *Client) (*http.Request, error) {
	request, err := req.Inner.Req(c)
	if err != nil {
		return nil, err
	}
	req.Token.SetAuthHeader(request)
	return request, nil
}
