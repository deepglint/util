package network

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

func ExternalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

func EnvIP() (string, error) {
	ip := os.Getenv("IP")
	if len(ip) == 0 {
		return ip, errors.New("Cannot get IP environment variable.")
	}
	return ip, nil
}

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

		subfields := strings.Split(fields[1], ",")
		for _, subfield := range subfields {
			nameservers = append(nameservers, subfield)
		}
	}
	return nameservers, nil
}
