/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"optimus/cmd"
	"optimus/utils"
)

func main() {
	println(utils.ProjectRoot())
	cmd.Execute()
}
