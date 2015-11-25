package filetool

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"

	"github.com/deepglint/glog"
)

// SelfDir gets compiled executable file directory
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// get filepath base name
func Basename(file string) string {
	return path.Base(file)
}

// get filepath dir name
func Dir(file string) string {
	return path.Dir(file)
}

func InsureDir(path string) error {
	if IsExist(path) {
		return nil
	}
	return os.MkdirAll(path, os.ModePerm)
}

func Ext(file string) string {
	return path.Ext(file)
}

// Search a file in paths.
// this is often used in search config file in /etc ~/
func SearchFile(filename string, paths ...string) (fullPath string, err error) {
	for _, path := range paths {
		if fullPath = filepath.Join(path, filename); IsExist(fullPath) {
			return
		}
	}
	err = errors.New(fullPath + " not found in paths")
	return
}

// get absolute filepath, based on built executable file
func RealPath(file string) (string, error) {
	if path.IsAbs(file) {
		return file, nil
	}
	wd, err := os.Getwd()
	return path.Join(wd, file), err
}

// list dirs under dirPath
func DirsUnder(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, nil
	}

	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	ret := []string{}
	for i := 0; i < sz; i++ {
		if fs[i].IsDir() {
			name := fs[i].Name()
			if name != "." && name != ".." {
				ret = append(ret, name)
			}
		}
	}

	return ret, nil

}

// list files under dirPath
func FilesUnder(dirPath string) ([]string, error) {
	if !IsExist(dirPath) {
		return []string{}, nil
	}

	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return []string{}, err
	}

	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}

	ret := []string{}
	for i := 0; i < sz; i++ {
		if !fs[i].IsDir() {
			ret = append(ret, fs[i].Name())
		}
	}

	return ret, nil

}

/*
	read file
*/
func LoadFileByte(path string) ([]byte, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		glog.Errorf("No such file: %s", path)
		return nil, err
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		glog.Errorf("Can not read file: %s", path)
		return nil, err
	}
	return content, nil
}

func LoadFileString(path string) (string, error) {
	content, err := LoadFileByte(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

/*
	write string to file
*/
func ResetString2File(path, content string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0664)
	if err != nil {
		glog.Errorf("Can not open file: %s", path)
		return err
	}
	defer file.Close()

	file.WriteString(content)
	return nil
}

// SelfPath gets compiled executable file absolute path
func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// get file modified time
func FileMTime(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// get file size as how many bytes
func FileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// delete file
func Unlink(file string) error {
	return os.Remove(file)
}

// rename file name
func Rename(file string, to string) error {
	return os.Rename(file, to)
}

// put string to file
func FilePutContent(file string, content string) (int, error) {
	fs, e := os.Create(file)
	if e != nil {
		return 0, e
	}
	defer fs.Close()
	return fs.WriteString(content)
}

// get string from text file
func FileGetContent(file string) (string, error) {
	if !IsFile(file) {
		return "", os.ErrNotExist
	}
	b, e := ioutil.ReadFile(file)
	if e != nil {
		return "", e
	}
	return string(b), nil
}

// it returns false when it's a directory or does not exist.
func IsFile(file string) bool {
	f, e := os.Stat(file)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

func CreateDirRecursively(path string, mode os.FileMode) (err error) {
	dir := filepath.Dir(path)
	if dir != path {
		err = CreateDirRecursively(dir, mode)
		if err != nil {
			return
		}
	}

	err = CreateDir(path, mode)
	return
}

//create dir
func CreateDir(path string, mode os.FileMode) (err error) {
	if err = os.MkdirAll(path, mode); err != nil {
		if os.IsPermission(err) {
			fmt.Println("No permissions.")
		}
	}
	return
}

//create file
func CreateFile(file string, mode interface{}) (string, error) {
	// var f *os.File
	var err error
	if !IsExist(file) {
		_, err = os.Create(file)
		if err != nil {
			return "", err
		}

	}
	err = os.Chmod(file, 0664)
	if err != nil {
		return "", err
	}
	return file, nil
}

type FileRepos []Repository

type Repository struct {
	Name     string
	FileTime int64
}

func (r FileRepos) Len() int {
	return len(r)
}

func (r FileRepos) Less(i, j int) bool {
	return r[i].FileTime < r[j].FileTime
}

func (r FileRepos) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// 获取所有文件
//如果文件达到最上限，按时间删除
func delFile(files []os.FileInfo, count int, fileDir string) {
	if len(files) <= count {
		return
	}

	result := new(FileRepos)

	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			*result = append(*result, Repository{Name: file.Name(), FileTime: file.ModTime().Unix()})
		}
	}

	sort.Sort(result)
	deleteNum := len(files) - count
	for k, v := range *result {
		if k+1 > deleteNum {
			break
		}
		Unlink(fileDir + v.Name)
	}

	return
}
