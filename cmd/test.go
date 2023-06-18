/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
	},
}

var e2eTestCmd = &cobra.Command{
	Use:   "e2e",
	Short: "Run E2E Tests",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("e2e tests called")
	},
}
var unitTestCmd = &cobra.Command{
	Use:   "unit",
	Short: "Run unit tests for all services",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("e2e tests called")
	},
}
var allTestCmd = &cobra.Command{
	Use:   "all",
	Short: "Run all tests across whole application",
	Long:  `Run all tests starting with unit tests, then integration tests and lastly e2e tests`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("e2e tests called")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.AddCommand(e2eTestCmd)
	testCmd.AddCommand(allTestCmd)
	testCmd.AddCommand(unitTestCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
