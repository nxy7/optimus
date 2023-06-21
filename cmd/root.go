/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"os/exec"

	"optimus/config"

	"github.com/spf13/cobra"
)

var AppConfig config.Config

var rootCmd = &cobra.Command{
	Use:   "optimus",
	Short: "Opinionated monorepo management framework",
	Long:  `Optimus is opinionated but extensible monorepo framework that can work with most web app workflows: standalone services, docker-compose projects and kubernetes clusters. It's easy to extend optimus to do most monorepo tasks using your favourite shell language.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	AppConfig = config.LoadConfig()

	for name, command := range AppConfig.AdditionalCommands {
		rootCmd.AddCommand(&cobra.Command{
			Use:   name,
			Short: command.Description,
			Run: func(cmd *cobra.Command, args []string) {
				e := exec.Command("bash", "-c", command.Run)
				e.Stdout = os.Stdout

				e.Run()
			},
		})
	}

	for _, svc := range AppConfig.Services {
		svcCmd := svc.ToCobraCommand()

		rootCmd.AddCommand(&svcCmd)
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
