package file

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	IsDirectory = iota
	IsRegular
	IsSymlink
)

type SysFile struct {
	Type  int
	Name  string
	Link  string
	Size  int64
	Mtime time.Time
	Perm  os.FileMode
}

type fs struct {
	files []SysFile
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
	inoFile := SysFile{
		Name:  path,
		Type:  tp,
		Perm:  f.Mode(),
		Mtime: f.ModTime(),
		Size:  f.Size(),
	}
	self.files = append(self.files, inoFile)
	return nil
}

func Visit(root string) (error, []SysFile) {
	var files []SysFile
	var self fs
	self.files = files
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		return self.visit(path, f, err)
	})

	return err, self.files
}

func IsExist(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}

func Remove(name string) error {
	return os.Remove(name)
}

func Move(oldname string, newname string) error {
	dirs := strings.Split(newname, "/")
	dirname := ""
	delta := 1

	if len(dirs[len(dirs)-1]) == 0 {
		delta = 2
	}

	for i := 0; i < len(dirs)-delta-1; i++ {
		dirname += dirs[i] + "/"
	}
	dirname += dirs[len(dirs)-delta-1]

	if !IsExist(dirname) {
		Mkdir(dirname)
	}

	return os.Rename(oldname, newname)
}

func Mkdir(name string) error {
	err := os.Mkdir(name, 0700)
	if err != nil && strings.Contains(err.Error(), "no such file or directory") {
		dirs := strings.Split(name, "/")
		newname := ""
		delta := 1

		if len(dirs[len(dirs)-1]) == 0 {
			delta = 2
		}

		for i := 0; i < len(dirs)-delta-1; i++ {
			newname += dirs[i] + "/"
		}
		newname += dirs[len(dirs)-delta-1]
		inner_err := Mkdir(newname)
		if inner_err == nil {
			err = os.Mkdir(name, 0700)
		} else {
			return inner_err
		}
	}
	return err
}
