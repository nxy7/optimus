package config

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	c := LoadConfig()
	j, err := json.Marshal(c)
	if err != nil {
		panic("could not jsonify")
	}
	fmt.Printf("c: %v\n", string(j))
}
