/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"optimus/config"

	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Start development of specific service",
	Long:  `Run develop`,
}

func init() {
	AppConfig = config.LoadConfig()
	for n, _ := range AppConfig.Services {
		devCmd.AddCommand(&cobra.Command{
			Use:   n,
			Short: "Run dev for " + n,
			Run: func(cmd *cobra.Command, args []string) {
				println("elo")
			},
		})
	}
	rootCmd.AddCommand(devCmd)
}
