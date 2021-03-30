package api

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/black-dragon74/dms-api/utils"
	"go.uber.org/zap"
	"net/http"
)

type Session struct {
	sid string
	lgr *zap.Logger
}

func (s Session) Validate() bool {
	// Make a get request to profileURL and check for reg no ID's value
	// It should not be null
	resp, err := s.Get(utils.ProfileURL, nil, nil)
	if err != nil {
		s.lgr.Error(fmt.Sprintf("[Session] [Validate] %s", err.Error()))
		return false
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		s.lgr.Error(fmt.Sprintf("[Session] [Validate] [NewDocumentFromReader] %s", err.Error()))
		return false
	}

	reg := doc.Find("#" + utils.IdForRegno)
	_, e := reg.Attr("value")

	return e
}

func NewSession(sessionID string, lgr *zap.Logger) Session {
	return Session{
		sid: sessionID,
		lgr: lgr,
	}
}

func (s Session) Get(url string, cookies *map[string]string, headers *map[string][]string) (*http.Response, error) {
	// Add session specific headers
	if cookies != nil {
		(*cookies)[utils.SessionCookie] = s.sid
	} else {
		cookies = &map[string]string{utils.SessionCookie: s.sid}
	}

	if headers != nil {
		(*headers)[utils.SessionCookie] = []string{s.sid}
	} else {
		headers = &map[string][]string{utils.SessionCookie: {s.sid}}
	}

	request := utils.NewRequest("GET", url, cookies, headers)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
