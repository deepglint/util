package filetool

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	FileStrings []string
	FileInfos   []os.FileInfo
)

/*
	reset fileStrings as it is global
*/
func ResetFileStrings() {
	FileStrings = []string{}
}

/*
	reset fileInfos as it is global
*/
func ResetFileInfos() {
	FileInfos = []os.FileInfo{}
}

func DirsOfPath(root string) []string {
	var ret []string
	// walk through root path
	absroot, _ := filepath.Abs(root)
	err := filepath.Walk(absroot, func(curpath string, info os.FileInfo, err error) error {
		if err != nil || nil == info {
			log.Printf("Error info: %s", err)
			return err
		}

		curdir := path.Dir(curpath)
		if curdir != absroot {
			return nil
		}
		name := path.Base(curpath)
		if strings.HasPrefix(name, ".") {
			return nil
		}

		if info.IsDir() {
			ret = append(ret, info.Name())
		}
		return nil
	})
	if err != nil {
		log.Printf("Error walking: %s", err)
	}
	return ret
}

func FilesOfPath(root string, abs bool) error {
	var err error
	if abs {
		root, _ = filepath.Abs(root)
	}
	// FileStrings = append(FileStrings, root)
	// walk through root path
	err = filepath.Walk(root, func(curpath string, info os.FileInfo, err error) error {
		if err != nil || nil == info {
			log.Printf("Error info: %s", err)
			return err
		}

		curdir := path.Dir(curpath)
		if curdir != root {
			return nil
		}
		name := path.Base(curpath)
		if strings.HasPrefix(name, ".") {
			return nil
		}

		filename := path.Join(curdir, info.Name())
		if info.IsDir() {
			FilesOfPath(filename, abs)
		}
		if info.IsDir() {
			return nil
		}
		FileStrings = append(FileStrings, filename)

		return nil
	})
	if err != nil {
		log.Printf("Error walking: %s", err)
		return err
	}
	return nil
}

/*
	iterate fileInfos
*/
func FileInfosOfPath(root string) error {
	// walk through root path
	absroot, _ := filepath.Abs(root)
	err := filepath.Walk(absroot, func(curpath string, info os.FileInfo, err error) error {
		if err != nil || nil == info {
			log.Printf("Error info: %s", err)
			return err
		}

		curdir := path.Dir(curpath)
		if curdir != absroot {
			return nil
		}
		name := path.Base(curpath)
		if strings.HasPrefix(name, ".") {
			return nil
		}

		if info.IsDir() {
			FileInfosOfPath(path.Join(curdir, info.Name()))
		}
		if info.IsDir() {
			return nil
		}
		FileInfos = append(FileInfos, info)
		return nil
	})
	if err != nil {
		log.Printf("Error walking: %s", err)
		return err
	}
	return nil
}
