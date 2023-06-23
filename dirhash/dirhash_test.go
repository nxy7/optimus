package dirhash

import (
	"fmt"
	"optimus/utils"
	"testing"
)

func TestDirHash(t *testing.T) {
	s, err := HashDir(utils.ProjectRoot(), DefaultIgnoredPaths())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(s)
}
