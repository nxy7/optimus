package config

import (
	"log"
	"optimus/utils"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Global             *Global
	BuildCmd           *Cmd
	PushCmd            *Cmd
	E2eTests           *Cmd
	Services           map[string]*Service
	AdditionalCommands map[string]*Cmd
}

type E2eTests struct {
	Cmd string
}

func LoadConfig() Config {
	p := utils.ProjectRoot()

	conf := LoadConfigFromPath(p)
	conf.MergeConfigs(DefaultConfig())

	return conf
}

func ParseConfig(a map[string]any, confPath string) Config {
	conf := Config{
		Services:           make(map[string]*Service),
		AdditionalCommands: make(map[string]*Cmd),
	}
	include := make([]string, 0)
	for k, v2 := range a {
		if k == "include" {
			strv2, o := v2.([]any)
			if !o {
				panic("Wrong include format")
			}

			for _, v3 := range strv2 {
				includePath, o := v3.(string)
				if !o {
					panic("include only accepts strings")
				}
				includePath = strings.Replace(includePath, "./", "/", 1)
				include = append(include, includePath)
			}
		} else if k == "global" {
			g := ParseGlobal(v2)
			conf.Global = &g
		} else if k == "e2e_tests" {
			c := ParseCmd(k, "", v2)
			conf.E2eTests = &c
		} else if k == "services" {
			servicesAny, o := v2.(map[string]any)
			if !o {
				panic("Unexpected services format")
			}
			for svcName, svcAny := range servicesAny {
				s := ParseService(svcName, svcAny)
				conf.Services[svcName] = &s
			}
		} else {
			cmd := ParseCmd(k, "", v2)
			conf.AdditionalCommands[k] = &cmd
		}

	}

	for _, v2 := range include {
		c2 := LoadConfigFromPath(confPath + v2)
		conf.MergeConfigs(c2)
	}
	return conf
}

func LoadConfigFromPath(p string) Config {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("optimus")
	v.AddConfigPath(p)

	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("Could not read config at path: %v\n%v", p, err)
	}

	var c map[string]any
	err = v.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	conf := ParseConfig(c, p)
	return conf
}

func (c1 *Config) MergeConfigs(c2 Config) {
	if c1.Global == nil && c2.Global != nil {
		c1.Global = c2.Global
	}

	if c1.E2eTests == nil && c2.E2eTests != nil {
		c1.E2eTests = c2.E2eTests
	}

	for k, s := range c2.Services {
		c1.Services[k] = s
	}

	// if c1.Global{}
}

func DefaultConfig() Config {
	return Config{
		Global:             &Global{ShellCmd: "bash -c"},
		BuildCmd:           nil,
		PushCmd:            nil,
		E2eTests:           nil,
		Services:           make(map[string]*Service),
		AdditionalCommands: make(map[string]*Cmd),
	}
}
