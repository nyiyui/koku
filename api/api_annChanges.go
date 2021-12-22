package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var annChangesURL, _ = url.Parse("announcements/changes")

// AnnChangesStreamName is the name of the event used for the announcement changes stream.
// For example, the client might receive this:
//     event: announcement-changes
//     data: {"id": 1, "author": {"id": 1, "slug": "su"}, "organization": {"id": 1, "slug": "club"}, "tags": [], "created_date": "2021-12-20T20:49:37.969222-05:00", "last_modified_date": "2021-12-21T17:40:14.080525-05:00", "title": "an announcement", "body": "", "is_public": false}
const annChangesStreamName = "announcement-changes"

// AnnChangesResp is the response for the data from the announcement changes stream.
type AnnChangesResp Ann

// AnnChanges forwards ann changes to ch.
func AnnChanges(cl *Client, ch chan<- AnnChangesResp, errs chan<- error) (err error) {
	defer close(ch)
	request, err := http.NewRequest(http.MethodGet, cl.BaseURL().ResolveReference(annChangesURL).String(), nil)
	if err != nil {
		err = fmt.Errorf("http req: %w", err)
		return
	}
	resp, err := cl.HTTPClient().Do(request)
	if err != nil {
		err = fmt.Errorf("http do: %w", err)
		return
	}
	defer func(body io.ReadCloser) {
		err2 := body.Close()
		if err2 != nil {
			err = fmt.Errorf("error closing response body: %w", err2)
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http status: %d", resp.StatusCode)
		return
	}

	go func() {
		var eventName string
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Bytes()
			switch {
			case bytes.HasPrefix(line, []byte("event: ")):
				eventName = string(line[7:])
			case bytes.HasPrefix(line, []byte("data: ")):
				switch eventName {
				case annChangesStreamName:
					var resp AnnChangesResp
					if err = json.Unmarshal(line[6:], &resp); err != nil {
						errs <- fmt.Errorf("json unmarshal: %w", err)
						continue
					}
					ch <- resp
				}
			}
		}
	}()
	return
}
