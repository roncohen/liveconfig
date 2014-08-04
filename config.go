/*
liveconfig provides a way to load configurations and notably reload configurations
in live systems in a threadsafe way
*/
package liveconfig

import (
	"reflect"
	"sync"
)

// A configuration
type Config struct {
	// An arbitrary object that represents configuration options
	config     interface{}
	configLock *sync.RWMutex
}

func NewConfig(config interface{}) *Config {
	return &Config{config, new(sync.RWMutex)}
}

// Load newConfig as the new configuration
func (c *Config) LoadConfig(newConfig interface{}) {
	c.configLock.Lock()
	c.config = newConfig
	c.configLock.Unlock()
	return
}

// Get the currently loaded configuration threadsafely.x
func (c *Config) GetConfig(confobj interface{}) {
	c.configLock.RLock()
	defer c.configLock.RUnlock()

	from := reflect.ValueOf(c.config)
	to := reflect.ValueOf(confobj).Elem()

	for i := 0; i < from.NumField(); i++ {
		f := from.Field(i)
		t := to.Field(i)

		if t.CanSet() {
			t.Set(f)
		}
	}
}
