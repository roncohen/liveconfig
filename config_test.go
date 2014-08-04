package liveconfig

import (
	"testing"
)

type TestConfig struct {
	Config1 string
}

func TestCanLoadAndGet(t *testing.T) {
	test_config := TestConfig{"Hello"}
	var out_config TestConfig
	c := NewConfig(test_config)

	c.GetConfig(&out_config)
	if out_config.Config1 != test_config.Config1 {
		t.Error("Config mismatch ", out_config, test_config)
	}
}

func TestCanReloadAndGet(t *testing.T) {
	test_config1 := TestConfig{"Hello1"}
	test_config2 := TestConfig{"Hello2"}
	var out_config TestConfig
	c := NewConfig(test_config1)

	c.LoadConfig(test_config2)

	c.GetConfig(&out_config)

	if out_config.Config1 != test_config2.Config1 {
		t.Error("Config missmatch")
	}
}
