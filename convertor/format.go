package convertor

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/deepglint/util/filetool"
	"github.com/golang/glog"
)

func FormatUglyFile2INI(sourceFile, targetFile string) (err error) {
	if !filetool.IsExist(sourceFile) {
		msg := fmt.Sprintf("file %s not found", sourceFile)
		return errors.New(msg)
	}

	contents, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(contents)
	reader := bufio.NewReader(buf)

	targetBuf := new(bytes.Buffer)

	var lineIni, lineStr string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			targetBuf.WriteString("\n")
			break
		}
		if len(line) == 0 {
			continue
		}
		lineStr = strings.TrimSpace(string(line))
		if strings.HasPrefix(lineStr, ";") || strings.HasPrefix(lineStr, "#") {
			continue
		}
		fields := strings.Fields(lineStr)
		switch len(fields) {
		case 0:
			continue
		case 1:
			if strings.HasPrefix(fields[0], "[") {
				lineIni = fmt.Sprintf("%s\n", fields[0])
				targetBuf.WriteString(lineIni)
			} else {
				lineIni = fmt.Sprintf("%s=\n", fields[0])
				targetBuf.WriteString(lineIni)
			}
		case 2:
			lineIni = fmt.Sprintf("%s = %s\n", fields[0], fields[1])
			targetBuf.WriteString(lineIni)
		default:
			var value string
			for i := 1; i < len(fields); i++ {
				value = fmt.Sprintf("%s %s", value, fields[i])
			}
			lineIni = fmt.Sprintf("%s = %s", fields[0], value)
			targetBuf.WriteString(lineIni)
		}
	}

	// glog.Infoln(targetBuf.String())

	if filetool.IsExist(targetFile) {
		os.Remove(targetFile)
	}
	_, err = filetool.CreateFile(targetFile, os.FileMode(0755))
	if err != nil {
		glog.Errorln(err)
		return err
	}

	_, err = filetool.WriteBytesToFile(targetFile, targetBuf.Bytes())

	return
}

func FormatINIFile2Ugly(sourceFile, targetFile string) (err error) {
	if !filetool.IsExist(sourceFile) {
		msg := fmt.Sprintf("file %s not found", sourceFile)
		return errors.New(msg)
	}

	contents, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(contents)
	reader := bufio.NewReader(buf)

	targetBuf := new(bytes.Buffer)

	var lineIni, lineStr string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			targetBuf.WriteString("\n")
			break
		}
		if len(line) == 0 {
			continue
		}
		lineStr = strings.TrimSpace(string(line))
		if strings.HasPrefix(lineStr, ";") || strings.HasPrefix(lineStr, "#") {
			continue
		}
		fields := strings.SplitN(lineStr, "=", 2)
		switch len(fields) {
		case 0:
			continue
		case 1:
			lineIni = fmt.Sprintf("%s\n", fields[0])
			targetBuf.WriteString(lineIni)
		case 2:
			lineIni = fmt.Sprintf("%s\r%s\n", fields[0], fields[1])
			targetBuf.WriteString(lineIni)
		}
	}

	if filetool.IsExist(targetFile) {
		os.Remove(targetFile)
	}
	_, err = filetool.CreateFile(targetFile, os.FileMode(0755))
	if err != nil {
		glog.Errorln(err)
		return err
	}

	_, err = filetool.WriteBytesToFile(targetFile, targetBuf.Bytes())

	return
}
