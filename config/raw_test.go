package config

import (
	"fmt"
	"testing"
)

func TestRawConfig(t *testing.T) {
	c := LoadRawConfig()
	fmt.Printf("%+v\n", c)
}
