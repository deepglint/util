package io

import (
	"os"
)

func CreateFileByString(filename string, content string) (err error) {
	fout, err := os.Create(filename)
	defer fout.Close()
	if err != nil {
		return
	}
	fout.WriteString(content)
	return
}

func CreateFileByBytes(filename string, content []byte) (err error) {
	fout, err := os.Create(filename)
	defer fout.Close()
	if err != nil {
		return
	}
	fout.Write(content)
	return
}
