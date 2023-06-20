/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// printConfigCmd represents the printConfig command
var printConfigCmd = &cobra.Command{
	Use:     "print-config",
	Aliases: []string{"pc"},
	Short:   "Print merged Optimus Configuration",
	Run: func(cmd *cobra.Command, args []string) {
		jsonified, err := json.MarshalIndent(AppConfig, "", "   ")
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v", string(jsonified))
	},
}

func init() {
	rootCmd.AddCommand(printConfigCmd)
}
