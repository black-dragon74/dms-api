package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	env string
	API APIConfig
}

func (c Config) GetEnv() string {
	return c.env
}

// Load loads the config.toml file for the server and logger configuration
func Load() Config {
	// Load the viper
	err := loadViper()
	if err != nil {
		log.Fatal("Unable to load viper")
	}

	myCfg := Config{
		env: viper.GetString("env"),
		API: APIConfig{
			redis: viper.GetBool("api.redis"),
			host:  viper.GetString("api.host"),
			port:  viper.GetInt("api.port"),
		},
	}

	return myCfg
}

func loadViper() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
