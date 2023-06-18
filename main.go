/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"optimus/cmd"
	"optimus/start"
	"optimus/utils"
)

func main() {
	start.Start(utils.ProjectRoot())

	cmd.Execute()
}
