package config

import (
	"fmt"
	"optimus/utils"
	"strings"

	"github.com/spf13/viper"
)

// If config field has few shapes (like string or some struct) we're using any to parse it into said struct later on
type rawConfig struct {
	global    any
	include   []string
	init      any
	e2e_tests any
	services  map[string]any
	purge     any
	cmds      []any
}

func LoadRawConfig() rawConfig {
	dirPath := utils.ProjectRoot()
	fmt.Println(dirPath)
	c := loadConfigFromPath(dirPath)

	return c
}

func loadConfigFromPath(p string) rawConfig {
	viper.SetConfigType("yaml")
	// viper.SetConfigType("json")
	viper.SetConfigName("optimus")
	viper.AddConfigPath(p)

	// viper.SetDefault("global", map[string]string{"shell_cmd": "bash -c"})
	err := viper.ReadInConfig()
	if err != nil {
		panic("Could not read config")
	}

	var c rawConfig
	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println(err)
		panic("Could not marchal config")
	}

	c.loadIncludes(p)

	return c
}

func (rc *rawConfig) loadIncludes(dirPath string) {
	for _, v := range rc.include {
		newPath := strings.Replace(v, "./", "/", 1)
		loadedConfig := loadConfigFromPath(dirPath + newPath)
		new := mergeRawConfigs(*rc, loadedConfig)
		*rc = new
	}
}

func mergeRawConfigs(c1 rawConfig, c2 rawConfig) rawConfig {
	if c1.global != nil && c2.global != nil {
		panic("Global field specified multiple times")
	}

	if c1.init != nil && c2.init != nil {
		panic("Init field specified multiple times")
	}

	if c1.init != nil && c2.init != nil {
		panic("Init field specified multiple times")
	}

	return rawConfig{
		global:    c1.global,
		include:   []string{},
		init:      nil,
		e2e_tests: nil,
		services:  map[string]any{},
		purge:     nil,
		cmds:      []any{},
	}
}
