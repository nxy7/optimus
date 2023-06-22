package cmd

import (
	"testing"
)

func TestConfigCanBeMarchaled(t *testing.T) {
	// Execute()
	printConfigCmd.Run(nil, []string{""})
}
