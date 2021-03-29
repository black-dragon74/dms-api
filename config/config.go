package config

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Env string
	API APIConfig
}

// Load loads the config.toml file for the server and logger configuration
func Load() Config {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Failed to load the config file", err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Unable to read from the file", err)
	}

	myCfg := Config{}
	err = toml.Unmarshal(data, &myCfg)
	if err != nil {
		log.Fatal("Unable to de-serialize the config", err)
	}

	return myCfg
}
