package cache

import (
	"testing"
)

func TestMakeCache(t *testing.T) {
	c, err := LoadCache()
	if err != nil {
		t.Log(err)
	}
	if c == nil {
		t.Error("Cache returned here should never be nil")
	}
}
