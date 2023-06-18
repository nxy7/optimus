/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"optimus/utils"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dirPath := utils.ProjectRoot()
		viper.SetConfigType("yaml")
		viper.SetConfigName("optimus")
		viper.AddConfigPath(dirPath)

		err := viper.ReadInConfig()
		if err != nil {
			println(err)
		}
		init := viper.GetString("init")
		c := exec.Command("bash", "-c", init)
		c.Stdout = os.Stdout
		err = c.Run()
		if err != nil {
			println(err)
		}

		// initByLine := strings.Split(init, "\n")
		// for _, line := range initByLine {
		// 	println(line)
		// }

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
