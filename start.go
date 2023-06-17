package main

import "github.com/spf13/viper"

func start(dirPath string) {
	viper.SetConfigFile("optimus")
	viper.AddConfigPath(dirPath)

	err := viper.ReadInConfig()
	if err != nil {
		println(err)
	}
	z := viper.Get("services")
	println(z)
}
