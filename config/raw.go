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
	Global    any
	Include   []string
	Init      any
	E2e_Tests any
	Services  map[string]any
	Purge     any
	Cmds      map[string]any
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

	var c map[string]any
	err = v.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("c: %v\n", c)

	z, o := c["init"].(string)
	if o {
		fmt.Printf("Is string %v", z)
	}

	ee := ParseCmd(c["e2e_tests"])
	fmt.Printf("Command %+v", ee)

	// c.loadIncludes(p)

	return RawConfig{}
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
