package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

func ProjectRoot() string {
	p, _ := os.Getwd()
	separator := string(os.PathSeparator)

	paths := strings.Split(p, separator)

	p = (func() string {
		for i := len(paths); i > 1; i-- {
			searchPath := paths[0:i]
			searchPathStr := strings.Join(searchPath, separator)

			files, err := ioutil.ReadDir(searchPathStr)
			if err != nil {
				panic(err)
			}

			for _, fi := range files {
				if fi.IsDir() && fi.Name() == ".git" {
					return searchPathStr
				}
			}
		}
		return ""
	})()
	if p == "" {
		panic("Not a git repository")
	}

	return p
}
