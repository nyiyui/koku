package api

import "testing"

func TestAuth(t *testing.T) {
	c := testClient()
	resp := AuthResp{}
	err := c.Do(AuthReq{
		Username: "superuser",
		Password: "superuser",
	}, &resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
