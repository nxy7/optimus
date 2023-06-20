/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// printConfigCmd represents the printConfig command
var printConfigCmd = &cobra.Command{
	Use:     "print-config",
	Aliases: []string{"pc"},
	Short:   "Print merged Optimus Configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Printing config")
		fmt.Printf("AppConfig: %+v\n", AppConfig)
	},
}

func init() {
	rootCmd.AddCommand(printConfigCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printConfigCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printConfigCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
