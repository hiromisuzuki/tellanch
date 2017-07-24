package config

import (
	"strconv"

	"github.com/spf13/viper"
)

type Host struct {
	User    string
	Address string
	Port    int
	Key     string
	Path    []string
}

// ConfigManager []*Host
type ConfigManager []*Host

func (c *ConfigManager) Load() {
	for no := 0; ; {
		no++
		key := "hosts.host" + strconv.Itoa(no)
		if !viper.IsSet(key) {
			break
		}
		*c = append(*c, &Host{
			User:    viper.GetString(key + ".user"),
			Address: viper.GetString(key + ".address"),
			Port:    viper.GetInt(key + ".port"),
			Key:     viper.GetString(key + ".key"),
			Path:    viper.GetStringSlice(key + ".path"),
		})
	}
}
