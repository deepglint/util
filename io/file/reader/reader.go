package file

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func ReadLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)

	if file, err = os.Open(path); err != nil {
		return
	}

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 1024))

	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	file.Close()
	log.Println("...")
	return
}

func ReadBytes(path string) (data []byte, size int, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)

	if file, err = os.Open(path); err != nil {
		return
	}

	data, err = ioutil.ReadAll(file)
	size = len(data)
	return
}
