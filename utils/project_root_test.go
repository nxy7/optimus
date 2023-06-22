package utils

import (
	"os"
	"testing"
)

func TestReturnsPath(t *testing.T) {
	p := ProjectRoot()
	_, err := os.ReadDir(p)
	if err != nil {
		t.Error(err)
	}
}
