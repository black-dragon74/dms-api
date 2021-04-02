package api

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/black-dragon74/dms-api/types"
	"github.com/black-dragon74/dms-api/utils"
)

func (d DMSSession) GetAttendance() (types.AttendanceModel, error) {
	// Validate the Session
	if !d.session.Validate() {
		return types.AttendanceModel{}, utils.ErrorLoginFailed
	}

	// Get the attendance page
	resp, err := d.session.Get(utils.AttendanceURL, nil)
	if err != nil {
		return types.AttendanceModel{}, err
	}
	defer resp.Body.Close()

	soup, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return types.AttendanceModel{}, err
	}

	tbl, err := utils.ParseHTMLTable(soup, utils.IdForAttendanceTable)
	if err != nil {
		switch err {
		case utils.ErrorTableNoBody:
			return types.AttendanceModel{}, utils.ErrorNoAttendance
		default:
			return types.AttendanceModel{}, err
		}
	}

	retVal := types.AttendanceModel{}
	attendanceStruct := []string{"index", "course", "status", "type", "section", "batch", "present", "absent", "total", "percentage"}
	for _, entry := range tbl.Body {
		tMap := make(map[string]string)
		for i, val := range entry {
			tMap[attendanceStruct[i]] = val
		}
		retVal = append(retVal, tMap)
	}

	return retVal, nil
}
