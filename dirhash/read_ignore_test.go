package dirhash

import (
	"fmt"
	"optimus/utils"
	"testing"
)

func TestReadIgnore(t *testing.T) {
	ignores := make([]string, 0)
	err := ReadIgnore(utils.ProjectRoot()+"/.gitignore", &ignores)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(ignores)
}
