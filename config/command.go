package config

import "log"

type Cmd struct {
	Run         string
	Description string
	File        string
	Shell       string
}

func ParseCmd(a any) Cmd {
	command := Cmd{
		Run:         "",
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
	legalFields := map[string]struct{}{"run": {}, "description": {}, "shell": {}, "file": {}}
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
		}
	}

	if command.Run != "" && command.File != "" {
		panic("Command cannot have both 'file' and 'cmd' fields set")
	}

	return command
}

type TestCmd struct {
	Cmd       Cmd
	DependsOn []Service
}
