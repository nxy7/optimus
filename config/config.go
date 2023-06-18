package config

import (
	"optimus/utils"

	"github.com/spf13/viper"
)

type Config struct {
	global             Global
	init               Init
	e2e_tests          E2eTests
	purge              Cmd
	services           []Service
	additionalCommands []Cmd
}

type Global struct {
	shell_cmd string
}

type Init struct {
	cmd  string
	file string
}

type E2eTests struct {
	cmd string
}

type Services struct {
}

type Service struct {
	name    string
	start   Cmd
	dev     Cmd
	postDev Cmd
	test    ServiceTestCmd
	dirHash string
}

type Cmd struct {
	cmd   string
	file  string
	shell string
}

type ServiceTestCmd struct {
	cmd       Cmd
	dependsOn []Service
}

func LoadConfig() Config {
	dirPath := utils.ProjectRoot()
	viper.SetConfigType("yaml")
	viper.SetConfigName("optimus")
	viper.AddConfigPath(dirPath)

	viper.SetDefault("global", map[string]string{"shell_cmd": "bash -c"})

	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		panic("Could not marchal config")
	}

	return Config{}
}
