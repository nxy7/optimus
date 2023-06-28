package cmd

import (
	"encoding/json"

	"optimus/config"
	"testing"
)

func TestCommandCanBeMarchaled(t *testing.T) {
	c := &config.Cmd{
		Run:         "",
		Root:        "",
		Name:        "",
		Description: "",
		File:        "",
		Shell:       "",
		CommandFunc: func() error {
			println("elo co tam")
			return nil
		},
	}

	_, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		panic(err)
	}
}
