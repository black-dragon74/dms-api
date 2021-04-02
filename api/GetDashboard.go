package api

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/black-dragon74/dms-api/types"
	"github.com/black-dragon74/dms-api/utils"
)

func (d DMSSession) GetDashboard() (types.DashboardModel, error) {
	// Validate the session
	if !d.session.Validate() {
		return types.DashboardModel{}, utils.ErrorLoginFailed
	}

	// Otherwise, make a get request to fetch the student profile
	resp, err := d.session.Get(utils.ProfileURL, nil)
	if err != nil {
		return types.DashboardModel{}, err
	}
	defer resp.Body.Close()

	// Init the reply model
	retVal := types.DashboardModel{}
	soup, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return types.DashboardModel{}, err
	}

	// Read and set the values to the reply
	retVal.AdmDetails.Name = utils.GetValueFromInput(soup, utils.IdForName)
	retVal.AdmDetails.RegNo = utils.GetValueFromInput(soup, utils.IdForRegno)
	retVal.AdmDetails.Program = utils.GetValueFromInput(soup, utils.IdForPgme)
	retVal.AdmDetails.AcadYear = utils.GetValueFromInput(soup, utils.IdForAcadYear)

	// Parse the marks table
	marksTable, err := utils.ParseHTMLTable(soup, utils.IdForMarksTable)
	if err != nil {
		return types.DashboardModel{}, err
	}

	// Read and update the data
	qualStruct := []string{"index", "board", "institution", "grade", "year", "max_marks", "obtained_marks", "percentage"}
	for _, qual := range marksTable.Body {
		tMap := make(map[string]string)
		for i, v := range qual {
			tMap[qualStruct[i]] = v
		}
		retVal.EduQualifications = append(retVal.EduQualifications, tMap)
	}

	// Update parent values
	retVal.ParentDetails.Email = utils.GetValueFromInput(soup, utils.IdForParentEmail)
	retVal.ParentDetails.EmergencyContact = utils.GetValueFromInput(soup, utils.IdForParentEmergency)
	retVal.ParentDetails.Father = utils.GetValueFromInput(soup, utils.IdForParentFather)
	retVal.ParentDetails.MobileNo = utils.GetValueFromInput(soup, utils.IdForParentMobile)
	retVal.ParentDetails.Mother = utils.GetValueFromInput(soup, utils.IdForParentMother)

	return retVal, nil
}
