package network

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strings"
)

func GetDNS() ([]string, error) {
	contents, err := ioutil.ReadFile("/etc/resolv.conf")
	if err != nil {
		return nil, err
	}
	var nameservers []string
	reader := bufio.NewReader(bytes.NewBuffer(contents))
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if strings.HasPrefix(string(line), "#") {
			continue
		}

		fields := strings.Fields(string(line))
		if len(fields) != 2 {
			continue
		}

		if fields[0] != "nameserver" {
			continue
		}

		if fields[1] == "127.0.0.1" {
			continue
		}

		nameservers = append(nameservers, fields[1])
	}
	return nameservers, nil
}
