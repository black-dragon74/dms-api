package api

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/black-dragon74/dms-api/types"
	"github.com/black-dragon74/dms-api/utils"
	"net/url"
	"strings"
)

func (d DMSSession) GetResult(semester string) (types.ResultModel, error) {
	// Validate the session
	if !d.session.Validate() {
		return types.ResultModel{}, utils.ErrorLoginFailed
	}

	// Prefetch the result URL
	resp, err := d.session.Get(utils.ResultURL, nil)
	if err != nil {
		return types.ResultModel{}, err
	}
	defer resp.Body.Close()

	preload, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return types.ResultModel{}, err
	}

	// Prepare the Headers
	headers := map[string][]string{
		"Cache-Control":    {utils.HeaderCacheControl},
		"Connection":       {utils.HeaderConnection},
		"Content-Type":     {utils.HeaderContentType},
		"DNT":              {utils.HeaderDNT},
		"Host":             {utils.HeaderHost},
		"Origin":           {utils.HeaderOrigin},
		"Referer":          {utils.HeaderReferer},
		"User-Agent":       {utils.HeaderUserAgent},
		"X-MicrosoftAjax":  {utils.HeaderXMicrosoftAjax},
		"X-Requested-With": {utils.HeaderXRequestedWith},
	}

	semester, err = utils.SwitchSemester(semester)
	if err != nil {
		return types.ResultModel{}, err
	}

	// Prepare the request body
	payload := url.Values{
		utils.IdForResultSemester:     {semester},
		utils.IdForEventTarget:        {"ctl00$ContentPlaceHolder1$ddlSemester"},
		utils.IdForViewState:          {utils.GetValueFromInput(preload, "#"+utils.IdForViewState)},
		utils.IdForViewStateGenerator: {utils.GetValueFromInput(preload, "#"+utils.IdForViewStateGenerator)},
		utils.IdForEventValidation:    {utils.GetValueFromInput(preload, "#"+utils.IdForEventValidation)},

		// Will always be the same, no point in declaring globally as a const
		"__EVENTARGUMENT":      {""},
		"__LASTFOCUS":          {""},
		"__ASYNCPOST":          {"true"},
		"ctl00$ScriptManager1": {"ctl00$ContentPlaceHolder1$mrak|ctl00$ContentPlaceHolder1$ddlSemester"},
	}.Encode()

	// Send the post request
	resp, err = d.session.Post(utils.ResultURL, &headers, strings.NewReader(payload))
	if err != nil {
		return types.ResultModel{}, err
	}
	defer resp.Body.Close()

	result, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return types.ResultModel{}, err
	}

	resTable, err := utils.ParseHTMLTable(result, utils.IdForResultTable)
	if err != nil {
		switch err {
		case utils.ErrorTableNotFound:
			// Semester is out of range
			return types.ResultModel{}, utils.ErrorInvalidSemester
		case utils.ErrorTableNoHeader:
			// There is no data for this semester
			return types.ResultModel{}, errors.New(fmt.Sprintf("result for semester %s is not available", semester))
		default:
			// Return whtever the error is
			return types.ResultModel{}, err
		}
	}

	resultStruct := []string{"index", "course_code", "course_name", "academic_session", "credits", "grade"}
	retVal := types.ResultModel{}

	for _, subject := range resTable.Body {
		tmpMap := make(map[string]string)
		for i, val := range subject {
			tmpMap[resultStruct[i]] = val
		}
		retVal = append(retVal, tmpMap)
	}

	return retVal, nil
}
