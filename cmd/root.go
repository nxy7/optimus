/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"optimus/config"
)

var AppConfig config.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "optimus",
	Short: "Opinionated monorepo management framework",
	Long:  `Optimus is opinionated and extensible monorepo framework that can work with most web app workflows: standalone services, docker-compose projects and kubernetes clusters. It's easy to extend optimus to do most monorepo tasks using your favourite shell language.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// parse config
	// config := config.Config {}
	AppConfig = config.LoadConfig()
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
