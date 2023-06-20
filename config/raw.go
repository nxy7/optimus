package config

import (
	"fmt"
	"log"
	"optimus/utils"
	"strings"

	"github.com/spf13/viper"
)

// If config field has few shapes (like string or some struct) we're using any to parse it into said struct later on
type rawConfig struct {
	Global    any            `mapstructure:"global"`
	Include   []string       `mapstructure:"include"`
	Init      any            `mapstructure:"init"`
	E2e_Tests any            `mapstructure:"e2e_tests"`
	Services  map[string]any `mapstructure:"services"`
	Purge     any            `mapstructure:"purge"`
	Cmds      []any          `mapstructure:"cmds"`
}

func LoadRawConfig() rawConfig {
	dirPath := utils.ProjectRoot()
	c := loadConfigFromPath(dirPath)

	return c
}

func loadConfigFromPath(p string) rawConfig {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("optimus")
	v.AddConfigPath(p)

	// v.SetDefault("global", map[string]string{"shell_cmd": "bash -c"})
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("Could not read config at path: %v\n%v", p, err)
	}

	var c rawConfig
	err = v.Unmarshal(&c)
	if err != nil {
		fmt.Println(err)
		panic("Could not marchal config")
	}

	c.loadIncludes(p)

	return c
}

func (rc *rawConfig) loadIncludes(dirPath string) {
	for _, v := range rc.Include {
		newPath := strings.Replace(v, "./", "/", 1)
		loadedConfig := loadConfigFromPath(dirPath + newPath)
		new := mergeRawConfigs(*rc, loadedConfig)
		*rc = new
	}
}

func mergeRawConfigs(c1 rawConfig, c2 rawConfig) rawConfig {
	if c1.Global != nil && c2.Global != nil {
		panic("Global field specified multiple times")
	}

	if c1.Init != nil && c2.Init != nil {
		panic("Init field specified multiple times")
	}

	if c1.Init != nil && c2.Init != nil {
		panic("Init field specified multiple times")
	}

	return rawConfig{
		Global:    c1.Global,
		Include:   []string{},
		Init:      nil,
		E2e_Tests: nil,
		Services:  map[string]any{},
		Purge:     nil,
		Cmds:      []any{},
	}
}
