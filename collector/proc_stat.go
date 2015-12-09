package collector

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/deepglint/util/filetool"
	log "github.com/deepglint/util/logtool"
	"github.com/deepglint/util/system/exec"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Proc struct {
	Pid     int
	Name    string
	Cmdline string
	State   string
}

type Process struct {
	PID   int
	CPU   float64
	MEM   float64
	RSS   float64
	Start string
	Time  string
}

func AllProcs() (ps []*Proc, err error) {
	var dirs []string
	dirs, err = filetool.DirsUnder("/proc")
	if err != nil {
		return
	}

	// id dir is a number, it should be a pid. but don't trust it
	dirs_len := len(dirs)
	if dirs_len == 0 {
		return
	}

	var pid int
	var name_state [2]string
	var cmdline string
	for i := 0; i < dirs_len; i++ {
		if pid, err = strconv.Atoi(dirs[i]); err != nil {
			err = nil
			continue
		} else {
			status_file := fmt.Sprintf("/proc/%d/status", pid)
			cmdline_file := fmt.Sprintf("/proc/%d/cmdline", pid)
			if !filetool.IsExist(status_file) || !filetool.IsExist(cmdline_file) {
				continue
			}

			name_state, err = ReadProcStatus(status_file)
			if err != nil {
				log.Error("read %s fail: %s", status_file, err)
				continue
			}

			cmdline, err = filetool.ReadFileToStringNoLn(cmdline_file)
			if err != nil {
				log.Error("read %s fail: %s", cmdline_file, err)
				continue
			}

			p := Proc{Pid: pid, Name: name_state[0], State: name_state[1], Cmdline: cmdline}
			ps = append(ps, &p)
		}
	}

	return
}

func ReadProcStatus(path string) (name_state [2]string, err error) {
	var content []byte
	content, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	reader := bufio.NewReader(bytes.NewBuffer(content))
	name_done := false
	state_done := false
	for {
		bs, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		line := string(bs)

		colonIndex := strings.Index(line, ":")
		if line[0:colonIndex] == "Name" {
			name_state[0] = strings.TrimSpace(line[colonIndex+1:])
			name_done = true
			continue
		}

		if line[0:colonIndex] == "State" {
			name_state[1] = strings.TrimSpace(line[colonIndex+1:])
			state_done = true
			continue
		}

		if name_done && state_done {
			break
		}

	}

	return
}

func GetCurrentProcesses(process_names map[string]int) (result map[string]Process, err error) {
	reg_label := regexp.MustCompile(`[.|/|a-z|A-Z|0-9|_| |-]+`)
	reg_time := regexp.MustCompile(`[a-z|A-Z|0-9|:|-]+`)
	result = make(map[string]Process)
	var lines []string
	lines, err = exec.Command("ps", "aux")
	if err != nil {
		return
	}
	for i := 0; i < len(lines); i++ {
		tokens := strings.Split(lines[i], " ")
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
			rss   float64
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
		rss, err = strconv.ParseFloat(ft[5], 64)
		start = reg_time.FindString(ft[8])
		time = reg_time.FindString(ft[9])
		if err != nil {
			continue
		}
		process := Process{pid, cpu, mem, rss, start, time}
		result[name] = process
	}
	return
}
