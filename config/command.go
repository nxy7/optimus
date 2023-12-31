package config

import (
	"encoding/json"
	"fmt"
	"log"
	"optimus/cache"
	"optimus/dirhash"

	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

type Cmd struct {
	Run string
	// Parent service can be empty in case of top level commands
	ParentService *Service
	// Root of the services used to run commands from
	Root string
	// If specified CacheRoot is used to cache result of command
	CacheRoot string

	DirHash             []byte
	Cache               bool
	DidRun              bool
	DidExitSuccessfully bool
	Name                string
	Description         string
	File                string
	Shell               string
	CommandFunc         func() error
}

func (c Cmd) ToCmdCache() cache.CommandCache {
	parentName := ""
	if c.ParentService != nil {
		parentName = c.ParentService.Name
	}
	return cache.CommandCache{
		Name:            c.Name,
		Output:          "",
		ParentService:   parentName,
		RanSuccessfully: c.DidExitSuccessfully,
		Hash:            c.DirHash,
	}
}

func (c Cmd) GetCmdCache(ca *cache.Cache) *cache.CommandCache {
	for _, cc := range ca.Commands {
		parentName := ""
		if c.ParentService != nil {
			parentName = c.ParentService.Name
		}
		if cc.Name == c.Name && cc.ParentService == parentName {
			return cc
		}
	}

	return nil
}

func (c Cmd) MarshalJSON() ([]byte, error) {
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
		Path:        c.Root,
		Name:        c.Name,
		Description: c.Description,
		File:        c.File,
		Shell:       c.Shell,
		CommandFunc: "Function",
	})
}

func ParseCmd(name string, root string, parentService *Service, a any) Cmd {
	command := Cmd{
		Run:           "",
		Name:          name,
		Root:          root,
		ParentService: parentService,
		Description:   "",
		File:          "",
		Shell:         "bash -c",
	}
	if parentService != nil && parentService.DirHash != nil {
		command.DirHash = parentService.DirHash
	} else {
		dh, err := dirhash.HashDir(root, dirhash.DefaultIgnoredPaths())
		if err != nil {
			panic(err)
		} else {
			command.DirHash = dh
		}
	}
	s, ok := a.(string)
	if ok {
		command.Run = s
		return command
	}

	obj, ok := a.(map[string]any)
	legalFields := map[string]struct{}{"run": {}, "description": {}, "shell": {}, "file": {}, "root": {}, "cache": {}, "cacheRoot": {}}
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
			command.Root = vStr
		} else if k == "cache" {
			command.Cache = vStr == "true"
		} else if k == "cacheRoot" {
			command.CacheRoot = vStr
		}
	}

	if command.Run != "" && command.File != "" {
		panic("Command cannot have both 'file' and 'cmd' fields set")
	}

	command.CommandFunc = command.GenerateCommandFunc()

	return command
}

func (c *Cmd) ToCobraCommand() cobra.Command {
	cobraCmd := cobra.Command{
		Use:   c.Name,
		Short: c.Description,
		Run: func(cmd *cobra.Command, args []string) {
			if c.CommandFunc == nil {
				c.CommandFunc = c.GenerateCommandFunc()
			}

			force := cmd.Flag("force").Value.String() == "true"
			if !force {
				// check checksum from lockfile and if it's the same as now then skip test
			}

			err := c.CommandFunc()
			if err != nil {
				fmt.Println("Command failed")
				fmt.Println(c)
				panic(err)
			}
		},
	}
	cobraCmd.Flags().BoolP("force", "f", false, "usage string")
	return cobraCmd
}

func (c *Cmd) GenerateCommandFunc() func() error {
	return func() error {
		e := exec.Command("bash", "-c", c.Run)
		e.Dir = c.Root
		e.Stdout = os.Stdout
		e.Stderr = os.Stderr
		e.Stdin = os.Stdin

		err := e.Run()
		c.DidRun = true
		if err == nil {
			c.DidExitSuccessfully = true
		}
		return err
	}
}

type TestCmd struct {
	Cmd       Cmd
	DependsOn []Service
}
