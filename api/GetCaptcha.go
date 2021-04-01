package api

import (
	"encoding/base64"
	"github.com/PuerkitoBio/goquery"
	"github.com/black-dragon74/dms-api/api/internal"
	"github.com/black-dragon74/dms-api/types"
	"github.com/black-dragon74/dms-api/utils"
	"io/ioutil"
	"net/http"
)

func GetCaptcha() (types.GetCaptchaModel, error) {
	// Get login URL
	resp, err := http.Get(utils.LoginURL)
	if err != nil {
		return types.GetCaptchaModel{}, err
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
		return types.GetCaptchaModel{}, err
	}

	// Get the captcha generator URL
	v, e := doc.Find("#" + utils.IdForCaptcha).Attr("src")
	if !e {
		return types.GetCaptchaModel{}, err
	}
	retVal.Generator = utils.DmsURL + v

	// Now get the new URL with session
	sess := internal.NewSession(retVal.SessionID)
	resp, err = sess.Get(retVal.Generator, nil)
	if err != nil {
		return types.GetCaptchaModel{}, err
	}
	defer resp.Body.Close()

	rawByte, _ := ioutil.ReadAll(resp.Body)
	retVal.EncodedImage = base64.StdEncoding.EncodeToString(rawByte)

	return retVal, nil
}
