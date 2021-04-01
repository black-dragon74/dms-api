package api

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/black-dragon74/dms-api/types"
	"github.com/black-dragon74/dms-api/utils"
	"net/url"
	"strings"
)

func (d DMSService) Login(userName string, password string, captcha string) (types.CaptchaAuthModel, error) {
	// Prepare the Headers
	reqHeaders := map[string][]string{
		"Accept":                    {utils.HeaderAccept},
		"Accept-Language":           {utils.HeaderAcceptLanguage},
		"Cache-Control":             {utils.HeaderCacheControl},
		"Connection":                {utils.HeaderConnection},
		"Content-Type":              {utils.HeaderContentType},
		"DNT":                       {utils.HeaderDNT},
		"Host":                      {utils.HeaderHost},
		"Origin":                    {utils.HeaderOrigin},
		"Referer":                   {utils.HeaderReferer},
		"Sec-Fetch-Dest":            {utils.HeaderSecFetchDest},
		"Sec-Fetch-Mode":            {utils.HeaderSecFetchMode},
		"Sec-Fetch-Site":            {utils.HeaderSecFetchSite},
		"Sec-Fetch-User":            {utils.HeaderSecFetchUser},
		"Upgrade-Insecure-Requests": {utils.HeaderUpgradeInsecureRequests},
		"User-Agent":                {utils.HeaderUserAgent},
	}

	// Prepare the Request Body
	reqPayload := url.Values{
		"__VIEWSTATE":          {utils.ASPLoginViewState},
		"__EVENTVALIDATION":    {utils.ASPLoginEventValidation},
		"__VIEWSTATEGENERATOR": {utils.ASPLoginViewStateGenerator},
		"__EVENTTARGET":        {utils.ASPLoginEventTarget},
		"txtUserid":            {userName},
		"txtpassword":          {password},
		"txtCaptcha":           {captcha},
	}.Encode()

	resp, err := d.session.Post(utils.LoginURL, nil, &reqHeaders, strings.NewReader(reqPayload))
	if err != nil {
		return types.CaptchaAuthModel{}, err
	}
	defer resp.Body.Close()

	retVal := types.CaptchaAuthModel{}
	retVal.SessionID = d.session.GetID()

	soup, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return types.CaptchaAuthModel{}, err
	}

	// Now parse the body and have a peek at different elements to check if login was successful
	name, exists := soup.Find("#" + utils.IdForName).Attr("value")
	if !exists {
		// A breif of how the checks are done for login
		// 1. Captcha must be valid at all times
		// 2. If captcha was valid, either the limit is exhausted or the credentials are wrong

		// Check #1
		captchaError := soup.Find(utils.IdForCaptchaError).Text()
		if strings.Contains(captchaError, "Captcha") {
			retVal.CaptchaFailed = true
		} else {
			// Check #2
			// Since captcha was not invalid now we have to check
			// If the daily limit is exhausted. If not, the credentials are invalid
			authErrror := soup.Find(utils.IdForCredentialsError).Text()
			if strings.Contains(authErrror, "exceded") {
				retVal.LoginSucceeded = true
			}
		}
	} else {
		retVal.UserName = name
		retVal.LoginSucceeded = true
	}

	return retVal, nil
}
