package config

import (
	"github.com/spf13/viper"
	"strconv"
)

type Host struct {
	User    string
	Address string
	Port    int
	Key     string
	Path    []string
}

// Manager []*Host
type Manager []*Host

func (c *Manager) Load() {
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
