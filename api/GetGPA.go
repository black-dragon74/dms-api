package api

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/black-dragon74/dms-api/types"
	"github.com/black-dragon74/dms-api/utils"
)

func (d DMSSession) GetGPA() (types.GPAModel, error) {
	if !d.session.Validate() {
		return types.GPAModel{}, utils.ErrorLoginFailed
	}

	// Get the GPA data
	resp, err := d.session.Get(utils.AttendanceURL, nil)
	if err != nil {
		return types.GPAModel{}, err
	}
	defer resp.Body.Close()

	// Create a DOM object
	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return types.GPAModel{}, err
	}

	gpaTable, err := utils.ParseHTMLTable(dom, utils.IdForCgpaTable)
	if err != nil {
		switch err {
		case utils.ErrorTableNoHeader:
			// There is no data for this semester
			return types.GPAModel{}, errors.New("GPA data is not available yet")
		default:
			// Return whtever the error is
			return types.GPAModel{}, err
		}
	}

	retVal := types.GPAModel{}

	for _, semester := range gpaTable.Body {
		for idx, marks := range semester {
			if marks == "0.00" {
				continue
			}
			retVal[gpaTable.Header[idx]] = marks
		}
	}

	return retVal, nil
}
