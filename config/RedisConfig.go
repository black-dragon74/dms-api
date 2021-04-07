package config

import (
	"fmt"
)

type RedisConfig struct {
	dbid int
	host string
	pass string
	port int
}

func (r RedisConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", r.host, r.port)
}

func (r RedisConfig) GetPassword() string {
	return r.pass
}

func (r RedisConfig) GetDBId() int {
	return r.dbid
}
