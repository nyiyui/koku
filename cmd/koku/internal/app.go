package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func newApp(credentialsPath, configPath string) (app *firebase.App, err error) {
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
