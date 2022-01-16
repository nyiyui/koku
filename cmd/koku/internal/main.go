package internal

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"nyiyui.ca/koku/api"
)

var instanceName string

// Main runs the program.
func Main() (err error) {
	var credentialsPath string
	var configPath string
	var rawBaseURL string
	var instanceName string
	flag.StringVar(&credentialsPath, "credentials", "", "path to the credentials file")
	flag.StringVar(&configPath, "config", "", "path to the config file")
	flag.StringVar(&rawBaseURL, "base-url", api.DefaultBaseURL.String(), "base url of the api")
	flag.StringVar(&instanceName, "instance", "", "name of this instance")
	flag.Parse()

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		err = fmt.Errorf("base-url: %w", err)
		return
	}
	apiCl := api.NewClient(&http.Client{}, baseURL)

	instanceName = "koku@" + instanceName + uuid.New().String()
	log.Printf("instance name: %s", instanceName)

	app, err := newApp(credentialsPath, configPath)
	if err != nil {
		err = fmt.Errorf("new app: %w", err)
		return
	}
	err = setupRealtime(context.Background(), apiCl, app)
	return
}
