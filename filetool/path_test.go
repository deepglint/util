package filetool

import (
	"testing"
)

const (
	root = "../../utils"
)

func TestFilesOfPath(t *testing.T) {
	FilesOfPath(root, true)

	for _, v := range FileStrings {
		t.Log(v)
	}

	ResetFileStrings()
	FilesOfPath(root, false)

	for _, v := range FileStrings {
		t.Log(v)
	}
}

// func TestFileInfosOfPath(t *testing.T) {
// 	FileInfosOfPath(root)
// 	for _, v := range FileInfos {
// 		t.Log(v.Name())
// 	}

// 	ResetFileInfos()

// 	FileInfosOfPath(root)
// 	for _, v := range FileInfos {
// 		t.Log("-----" + v.Name())
// 	}
// }
