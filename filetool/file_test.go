package filetool

import (
	"testing"
)

func TestRename(t *testing.T) {
	source := "/Users/Cheng/a"
	target := "/Users/Cheng/b"
	Rename(source, target)
}
