package config

import "sync"

type Config struct {
	DB string
}

func (c *Config) SetDB(val string) {
	c.DB = val
}
func (c *Config) GetDB() string {
	return c.DB
}

var (
	once     sync.Once
	instance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{"localhost:3306"}
	})
	return instance
}
