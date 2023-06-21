package config

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	c := LoadConfig()
	fmt.Printf("c: %v\n", c)
}
