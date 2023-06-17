/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"optimus/cmd"
	"optimus/utils"
)

func main() {
	start(utils.ProjectRoot())

	cmd.Execute()
}
