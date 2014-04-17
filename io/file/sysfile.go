package file

import (
	"os"
	"path/filepath"
	"time"
)

const (
	IsDirectory = iota
	IsRegular
	IsSymlink
)

type SysFile struct {
	fType  int
	fName  string
	fLink  string
	fSize  int64
	fMtime time.Time
	fPerm  os.FileMode
}

type fs struct {
	files []*SysFile
}

func (self *fs) visit(path string, f os.FileInfo, err error) error {
	if f == nil {
		return err
	}
	var tp int
	if f.IsDir() {
		tp = IsDirectory
	} else if (f.Mode() & os.ModeSymlink) > 0 {
		tp = IsSymlink
	} else {
		tp = IsRegular
	}
	inoFile := &SysFile{
		fName:  path,
		fType:  tp,
		fPerm:  f.Mode(),
		fMtime: f.ModTime(),
		fSize:  f.Size(),
	}
	self.files = append(self.files, inoFile)
	return nil
}

func Visit(root string) []*SysFile {
	self := fs{
		files: make([]*SysFile, 0),
	}
	filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		return self.visit(path, f, err)
	})

	return self.files
}
