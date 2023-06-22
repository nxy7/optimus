package config

import (
	"encoding/json"
	"fmt"
	"log"

	// "optimus/utils"

	"os"
	"os/exec"

	// "strings"

	"github.com/spf13/cobra"
)

type Cmd struct {
	Run         string
	Path        string
	Name        string
	Description string
	File        string
	Shell       string
	CommandFunc func() error
}

func (c Cmd) MarshalJSON() ([]byte, error) {
	fmt.Println("elo")
	return json.Marshal(&struct {
		Run         string
		Path        string
		Name        string
		Description string
		File        string
		Shell       string
		CommandFunc string
	}{
		Run:         c.Run,
		Path:        c.Path,
		Name:        c.Name,
		Description: c.Description,
		File:        c.File,
		Shell:       c.Shell,
		CommandFunc: "Function",
	})
}

func ParseCmd(name string, root string, a any) Cmd {
	command := Cmd{
		Run:         "",
		Name:        name,
		Path:        root,
		Description: "",
		File:        "",
		Shell:       "bash -c",
	}
	s, ok := a.(string)
	if ok {
		command.Run = s
		return command
	}

	obj, ok := a.(map[string]any)
	legalFields := map[string]struct{}{"run": {}, "description": {}, "shell": {}, "file": {}, "root": {}}
	if !ok {
		panic("Invalid Cmd shape")
	}

	for k, v := range obj {
		_, ok := legalFields[k]
		if !ok {
			log.Panicf("%v is not a legal command field", k)
		}

		vStr := v.(string)
		if k == "run" {
			command.Run = vStr
		} else if k == "description" {
			command.Description = vStr
		} else if k == "shell" {
			command.Shell = vStr
		} else if k == "file" {
			command.File = vStr
		} else if k == "root" {
			command.Path = vStr
		}
	}

	if command.Run != "" && command.File != "" {
		panic("Command cannot have both 'file' and 'cmd' fields set")
	}

	command.CommandFunc = command.GenerateCommandFunc()

	return command
}

func (c *Cmd) ToCobraCommand() cobra.Command {
	return cobra.Command{
		Use:   c.Name,
		Short: c.Description,
		Run: func(cmd *cobra.Command, args []string) {
			if c.CommandFunc == nil {
				c.CommandFunc = c.GenerateCommandFunc()
			}
			err := c.CommandFunc()
			if err != nil {
				fmt.Println("Command failed")
				fmt.Println(c)
				panic(err)
			}
		},
	}
}

func (c *Cmd) GenerateCommandFunc() func() error {
	return func() error {
		e := exec.Command("bash", "-c", c.Run)
		e.Dir = c.Path
		e.Stdout = os.Stdout
		e.Stderr = os.Stderr
		e.Stdin = os.Stdin

		err := e.Run()
		return err
	}
}

type TestCmd struct {
	Cmd       Cmd
	DependsOn []Service
}
