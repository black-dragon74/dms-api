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

func (d DMSSession) GetInternalMarks(semester string) (types.InternalMarksModel, error) {
	// Validate
	if !d.session.Validate() {
		return types.InternalMarksModel{}, utils.ErrorLoginFailed
	}

	// Prefetch the internals URL
	resp, err := d.session.Get(utils.AttendanceURL, nil)
	if err != nil {
		return types.InternalMarksModel{}, err
	}
	defer resp.Body.Close()

	preload, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return types.InternalMarksModel{}, err
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
		return types.InternalMarksModel{}, err
	}

	// Prepare the request body
	payload := url.Values{
		utils.IdForResultSemester:     {semester},
		utils.IdForEventTarget:        {"ctl00$ContentPlaceHolder1$ddlSemester"},
		utils.IdForViewState:          {utils.GetValueFromInput(preload, "#"+utils.IdForViewState)},
		utils.IdForViewStateGenerator: {utils.GetValueFromInput(preload, "#"+utils.IdForViewStateGenerator)},
		utils.IdForEventValidation:    {utils.GetValueFromInput(preload, "#"+utils.IdForEventValidation)},

		// Will always be the same, no point in declaring globally as a const
		"__VIEWSTATEENCRYPTED": {""},
	}.Encode()

	// Send the request
	resp, err = d.session.Post(utils.AttendanceURL, &headers, strings.NewReader(payload))
	if err != nil {
		return types.InternalMarksModel{}, err
	}
	defer resp.Body.Close()

	result, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return types.InternalMarksModel{}, err
	}

	resTable, err := utils.ParseHTMLTable(result, utils.IdForInternalsTable)
	if err != nil {
		switch err {
		case utils.ErrorTableNotFound:
			return types.InternalMarksModel{}, utils.ErrorInvalidSemester
		case utils.ErrorTableNoHeader:
			return types.InternalMarksModel{}, errors.New(fmt.Sprintf("Result for semester %s is not available", semester))
		default:
			return types.InternalMarksModel{}, err
		}
	}

	retVal := types.InternalMarksModel{}

	for _, entry := range resTable.Body {
		tMap := make(map[string]string)
		for idx, item := range entry {
			tMap[resTable.Header[idx]] = strings.TrimSpace(utils.StripParen(item))
		}
		retVal = append(retVal, tMap)
	}

	return retVal, nil
}
