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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	AppConfig = config.LoadConfig()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// parse config
	// config := config.Config {}
	AppConfig = config.LoadConfig()

	for name, command := range AppConfig.AdditionalCommands {
		rootCmd.AddCommand(&cobra.Command{
			Use: name,
			Run: func(cmd *cobra.Command, args []string) {
				e := exec.Command("bash", "-c", command.Run)
				e.Stdout = os.Stdout

				e.Run()
			},
		})
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
