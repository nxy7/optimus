/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"optimus/config"

	"github.com/spf13/cobra"
)

var AppConfig config.Config

var rootCmd = &cobra.Command{
	Use:   "optimus",
	Short: "Opinionated monorepo management framework",
	Long:  `Optimus is extensible monorepo framework that can work with most web app workflows: standalone services, docker-compose projects and kubernetes clusters. It's easy to extend optimus to do most monorepo tasks using your favourite shell language. Optimus supports caching of command results, so if your project didn't change then we won't rerun the command (unless you want to).`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	AppConfig = config.LoadConfig()

	for _, command := range AppConfig.AdditionalCommands {
		cobraCmd := command.ToCobraCommand()
		rootCmd.AddCommand(&cobraCmd)
	}

	for _, svc := range AppConfig.Services {
		svcCmd := svc.ToCobraCommand()

		rootCmd.AddCommand(&svcCmd)
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
