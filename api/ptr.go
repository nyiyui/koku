package api

import (
	"errors"
)

// AnnID points to an Ann.
type AnnID uint

// Deref dereferences itself.
func (id AnnID) Deref(c *Client) (Ann, error) {
	resp := AnnResp{}
	err := c.Do(AnnReq{}, &resp)
	if err != nil {
		return Ann{}, err
	}
	for _, ann := range resp {
		if ann.ID == id {
			return ann, nil
		}
	}
	return Ann{}, errors.New("not found")
}

// Username points to a User.
type Username string

// Deref dereferences itself.
func (username Username) Deref(c *Client) (resp UserResp, err error) {
	err = c.Do(UserReq{
		Username: string(username),
	}, &resp)
	return
}

// TagID points to a Tag.
type TagID = uint

// OrgName is the name of an organization, which may not be unique.
type OrgName = string

// OrgSlug points to an Org.
type OrgSlug string

// Deref dereferences itself.
func (slug OrgSlug) Deref(c *Client) (org Org, err error) {
	resp := OrgsResp{}
	err = c.Do(OrgsReq{}, &resp)
	if err != nil {
		return Org{}, err
	}
	for _, org := range resp {
		if org.Slug == slug {
			return org, nil
		}
	}
	return Org{}, errors.New("not found")
}

// EventID points to an Event.
type EventID = int // negative is evs not on server (generated)

// TermID points to a Term.
type TermID = uint

// CourseID points to a Course.
type CourseID = uint

// CourseCode points to a Course.
type CourseCode string

// Deref dereferences itself.
func (code CourseCode) Deref(c *Client) (Course, error) {
	resp := MeTimetableResp{}
	err := c.Do(MeTimetableReq{c.auth}, &resp)
	if err != nil {
		return Course{}, err
	}
	for _, course := range resp.Courses {
		if course.Code == code {
			return course, nil
		}
	}
	return Course{}, errors.New("not found")
}
