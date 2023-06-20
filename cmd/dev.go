/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"optimus/config"

	"github.com/spf13/cobra"
)

// devCmd represents the dev command
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Start development of specific service",
	Long:  `Run develop`,
}

func init() {
	for n, t := range AppConfig.Services {
		devCmd.AddCommand(&cobra.Command{
			Use:   n,
			Short: "Start development for " + n,
			Run: func(cmd *cobra.Command, args []string) {
				switch t.(type) {
				case string:
					println("it's string")
					fmt.Println(t)
				case config.Cmd:
					println("it's cmd")
					fmt.Println(t)

				}
				// svc := t.(map[string]any)
				// shellScript := svc["dev"].(string)
				// fmt.Println(shellScript)

				// c := exec.Command("bash", "-c", shellScript)
				// c.Stdout = os.Stdout
				// c.Stderr = os.Stderr
				// err := c.Run()
				// if err != nil {
				// 	fmt.Println(err)
				// }

			},
		})
	}
	rootCmd.AddCommand(devCmd)
}
