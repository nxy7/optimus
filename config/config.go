package config

import (
	"fmt"
	"log"
	"optimus/utils"

	"github.com/spf13/viper"
)

type Config struct {
	Global             Global
	BuildCmd           Cmd
	PushCmd            Cmd
	E2eTests           Cmd
	Services           map[string]Service
	AdditionalCommands map[string]Cmd
}

type E2eTests struct {
	Cmd string
}

func LoadConfig() Config {
	p := utils.ProjectRoot()
	conf := LoadConfigFromPath(p)
	return conf
}

func LoadConfigFromPath(p string) Config {
	conf := DefaultConfig()
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

	for k, v2 := range c {
		if k == "include" {
			strv2, o := v2.([]any)
			if !o {
				panic("Wrong include format")
			}

			for _, v3 := range strv2 {
				v3str, o := v3.(string)
				if !o {
					panic("include only accepts strings")
				}
				_ = LoadConfigFromPath(v3str)
			}
		}
	}

	return conf
}

func DefaultConfig() Config {
	return Config{
		Global:   Global{ShellCmd: "bash -c"},
		BuildCmd: Cmd{},
		PushCmd:  Cmd{},
		E2eTests: Cmd{},
		Services: map[string]Service{
			"frontend": Service{
				AdditionalCommands: make(map[string]Cmd),
			},
		},
		AdditionalCommands: map[string]Cmd{
			"testcmd": Cmd{
				Description: "Gówno",
				Run:         "echo 2",
			},
		},
	}
}
