package config

import (
	"fmt"
	"log"
	"optimus/utils"
	"strings"

	"github.com/spf13/viper"
)

// If config field has few shapes (like string or some struct) we're using any to parse it into said struct later on
type RawConfig struct {
	Global    any            `mapstructure:"global"`
	Include   []string       `mapstructure:"include"`
	Init      any            `mapstructure:"init"`
	E2e_Tests any            `mapstructure:"e2e_tests"`
	Services  map[string]any `mapstructure:"services"`
	Purge     any            `mapstructure:"purge"`
	Cmds      map[string]any `mapstructure:"cmds"`
}

func LoadRawConfig() RawConfig {
	dirPath := utils.ProjectRoot()
	c := loadConfigFromPath(dirPath)

	return c
}

func loadConfigFromPath(p string) RawConfig {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("optimus")
	v.AddConfigPath(p)

	// v.SetDefault("global", map[string]string{"shell_cmd": "bash -c"})
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("Could not read config at path: %v\n%v", p, err)
	}

	var c RawConfig
	err = v.Unmarshal(&c)
	if err != nil {
		fmt.Println(err)
		panic("Could not marchal config")
	}

	c.loadIncludes(p)

	return c
}

func (rc *RawConfig) loadIncludes(dirPath string) {
	for _, v := range rc.Include {
		newPath := strings.Replace(v, "./", "/", 1)
		loadedConfig := loadConfigFromPath(dirPath + newPath)
		rc.mergeRawConfigs(loadedConfig)
	}
}

func (rc *RawConfig) mergeRawConfigs(newConfig RawConfig) {
	for k, v := range newConfig.Services {
		rc.Services[k] = v
	}
	for i, v := range newConfig.Cmds {
		rc.Services[i] = v
	}
}
