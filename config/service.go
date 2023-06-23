package config

import (
	"optimus/dirhash"
	"optimus/utils"
	"os"

	"strings"

	"github.com/spf13/cobra"
)

type Service struct {
	Name string

	// in the future I'd like to be automatically get service version from package.json/cargo.toml and other files, so it can be used to tag images
	// Version            string
	// Language           string

	// Path that holds microservice
	Root     string
	Commands map[string]*Cmd

	// Hash of whole directory
	DirHash string
}

func (s *Service) UpdateDirhash() {
	path := strings.Replace(s.Root, "./", "/", 1)
	path = utils.ProjectRoot() + path
	hash, err := dirhash.HashDir(path, make([]string, 0))
	if err != nil {
		panic(err)
	}
	s.DirHash = hash
}

func ParseService(name string, a any, configPath string) Service {
	// fmt.Println("Parsing service")
	s := Service{
		Name:     name,
		Root:     configPath + string(os.PathSeparator) + name,
		Commands: make(map[string]*Cmd, 0),
	}

	amap, o := a.(map[string]any)
	if !o {
		panic("Invalid service format")
	}

	r, o := amap["root"]
	if o {
		rootVal := r.(string)
		rootVal = strings.Replace(rootVal, "./", "", 1)
		s.Root = configPath + string(os.PathSeparator) + rootVal
	}

	for k, v := range amap {
		if k == "dev" {
			c := ParseCmd(k, s.Root, &s, v)
			s.Commands["dev"] = &c
		} else if k == "build" {
			c := ParseCmd(k, s.Root, &s, v)
			s.Commands["build"] = &c
		} else if k == "test" {
			c := ParseCmd(k, s.Root, &s, v)
			s.Commands["test"] = &c
		}
	}

	// serviceHash :=
	return s
}

func (s *Service) ToCobraCommand() cobra.Command {
	svcCmd := cobra.Command{
		Use:   s.Name,
		Short: "Commands related to " + s.Name + " service.",
	}

	for _, c := range s.Commands {
		cobraCmd := c.ToCobraCommand()
		svcCmd.AddCommand(&cobraCmd)
	}

	return svcCmd
}
