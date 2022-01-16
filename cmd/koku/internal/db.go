package internal

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"nyiyui.ca/koku/api"
)

// DB is a wrapper around a Firestore database.
type DB struct {
	c *firestore.Client
}

func newDB(ctx context.Context, projectID string) (*DB, error) {
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return &DB{c: client}, nil
}

func fromApp(ctx context.Context, app *firebase.App) (*DB, error) {
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to make firestore client: %v", err)
	}
	return &DB{c: client}, nil
}

// Close closes the connection to the database.
func (db *DB) Close() error {
	return db.c.Close()
}

// Subs is a slice of subscriptions.
type Subs []Sub

func (s *Subs) matchTag(tag string) Subs {
	var subs Subs
	for _, sub := range *s {
		if sub.matchTag(tag) {
			subs = append(subs, sub)
		}
	}
	return subs
}

func (s *Subs) matchClub(club string) Subs {
	var subs Subs
	for _, sub := range *s {
		if sub.matchClub(club) {
			subs = append(subs, sub)
		}
	}
	return subs
}

// Sub is a subscription.
type Sub struct {
	ID                string
	Username          string    `firestore:"username"`
	RefreshToken      string    `firestore:"refresh_token"`
	RegistrationToken string    `firestore:"registration_token"`
	FollowingClubs    []string  `firestore:"following_clubs"`
	FollowingTags     []string  `firestore:"following_tags"`
	Verified          bool      `firestore:"verified"`
	VerifiedAt        time.Time `firestore:"verified_at"`
	VerifiedBy        string    `firestore:"verified_by"`
}

func (s *Sub) matchTag(tag string) bool {
	for _, t := range s.FollowingTags {
		if t == tag {
			return true
		}
	}
	return false
}

func (s *Sub) matchClub(club string) bool {
	for _, c := range s.FollowingClubs {
		if c == club {
			return true
		}
	}
	return false
}

func (db *DB) getSubscriptions(ctx context.Context) (Subs, error) {
	iter := db.c.Collection("subscriptions").Documents(ctx)
	var subscriptions Subs
	// TODO: use goroutines
	for {
		snap, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		subscription := Sub{
			ID: snap.Ref.ID,
		}
		err = snap.DataTo(&subscription)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal subscription: %v", err)
		}
		if !subscription.Verified {
			err = db.verifySubscription(&subscription)
			if err != nil {
				return nil, fmt.Errorf("failed to verify subscription %s: %v", snap.Ref.ID, err)
			}
		}
		subscriptions = append(subscriptions, subscription)
	}
	return subscriptions, nil
}

func (db *DB) verifySubscription(sub *Sub) (err error) {
	cl := api.DefaultClient()

	var authRe api.AuthReResp
	err = cl.Do(api.AuthReReq{
		Auth: api.Auth{Refresh: sub.RefreshToken},
	}, &authRe)
	if err != nil {
		err = fmt.Errorf("failed to refresh token: %v", err)
		return
	}

	var me api.MeResp
	err = cl.Do(api.MeReq{Auth: api.Auth{Access: authRe.Access}}, &me)
	if err != nil {
		err = fmt.Errorf("failed to get me: %v", err)
		return
	}

	sub.Username = me.Username
	sub.FollowingClubs = me.Organizations
	sub.FollowingTags = me.TagsFollowing
	sub.Verified = true
	sub.VerifiedAt = time.Now()
	sub.VerifiedBy = instanceName
	return
}
