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

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Run 'purge' command, that's typically used to cleanup all project resources",
	Run: func(cmd *cobra.Command, args []string) {
		dirPath := utils.ProjectRoot()
		viper.SetConfigType("yaml")
		viper.SetConfigName("optimus")
		viper.AddConfigPath(dirPath)

		err := viper.ReadInConfig()
		if err != nil {
			println(err)
		}
		init := viper.GetString("purge")
		c := exec.Command("bash", "-c", init)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		err = c.Run()
		if err != nil {
			// println("command failed: ", err)
			println(err.Error())
		}

		// initByLine := strings.Split(init, "\n")
		// for _, line := range initByLine {
		// 	println(line)
		// }

	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// purgeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// purgeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
