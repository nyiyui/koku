package api

import (
	"net/http"
	"net/url"
)

var meURL, _ = url.Parse("me")

type MeReq struct {
	Auth
}

func (req MeReq) Req(c *Client) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(meURL).String(), nil)
	if err != nil {
		return nil, err
	}
	req.Auth.setAHeader(request.Header)
	return request, nil
}

type MeResp struct {
	Username       string   `json:"username"`
	FirstName      string   `json:"first_name"`
	LastName       string   `json:"last_name"`
	Bio            string   `json:"bio"`
	Timezone       string   `json:"timezone"`
	GraduatingYear *int     `json:"graduating_year"`
	Organizations  []string `json:"organizations"`
	TagsFollowing  []string `json:"tags_following"`
}
