package initialize

import (
	"encoding/json"
	"errors"
	"github.com/black-dragon74/dms-api/config"
	"github.com/black-dragon74/dms-api/types"
	"github.com/black-dragon74/dms-api/utils"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func DataStore(lgr *zap.Logger, cfg *config.Config) (*types.DataStoreModel, error) {
	lgr.Info("[Initilaize] [DataStore] Loading stores in memory")
	store := &types.DataStoreModel{}

	// Use waitgroups to keep track of sub-goroutines
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// Load both the stores into memory, concurrently
	go loadContactsStore(lgr, &store.ContactsData, wg)
	go loadMessStore(lgr, &store.MessMenuData, wg)

	// Wait for go-routines to exit
	wg.Wait()

	// Validate
	if len(store.ContactsData) == 0 || store.MessMenuData.LastUpdatedAt == "" {
		return nil, errors.New("failed to init the stores")
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
			loadMessStore(lgr, &store.MessMenuData, nil)

		case <-contactsTicker.C:
			lgr.Info("[Initialize] [DataStore] [WatchStoreForChanges] [ContactsTicker] Tick")
			loadContactsStore(lgr, &store.ContactsData, nil)
		}
	}
}

// loadContactsStore queries the contacts data store and loads the response into memory
func loadContactsStore(lgr *zap.Logger, store *[]types.ContactsModel, wg *sync.WaitGroup) {
	// Notify the wait group, it will be nil if called by monitoring go-routine
	if wg != nil {
		defer wg.Done()
	}

	// Fetch the contacts store
	resp, err := http.Get(utils.URLContactsDataStore)
	if err != nil {
		lgr.Error("[Initialize] [DataStore] Unable to fetch the contacts data store")
		return
	}
	defer resp.Body.Close()

	// Unmarshal to the type
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lgr.Error("[Initialize] [DataStore] Unable to read the contacts store response")
		return
	}

	err = json.Unmarshal(data, store)
	if err != nil {
		lgr.Error("[Initialize] [DataStore] Contatcs store returned an unexpected response")
		return
	}
	lgr.Info("[Initialize] [DataStore] Contacts store loaded successfully")
}

// loadMessStore queries the mess menu data store and loads the response into memory
func loadMessStore(lgr *zap.Logger, store *types.MessMenuModel, wg *sync.WaitGroup) {
	// Notify the wait group, it will be nil if called by monitoring go-routine
	if wg != nil {
		defer wg.Done()
	}

	resp, err := http.Get(utils.URLMessDataStore)
	if err != nil {
		lgr.Error("[Initialize] [DataStore] Unable to fetch the mess data store")
		return
	}
	defer resp.Body.Close()

	// Unmarshal to the type
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		lgr.Error("[Initialize] [DataStore] Unable to read the mess store response")
		return
	}

	err = json.Unmarshal(data, store)
	if err != nil {
		lgr.Error("[Initialize] [DataStore] Mess store returned an unexpected response")
		return
	}
	lgr.Info("[Initialize] [DataStore] Mess store loaded successfully")
}
