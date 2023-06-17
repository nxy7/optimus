package utils

import "os"

func ProjectRoot() string {
	p, _ := os.Getwd()

	return p
}
