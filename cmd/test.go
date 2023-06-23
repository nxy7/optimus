/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"optimus/cache"
	"os"

	"sync"

	"github.com/spf13/cobra"
	// "google.golang.org/appengine/log"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run test for all services, returns 0 exit code only if all tests pass",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			panic("test command only accepts up to 1 argument")
		}

		ca, err := cache.LoadCache()
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("ca: %v\n", ca)
		}

		services := AppConfig.Services
		errors := make([]error, 0)
		var wg sync.WaitGroup
		for _, s := range services {
			sTestCmd := s.Commands["test"]
			if sTestCmd != nil && s != nil {
				wg.Add(1)
				testCmd := sTestCmd.CommandFunc
				if testCmd == nil {
					testCmd = sTestCmd.GenerateCommandFunc()
				}
				go func() {
					err := testCmd()
					if err != nil {
						errors = append(errors, err)
					}
					wg.Done()
				}()

			}
		}
		wg.Wait()

		for _, c := range AppConfig.FinishedCommands() {
			cc := c.ToCmdCache()
			cc.UpdateCache(ca)
		}

		fmt.Println(ca)
		err = ca.SaveCache()
		if err != nil {
			log.Fatal(err)
		}
		if len(errors) > 0 {
			for _, err := range errors {
				fmt.Println(err)
			}

			log.Println("Not all tests passed, exiting with code 1")
			os.Exit(1)
		} else {
			log.Println("All tests passed")
		}
	},
}

func init() {
	testCmd.Flags().BoolP("ignore-cache", "i", false, "Run all tests without using cached results")
	rootCmd.AddCommand(testCmd)
}
