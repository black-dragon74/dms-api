package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/black-dragon74/dms-api/config"
	"github.com/black-dragon74/dms-api/utils"
	"github.com/go-redis/redis/v8"
	"io"
	"io/ioutil"
	"net/http"
)

type Session struct {
	sid string
	cfg *config.Config
	rds *redis.Client
}

func (s Session) GetID() string {
	return s.sid
}

func (s Session) GetRedisClient() *redis.Client {
	return s.rds
}

func (s Session) Validate() bool {
	// If redis store is to be used, check from it else follow conventional method
	if s.cfg.API.UseRedis() {
		valid := validateSessionFromCache(s)

		if valid {
			if e := utils.UpdateSessionExpiry(s.GetID(), s.GetRedisClient()); e != nil {
				fmt.Printf("[ERROR] [API] [Session] [UpdateSessionExpiry] %s\n", e.Error())
			}
		}

		return valid
	}

	// Make a get request to profileURL and check for reg no ID's value
	// It should not be null
	resp, err := s.Get(utils.ProfileURL, nil)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return false
	}

	reg := utils.GetValueFromInput(doc, utils.IdForRegno)

	return reg != ""
}

func NewSession(sessionID string, cfg *config.Config, rds *redis.Client) Session {
	return Session{
		sid: sessionID,
		cfg: cfg,
		rds: rds,
	}
}

func (s Session) Get(url string, headers *map[string][]string) (*http.Response, error) {
	// Add session specific headers
	cookies := &map[string]string{utils.SessionCookie: s.sid}

	if headers != nil {
		(*headers)[utils.SessionCookie] = []string{s.sid}
	} else {
		headers = &map[string][]string{utils.SessionCookie: {s.sid}}
	}

	request := utils.NewRequest(http.MethodGet, url, cookies, headers, nil)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s Session) Post(url string, headers *map[string][]string, body io.Reader) (*http.Response, error) {
	// Add session specific headers
	cookies := &map[string]string{utils.SessionCookie: s.sid}

	if headers != nil {
		(*headers)[utils.SessionCookie] = []string{s.sid}
	} else {
		headers = &map[string][]string{utils.SessionCookie: {s.sid}}
	}

	bodyCloser := ioutil.NopCloser(body)
	request := utils.NewRequest(http.MethodPost, url, cookies, headers, &bodyCloser)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func validateSessionFromCache(sess Session) bool {
	// Query the client for the session ID
	_, err := sess.rds.Get(context.Background(), sess.sid).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		fmt.Printf("[ERROR] [API] [Session] [ValidateSessionFromCache] %s\n", err.Error())
	}

	return err == nil
}
