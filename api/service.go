package api

import (
	"github.com/black-dragon74/dms-api/config"
	"io/ioutil"
)

type DMSService struct {
	Session Session
	cfg     *config.Config
}

func NewDMSService(sid string, config *config.Config) DMSService {
	return DMSService{
		Session: newSession(sid),
		cfg:     config,
	}
}

// GetMessMenu reads the menu JSON file and returns the raw data
func (svc DMSService) GetMessMenu() ([]byte, error) {

	// Read the store file defined in `config.toml`
	data, err := ioutil.ReadFile(svc.cfg.API.GetMessMenuDataStore())
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

// GetContacts reads the contacts JSON file and returns the raw data
func (svc DMSService) GetContacts() ([]byte, error) {

	// Read the store file defined in `config.toml`
	data, err := ioutil.ReadFile(svc.cfg.API.GetFacultyDataStore())
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}
