package api

import (
	"net/http"
	"net/url"
	"strconv"
)

var orgURL, _ = url.Parse("organization")

type OrgReq struct {
	Id int
}

func (req OrgReq) Req(c *Client) (*http.Request, error) {
	username, err := url.Parse(url.PathEscape(strconv.FormatInt(int64(req.Id), 10)))
	if err != nil {
		return nil, err
	}
	return http.NewRequest(http.MethodPost, c.baseURL.ResolveReference(orgURL).ResolveReference(username).String(), nil)
}

type OrgResp = Org
