package api

import (
	"net/http"
)

// Auth stores basic information for authentication.
type Auth struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func (a Auth) aHeader() string               { return "Bearer " + a.Access }
func (a Auth) rHeader() string               { return "Bearer " + a.Refresh }
func (a Auth) setRHeader(header http.Header) { header.Set("Authorization", a.rHeader()) }
func (a Auth) setAHeader(header http.Header) { header.Set("Authorization", a.aHeader()) }
