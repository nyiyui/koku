package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"nyiyui.ca/koku/api"
)

func setupRealtime(ctx context.Context, apiCl *api.Client, app *firebase.App) (err error) {
	client, err := app.Messaging(ctx)
	if err != nil {
		return
	}

	d, err := fromApp(ctx, app)
	if err != nil {
		err = fmt.Errorf("setupRealtime: newDB: %w", err)
		return
	}

	notifCh := make(chan api.NotifNewResp)
	errs := make(chan error)
	go api.NotifNew(apiCl, notifCh, errs)
	go sendOnCh(client, notifCh, errs, d)
	if err != nil {
		err = fmt.Errorf("api: %w", err)
		return
	}
	for err := range errs {
		log.Println("error: ", err)
	}
	return
}

// topicsToCond generates a condition from a slice of topics.
func topicsToCond(topics []string) (cond string) {
	if len(topics) == 0 {
		return
	}
	topics2 := make([]string, 0, len(topics))
	for _, t := range topics {
		if t == "" {
			continue
		}
		topics2 = append(topics2, "'"+t+"' in topics") // TODO: escape
	}
	return strings.Join(topics2, " || ")
}

// tagsToTopics converts a slice of tags to a slice of topics.
func tagsToTopics(tags []api.Tag) (topics []string) {
	for _, t := range tags {
		topics = append(topics, t.Name)
	}
	return
}

// msg creates a message from notification data.
func msg(data interface{}) (*messaging.Message, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	switch data := data.(type) {
	case api.Ann:
		return &messaging.Message{
			Data: map[string]string{
				"type": "announcement-changes",
				"data": string(jsonData),
			},
			Notification: &messaging.Notification{
				Title: "Announcement Changed",
			},
			Condition: topicsToCond(append(tagsToTopics(data.Tags), data.Org.Name)),
		}, nil
	case api.BlogPost:
		return &messaging.Message{
			Data: map[string]string{
				"type": "blogpost-changes",
				"data": string(jsonData),
			},
			Notification: &messaging.Notification{
				Title: "Blog Post Changeed",
			},
			Condition: topicsToCond(tagsToTopics(data.Tags)),
		}, nil
	default:
		return nil, fmt.Errorf("unknown type: %T", data)
	}
}

// sendOnCh sends messages on the given channel.
// It is blocking.
func sendOnCh(cl *messaging.Client, ch <-chan api.NotifNewResp, errs chan<- error, d *DB) {
	log.Println("sendOnCh")
	//ctx := context.Background()
	//subs, err := d.getSubscriptions(ctx)
	//if err != nil {
	//	errs <- fmt.Errorf("sendOnCh: getSubscriptions: %w", err)
	//	return
	//}
	log.Println("sendOnCh: got subscriptions")
	for resp := range ch {
		log.Println("sendOnCh:", resp)
		//subs := subs.matchTag(resp.Tag)
		m, err := msg(resp)
		if err != nil {
			errs <- err
			continue
		}
		raw, _ := json.MarshalIndent(m, "", "  ")
		log.Println("sendOnCh:", string(raw))
		messages := []*messaging.Message{m}
		br, err := cl.SendAll(context.Background(), messages)
		if err != nil {
			errs <- err
			continue
		}

		log.Printf("sendOnCh: success=%d failure=%d", br.SuccessCount, br.FailureCount)
		for i, r := range br.Responses {
			log.Printf("sendOnCh: i=%d success=%t messageID=%s error=%v", i, r.Success, r.MessageID, r.Error)
		}
	}
}
