/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"optimus/cache"
	"optimus/config"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Run 'build' command for all services inside the project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			panic("test command only accepts up to 1 argument")
		}

		ca, err := cache.LoadCache()
		if err != nil {
			panic(err)
		}

		services := AppConfig.Services
		errors := RunServicesCommand(services, "build")

		err = ca.SaveCache()
		if err != nil {
			log.Fatal(err)
		}
		if len(errors) > 0 {
			for _, err := range errors {
				fmt.Println(err)
			}

			log.Println("Not all builds exited succesfully")
			os.Exit(1)
		} else {
			log.Println("All builds were successful")
		}
	},
}

func RunServicesCommand(services map[string]*config.Service, commandName string) []error {
	ca, err := cache.LoadCache()
	if err != nil {
		panic(err)
	}
	errors := make([]error, 0)
	var wg sync.WaitGroup
	for _, s := range services {
		sCommand := s.Commands[commandName]
		if sCommand == nil {
			continue
		}
		cachedRes := sCommand.GetCmdCache(ca)
		if cachedRes != nil {
			if string(sCommand.DirHash) == string(cachedRes.Hash) && cachedRes.RanSuccessfully {
				fmt.Printf("Command %v is in cache\n", sCommand.ParentService.Name)
				continue
				// fmt.Printf("Cmd:\n%+v\nCached:\n%+v\n", sTestCmd, cachedRes)
			} else {
				fmt.Printf("\nDifferent dirhash for %v\nCmd: %v\nCache: %v\n\n", sCommand.ParentService.Name, sCommand.DirHash, cachedRes.Hash)
			}

		}
		if sCommand != nil && s != nil {
			wg.Add(1)

			testCmd := sCommand.CommandFunc
			if testCmd == nil {
				testCmd = sCommand.GenerateCommandFunc()
			}
			go func() {
				err := testCmd()
				if err != nil {
					errors = append(errors, err)
				}
				cached := sCommand.ToCmdCache()
				cached.UpdateCache(ca)
				wg.Done()
			}()

		}
	}
	wg.Wait()
	return errors
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
