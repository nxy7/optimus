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
		} else if ca == nil {
			panic("cache empty?")
		}

		services := AppConfig.Services
		errors := make([]error, 0)
		var wg sync.WaitGroup
		for _, s := range services {
			sTestCmd := s.Commands["test"]
			cachedRes := sTestCmd.GetCmdCache(ca)
			if cachedRes != nil {
				if string(sTestCmd.DirHash) == string(cachedRes.Hash) {
					fmt.Printf("Command %v is in cache\n", sTestCmd.ParentService.Name)
					continue
					// fmt.Printf("Cmd:\n%+v\nCached:\n%+v\n", sTestCmd, cachedRes)
				} else {
					fmt.Printf("\nDifferent dirhash for %v\nCmd: %v\nCache: %v\n\n", sTestCmd.ParentService.Name, sTestCmd.DirHash, cachedRes.Hash)
				}

			}
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
					cached := sTestCmd.ToCmdCache()
					cached.UpdateCache(ca)
					wg.Done()
				}()

			}
		}
		wg.Wait()

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
	testCmd.Flags().BoolP("only-services", "s", false, "Only run tests for services")
	testCmd.Flags().StringP("filter", "f", "", "Phrase that tests should be filtered by")
	rootCmd.AddCommand(testCmd)
}
