package initialize

import (
	"github.com/black-dragon74/dms-api/config"
	"github.com/black-dragon74/dms-api/types"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

func DataStore(lgr *zap.Logger, cfg *config.Config) (*types.GlobalDataStore, error) {
	lgr.Info("[Initilaize] [DataStore] Loading stores in memory")
	store := &types.GlobalDataStore{}

	data, err := read(cfg.API.GetMessMenuDataStore())
	if err != nil {
		return &types.GlobalDataStore{}, err
	}
	store.MessMenuData = data

	data, err = read(cfg.API.GetFacultyDataStore())
	if err != nil {
		return &types.GlobalDataStore{}, err
	}
	store.ContactsData = data

	go watchStoreForChanges(cfg, store, lgr)

	return store, nil
}

func read(file string) ([]byte, error) {
	return ioutil.ReadFile(file)
}

func watchStoreForChanges(cfg *config.Config, store *types.GlobalDataStore, lgr *zap.Logger) {
	lgr.Info("[Initialize] [DataStore] [WatchStoreForChanges] Actively monitoring store for changes")
	messTicker := time.NewTicker(5 * time.Minute)
	contactsTicker := time.NewTicker(4 * time.Hour)

	for {
		select {
		case <-messTicker.C:
			lgr.Info("[Initialize] [DataStore] [WatchStoreForChanges] [MessTicker] Tick")
			data, err := read(cfg.API.GetMessMenuDataStore())
			if err != nil {
				return
			}
			store.MessMenuData = data

		case <-contactsTicker.C:
			lgr.Info("[Initialize] [DataStore] [WatchStoreForChanges] [ContactsTicker] Tick")
			data, err := read(cfg.API.GetFacultyDataStore())
			if err != nil {
				return
			}
			store.ContactsData = data
		}
	}
}
