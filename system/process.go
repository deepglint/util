package system

import (
	"bytes"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type Process struct {
	PID   int
	CPU   float64
	MEM   float64
	RSS   uint64
	Start string
	Time  string
}

func GetCurrentProcesses(process_names map[string]int) (result map[string]Process, err error) {
	reg_label := regexp.MustCompile(`[0-9|a-z|A-Z|.|/|-|:|\[|\]|_|+| ]+`)
	result = make(map[string]Process)
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return
	}
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			break
		}
		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}
		var (
			name  string
			pid   int
			cpu   float64
			mem   float64
			rss   uint64
			start string
			time  string
		)
		var tep_name string
		for i := 10; i < len(ft); i++ {
			if i == len(ft) {
				tep_name += ft[i]
			} else {
				tep_name += ft[i] + " "
			}
		}
		name = reg_label.FindString(tep_name)
		if process_names[name] != 1 {
			continue
		}
		pid, err = strconv.Atoi(ft[1])
		cpu, err = strconv.ParseFloat(ft[2], 64)
		mem, err = strconv.ParseFloat(ft[3], 64)
		rss, err = strconv.ParseUint(ft[5], 10, 64)
		start = reg_label.FindString(ft[8])
		time = reg_label.FindString(ft[9])
		if err != nil {
			continue
		}
		process := Process{pid, cpu, mem, rss, start, time}
		result[name] = process
	}
	return
}
