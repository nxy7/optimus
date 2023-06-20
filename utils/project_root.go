package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ProjectRoot() string {
	p, _ := os.Getwd()
	separator := string(os.PathSeparator)

	paths := strings.Split(p, separator)
	for i := len(paths); i > 1; i-- {
		searchPath := paths[0:i]

		files, err := ioutil.ReadDir(separator + strings.Join(searchPath, separator))
		if err != nil {
			panic(err)
		}

		fmt.Println(searchPath)
		fmt.Println(files)
	}

	return p
}
