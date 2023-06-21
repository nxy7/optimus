package config

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

type Service struct {
	Name string

	// in the future I'd like to be automatically get service version from package.json/cargo.toml and other files, so it can be used to tag images
	// Version            string
	// Language           string
	Root               string
	Start              *Cmd
	Build              *Cmd
	Dev                *Cmd
	PostDev            *Cmd
	AdditionalCommands map[string]*Cmd
	Test               *TestCmd
	DirHash            string
}

func ParseService(name string, a any) Service {
	// fmt.Println("Parsing service")
	s := Service{
		Name: name,
	}

	amap, o := a.(map[string]any)
	if !o {
		panic("Invalid service format")
	}

	for k, v := range amap {
		if k == "dev" {
			c := ParseCmd(v)
			s.Dev = &c
		} else if k == "build" {
			c := ParseCmd(v)
			s.Build = &c
		} else if k == "root" {
			str := v.(string)
			s.Root = str
		}
	}

	// fmt.Println(s)
	return s
}

func (s *Service) ToCobraCommand() cobra.Command {
	svcCmd := cobra.Command{
		Use:   s.Name,
		Short: "Commands related to " + s.Name + " service.",
	}
	allCmds := s.AdditionalCommands
	if allCmds == nil {
		allCmds = make(map[string]*Cmd)
	}
	if s.Build != nil {
		allCmds["build"] = s.Build
	}
	if s.Dev != nil {
		allCmds["dev"] = s.Dev
	}
	if s.Test != nil {
		allCmds["test"] = &s.Test.Cmd
	}

	for k, c := range allCmds {
		svcCmd.AddCommand(&cobra.Command{
			Use:   k,
			Short: k + " command",
			Run: func(cmd *cobra.Command, args []string) {
				e := exec.Command("bash", "-c", c.Run)
				e.Stdout = os.Stdout

				e.Run()

			},
		})
	}

	return svcCmd
}
