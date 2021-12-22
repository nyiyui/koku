package internal

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
	"nyiyui.ca/koku/api"
)

func Main() (err error) {
	var credentialsPath string
	var configPath string
	var baseURL_ string
	flag.StringVar(&credentialsPath, "credentials", "", "path to the credentials file")
	flag.StringVar(&configPath, "config", "", "path to the config file")
	flag.StringVar(&baseURL_, "base-url", api.DefaultBaseURL.String(), "base url of the api")
	flag.Parse()

	baseURL, err := url.Parse(baseURL_)
	if err != nil {
		err = fmt.Errorf("base-url: %w", err)
		return
	}
	apiCl := api.NewClient(&http.Client{}, baseURL)

	app, err := NewApp(credentialsPath, configPath)
	if err != nil {
		err = fmt.Errorf("new app: %w", err)
		return
	}
	err = Setup(apiCl, app)
	return
}

func NewApp(credentialsPath, configPath string) (app *firebase.App, err error) {
	opt := option.WithCredentialsFile(credentialsPath)
	cfg := &firebase.Config{}
	file, err := os.Open(configPath)
	if err != nil {
		err = fmt.Errorf("config open: %w", err)
		return
	}
	defer func(file *os.File) {
		if err2 := file.Close(); err2 != nil {
			err = fmt.Errorf("config close: %w", err2)
		}
	}(file)
	err = json.NewDecoder(file).Decode(cfg)
	if err != nil {
		err = fmt.Errorf("config unmarshal: %w", err)
		return
	}
	app, err = firebase.NewApp(context.Background(), cfg, opt)
	return
}

func Setup(apiCl *api.Client, app *firebase.App) (err error) {
	client, err := app.Messaging(context.Background())
	if err != nil {
		return
	}

	annCh := make(chan api.AnnChangesResp)
	errs := make(chan error)
	err = api.AnnChanges(apiCl, annCh, errs)
	if err != nil {
		err = fmt.Errorf("api: %w", err)
		return
	}
	go sendOnCh(client, annCh, errs)
	for err := range errs {
		log.Println("error: ", err)
	}
	return
}

func sendOnCh(cl *messaging.Client, ch <-chan api.AnnChangesResp, errs chan<- error) {
	log.Println("sendOnCh")
	for resp := range ch {
		log.Println("sendOnCh:", resp)
		data, err := json.Marshal(resp)
		messages := []*messaging.Message{
			{
				Data: map[string]string{
					"type": "announcement-changes",
					"data": string(data),
				},
				Notification: &messaging.Notification{
					Title: "Announcement Changeed",
				},
				Topic: "announcement-changes",
			},
		}

		br, err := cl.SendAll(context.Background(), messages)
		if err != nil {
			errs <- err
		}

		log.Printf("sendOnCh: success=%d failure=%d", br.SuccessCount, br.FailureCount)
		for i, r := range br.Responses {
			log.Printf("sendOnCh: i=%d success=%t messageID=%s error=%v", i, r.Success, r.MessageID, r.Error)
		}
	}
}
