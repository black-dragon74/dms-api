package config

import "fmt"

// APIConfig reflects the type for file `config.toml`
type APIConfig struct {
	redis               bool
	monitorDataStore    bool
	enableMessStore     bool
	enableContactsStore bool
	host                string
	port                int
}

func (a APIConfig) UseRedis() bool {
	return a.redis
}

func (a APIConfig) MonitorDataStore() bool {
	return a.monitorDataStore
}

func (a APIConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", a.host, a.port)
}

func (a APIConfig) MessStoreEnabled() bool {
	return a.enableMessStore
}

func (a APIConfig) ContactsStoreEnabled() bool {
	return a.enableContactsStore
}
