package start

import "github.com/spf13/viper"

func Start(dirPath string) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("optimus")
	viper.AddConfigPath(dirPath)

	err := viper.ReadInConfig()
	if err != nil {
		println("command failed: ", err)
	}
	z := viper.GetString("start")
	println(z)
}
