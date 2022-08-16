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
	lgr.Info("[Initialize] [DataStore] Loading stores in memory")
	store := &types.DataStoreModel{}

	// Use wait groups to keep track of sub-goroutines
	wg := &sync.WaitGroup{}
	wg.Add(2)

	// Load both the stores into memory, concurrently
	go loadContactsStore(lgr, &cfg.API, &store.ContactsData, wg)
	go loadMessStore(lgr, &cfg.API, &store.MessMenuData, wg)

	// Wait for go-routines to exit
	wg.Wait()

	// Validate contacts store first, only if it is enabled
	if cfg.API.ContactsStoreEnabled() {
		if len(store.ContactsData) == 0 {
			return nil, errors.New("failed to read from the contacts datastore")
		}
	}

	// Validate the mess menu store
	if cfg.API.MessStoreEnabled() {
		if store.MessMenuData.LastUpdatedAt == "" {
			return nil, errors.New("failed to read from the mess menu datastore")
		}
	}

	// Watch for changes if requested
	if cfg.API.MonitorDataStore() {
		if !(cfg.API.MessStoreEnabled() && cfg.API.ContactsStoreEnabled()) {
			lgr.Info("[Initialize] [DataStore] [WatchStoreForChanges] Monitoring disabled as both the stores are disabled")
			return store, nil
		}

		go watchStoreForChanges(store, cfg, lgr)
	} else {
		lgr.Info("[Initialize] [DataStore] [WatchStoreForChanges] Monitoring disabled by config")
	}

	return store, nil
}

func watchStoreForChanges(store *types.DataStoreModel, cfg *config.Config, lgr *zap.Logger) {
	lgr.Info("[Initialize] [DataStore] [WatchStoreForChanges] Actively monitoring stores for changes")
	messTicker := time.NewTicker(5 * time.Minute)
	contactsTicker := time.NewTicker(24 * time.Hour)

	for {
		select {
		case <-messTicker.C:
			loadMessStore(lgr, &cfg.API, &store.MessMenuData, nil)

		case <-contactsTicker.C:
			loadContactsStore(lgr, &cfg.API, &store.ContactsData, nil)
		}
	}
}

// loadContactsStore queries the contact data store and loads the response into memory
func loadContactsStore(lgr *zap.Logger, cfg *config.APIConfig, store *[]types.ContactsModel, wg *sync.WaitGroup) {
	// Notify the wait group, it will be nil if called by monitoring go-routine
	if wg != nil {
		defer wg.Done()
	}

	// If we are not supposed to load, return right away
	if !cfg.ContactsStoreEnabled() {
		lgr.Info("[Initialize] [DataStore] [LoadContactsStore] Loading of contacts store disabled via config")
		return
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
func loadMessStore(lgr *zap.Logger, cfg *config.APIConfig, store *types.MessMenuModel, wg *sync.WaitGroup) {
	// Notify the wait group, it will be nil if called by monitoring go-routine
	if wg != nil {
		defer wg.Done()
	}

	// If we are not supposed to load, return right away
	if !cfg.MessStoreEnabled() {
		lgr.Info("[Initialize] [DataStore] [LoadMessStore] Loading of mess store disabled via config")
		return
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
