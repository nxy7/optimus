package config

import (
	"fmt"
	"testing"
)

func TestRawConfig(t *testing.T) {
	c := LoadRawConfig()
	// t.Logf("config", c)
	// t.Logf(c)
	fmt.Printf("%+v\n", c)
	// fmt.Printf("c: %v\n", c)
	// fmt.Println(c)
	// fmt.Fprintln(os.Stdout, "config: ", c)
}
