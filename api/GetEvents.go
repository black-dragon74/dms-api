package api

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/black-dragon74/dms-api/types"
	"github.com/black-dragon74/dms-api/utils"
)

func (d DMSSession) GetEvents() (types.EventsModel, error) {
	if !d.session.Validate() {
		return types.EventsModel{}, utils.ErrorLoginFailed
	}

	resp, err := d.session.Get(utils.EventURL, nil)
	if err != nil {
		return types.EventsModel{}, err
	}
	defer resp.Body.Close()

	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return types.EventsModel{}, err
	}

	table, err := utils.ParseHTMLTable(dom, utils.IdForEventsTable)
	if err != nil {
		return types.EventsModel{}, err
	}

	eventsStruct := []string{"name", "date", "description"}
	retVal := types.EventsModel{}

	for _, event := range table.Body {
		tMap := make(map[string]string)
		for i, e := range event {
			tMap[eventsStruct[i]] = e
		}
		retVal = append(retVal, tMap)
	}

	return retVal, nil
}
