package api

import (
	"encoding/base64"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/black-dragon74/dms-api/types"
	"github.com/black-dragon74/dms-api/utils"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type DMSService struct {
	session Session
}

func GetCaptcha(lgr *zap.Logger) ([]byte, error) {
	// Get login URL
	resp, err := http.Get(utils.LoginURL)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	// Get the session ID from the request
	retVal := types.GetCaptchaModel{}
	sessionID := utils.GetSessionFromResponse(resp)

	if sessionID != "" {
		retVal.SessionID = sessionID
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	// Get the captcha generator URL
	v, e := doc.Find("#" + utils.IdForCaptcha).Attr("src")
	if !e {
		return []byte{}, err
	}
	retVal.Generator = utils.DmsURL + v

	// Now get the new URL with session
	sess := NewSession(retVal.SessionID, lgr)
	resp, err = sess.Get(retVal.Generator, nil, nil)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	rawByte, _ := ioutil.ReadAll(resp.Body)
	retVal.EncodedImage = base64.StdEncoding.EncodeToString(rawByte)

	data, err := json.Marshal(retVal)

	return data, nil
}
