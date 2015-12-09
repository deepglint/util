package collector

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/deepglint/util/filetool"
	log "github.com/deepglint/util/logtool"
	"github.com/deepglint/util/systool"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type LoginStruct struct {
	Username string
	Port     string
	From     string
	Latest   string
}

func LastLogin() ([]*LoginStruct, error) {
	var logins []*LoginStruct = make([]*LoginStruct, 0)
	bs, err := systool.CmdOutBytes("lastlog")
	if err != nil {
		log.Error("Error exec lastlog: %s", err)
		return nil, err
	}

	reader := bufio.NewReader(bytes.NewBuffer(bs))
	// ignore the first line
	line, _, err := reader.ReadLine()
	if err == io.EOF || err != nil {
		return nil, err
	}
	for {
		line, _, err = reader.ReadLine()
		if err == io.EOF || err != nil {
			break
		}

		lineStr := string(line)
		regex := regexp.MustCompile("\\s{2,}")
		splitline := regex.Split(lineStr, -1)
		if len(splitline) != 4 {
			continue
		}

		login := &LoginStruct{splitline[0], splitline[1], splitline[2], strings.Split(splitline[3], "+")[0]}
		logins = append(logins, login)
	}
	return logins, nil
}

func SystemDate() (string, error) {
	return systool.CmdOutNoLn("/bin/date")
}

func SystemUptime() (arr [3]int64, err error) {
	var content string
	content, err = filetool.ReadFileToStringNoLn("/proc/uptime")
	if err != nil {
		return
	}

	fields := strings.Fields(content)
	if len(fields) < 2 {
		err = errors.New("/proc/uptime parse error")
		return
	}

	secStr := fields[0]
	var secF float64
	secF, err = strconv.ParseFloat(secStr, 64)
	if err != nil {
		return
	}

	minTotal := secF / 60.0
	hourTotal := minTotal / 60.0

	days := int64(hourTotal / 24.0)
	hours := int64(hourTotal) - days*24
	mins := int64(minTotal) - (days * 60 * 24) - (hours * 60)

	return [3]int64{days, hours, mins}, nil
}
