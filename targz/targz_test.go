package targz

import (
	"testing"
)

const (
	file   = "cur.tar.gz"
	path   = "../../utils"
	target = "data"
)

func TestCompressTargz(t *testing.T) {
	CompressTargz(file, path)
}

func TestUncompressTargz(t *testing.T) {
	UncompressTargz(file, target)
}
