package initialize

import (
	"encoding/json"
	"github.com/black-dragon74/dms-api/config"
	"github.com/black-dragon74/dms-api/types"
	"github.com/black-dragon74/dms-api/utils"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

func DataStore(lgr *zap.Logger, cfg *config.Config) (*types.DataStoreModel, error) {
	lgr.Info("[Initilaize] [DataStore] Loading stores in memory")
	store := &types.DataStoreModel{}

	// Load both the stores into memory
	err := loadContactsStore(lgr, &store.ContactsData)
	if err != nil {
		return nil, err
	}

	err = loadMessStore(lgr, &store.MessMenuData)
	if err != nil {
		return nil, err
	}

	// Watch for changes if requested
	if cfg.API.MonitorDataStore() {
		go watchStoreForChanges(store, lgr)
	} else {
		lgr.Info("[Initialize] [DataStore] [WatchStoreForChanges] Monitoring disabled by config")
	}

	return store, nil
}

func watchStoreForChanges(store *types.DataStoreModel, lgr *zap.Logger) {
	lgr.Info("[Initialize] [DataStore] [WatchStoreForChanges] Actively monitoring stores for changes")
	messTicker := time.NewTicker(5 * time.Minute)
	contactsTicker := time.NewTicker(24 * time.Hour)

	for {
		select {
		case <-messTicker.C:
			lgr.Info("[Initialize] [DataStore] [WatchStoreForChanges] [MessTicker] Tick")
			_ = loadMessStore(lgr, &store.MessMenuData)

		case <-contactsTicker.C:
			lgr.Info("[Initialize] [DataStore] [WatchStoreForChanges] [ContactsTicker] Tick")
			_ = loadContactsStore(lgr, &store.ContactsData)
		}
	}
}

// loadContactsStore queries the contacts data store and loads the response into memory
func loadContactsStore(lgr *zap.Logger, store *[]types.ContactsModel) error {
	// Fetch the contacts store
	resp, err := http.Get(utils.URLContactsDataStore)
	if err != nil {
		lgr.Error("[Initialize] [DataStore] Unable to fetch the contacts data store")
		return err
	}
	defer resp.Body.Close()

	// Unmarshal to the type
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lgr.Error("[Initialize] [DataStore] Unable to read the contacts store response")
		return err
	}

	err = json.Unmarshal(data, store)
	if err != nil {
		lgr.Error("[Initialize] [DataStore] Contatcs store returned an unexpected response")
		return err
	}
	lgr.Info("[Initialize] [DataStore] Contacts store loaded successfully")

	return nil
}

// loadMessStore queries the mess menu data store and loads the response into memory
func loadMessStore(lgr *zap.Logger, store *types.MessMenuModel) error {
	resp, err := http.Get(utils.URLMessDataStore)
	if err != nil {
		lgr.Error("[Initialize] [DataStore] Unable to fetch the mess data store")
		return err
	}
	defer resp.Body.Close()

	// Unmarshal to the type
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lgr.Error("[Initialize] [DataStore] Unable to read the mess store response")
		return err
	}

	err = json.Unmarshal(data, store)
	if err != nil {
		lgr.Error("[Initialize] [DataStore] Mess store returned an unexpected response")
		return err
	}
	lgr.Info("[Initialize] [DataStore] Mess store loaded successfully")

	return nil
}
