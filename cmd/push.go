/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"optimus/cache"
	"os"

	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Run 'push' command of all services",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			panic("test command only accepts up to 1 argument")
		}

		ca, err := cache.LoadCache()
		if err != nil {
			panic(err)
		}

		services := AppConfig.Services
		errors := RunServicesCommand(services, "push")

		err = ca.SaveCache()
		if err != nil {
			log.Fatal(err)
		}
		if len(errors) > 0 {
			for _, err := range errors {
				fmt.Println(err)
			}

			log.Println("Not all push commands were successful")
			os.Exit(1)
		} else {
			log.Println("All pushes were successful")
		}
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
