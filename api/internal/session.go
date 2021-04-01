package internal

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/black-dragon74/dms-api/utils"
	"io"
	"io/ioutil"
	"net/http"
)

type Session struct {
	sid string
}

func (s Session) GetID() string {
	return s.sid
}

func (s Session) Validate() bool {
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

func NewSession(sessionID string) Session {
	return Session{
		sid: sessionID,
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
