package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

var notifNewURL, _ = url.Parse("notifications/new")

// NotifNewResp is the response for the data from the announcement changes stream.
type NotifNewResp interface{}

// NotifNew forwards ann changes to ch.
func NotifNew(cl *Client, ch chan<- NotifNewResp, errs chan<- error) {
	defer close(ch)
	request, err := http.NewRequest(http.MethodGet, cl.BaseURL().ResolveReference(notifNewURL).String(), nil)
	if err != nil {
		err = fmt.Errorf("http req: %w", err)
		return
	}
	resp, err := cl.HTTPClient().Do(request)
	if err != nil {
		err = fmt.Errorf("http do: %w", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http status: %d", resp.StatusCode)
		return
	}
	defer func(body io.ReadCloser) {
		err2 := body.Close()
		if err2 != nil {
			err = fmt.Errorf("error closing response body: %w", err2)
		}
	}(resp.Body)

	scanner := bufio.NewScanner(resp.Body)
	if err != nil {
		errs <- err
		return
	}
	var eventName string
	for scanner.Scan() {
		if scanner.Err() != nil {
			errs <- fmt.Errorf("scanner error: %w", scanner.Err())
			continue
		}
		line := scanner.Bytes()
		log.Printf("event %s", line)
		switch {
		case bytes.HasPrefix(line, []byte("event: ")):
			eventName = string(line[7:])
		case bytes.HasPrefix(line, []byte("data: ")):
			switch eventName {
			case "announcement_change":
				var resp Ann
				if err = json.Unmarshal(line[6:], &resp); err != nil {
					errs <- fmt.Errorf("json unmarshal: %w", err)
					continue
				}
				ch <- resp
			case "blogpost_change":
				var resp BlogPost
				if err = json.Unmarshal(line[6:], &resp); err != nil {
					errs <- fmt.Errorf("json unmarshal: %w", err)
					continue
				}
				ch <- resp
			}
		}
	}
	return
}
