package config

import (
	"fmt"
	"optimus/utils"

	"github.com/spf13/viper"
)

type Config struct {
	Global Global `mapstructure:"global"`
	// Include            []string `mapstructure:"include"`
	// Init               any
	// E2eTests           any
	// Purge              Cmd
	Services map[string]any
	// AdditionalCommands []Cmd
}

type Global struct {
	ShellCmd string `mapstructure:"shell_cmd"`
}

type Init struct {
	Cmd  string
	File string
}

type E2eTests struct {
	Cmd string
}

type Services struct {
}

type Service struct {
	Name    string
	Start   Cmd
	Dev     Cmd
	PostDev Cmd
	Test    ServiceTestCmd
	DirHash string
}

type Cmd struct {
	Cmd   string
	File  string
	Shell string
}

type ServiceTestCmd struct {
	Cmd       Cmd
	DependsOn []Service
}

func LoadConfig() Config {
	dirPath := utils.ProjectRoot()
	viper.SetConfigType("yaml")
	viper.SetConfigName("optimus")
	viper.AddConfigPath(dirPath)

	// viper.SetDefault("global", map[string]string{"shell_cmd": "bash -c"})
	err := viper.ReadInConfig()
	if err != nil {
		panic("Could not read config")
	}

	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println(err)
		panic("Could not marchal config")
	}

	return c
}
