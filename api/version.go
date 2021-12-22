package api

import (
	"errors"

	"golang.org/x/mod/semver"
)

// APIVersion is the API version this library is compatible with.
const APIVersion = "v3.2.0"

func init() {
	// make sure API version (in source code) is valid.
	if !semver.IsValid(APIVersion) {
		panic("api: APIVersion is invalid")
	}
}

// CheckAPIVersion checks whether the API server supports the API version the Client is compatible with.
func (c *Client) CheckAPIVersion() (string, bool, error) {
	req := VersionReq{}
	resp := VersionResp{}
	err := c.Do(req, &resp)
	if err != nil {
		return "", false, err
	}
	resp.Version = "v" + resp.Version
	if !semver.IsValid(resp.Version) {
		return "", false, errors.New("api: invalid version given")
	}
	if semver.Compare(APIVersion, resp.Version) == 1 {
		return resp.Version, false, nil
	}
	if semver.Major(APIVersion) != semver.Major(resp.Version) {
		return resp.Version, false, nil
	}
	return resp.Version, true, nil
}
