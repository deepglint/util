package collector

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/deepglint/util/filetool"
)

type Memory struct {
	MemBuffers uint64
	MemCached  uint64
	MemTotal   uint64
	MemFree    uint64
	MemUse     uint64
	MemUsage   float64
}

type Meminfo struct {
	Buffers   uint64
	Cached    uint64
	MemTotal  uint64
	MemFree   uint64
	SwapTotal uint64
	SwapUsed  uint64
	SwapFree  uint64
}

var Multi uint64 = 1024

func MemoryInfo() (result Memory, err error) {
	var lines []string
	lines, err = filetool.ReadLines("/proc/meminfo")
	if err != nil {
		return
	}

	reg_content := regexp.MustCompile(`[0-9]+`)
	var (
		mem_total   uint64
		mem_free    uint64
		mem_buffers uint64
		mem_cached  uint64
		mem_use     uint64
		mem_usage   float64
	)
	mem_total, err = strconv.ParseUint(reg_content.FindString(lines[0]), 10, 64)
	mem_free, err = strconv.ParseUint(reg_content.FindString(lines[1]), 10, 64)
	mem_buffers, err = strconv.ParseUint(reg_content.FindString(lines[2]), 10, 64)
	mem_cached, err = strconv.ParseUint(reg_content.FindString(lines[3]), 10, 64)
	mem_use = mem_total - mem_free - mem_buffers - mem_cached
	if mem_total == 0 {
		mem_usage = 0
	} else {
		mem_usage = float64(mem_use) / float64(mem_total)
	}

	if err != nil {
		return
	}
	result = Memory{mem_total, mem_free, mem_buffers, mem_cached, mem_use, mem_usage}
	return
}

func MemInfo() (*Meminfo, error) {
	want := map[string]bool{
		"Buffers:":   true,
		"Cached:":    true,
		"MemTotal:":  true,
		"MemFree:":   true,
		"SwapTotal:": true,
		"SwapFree:":  true}

	contents, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return nil, err
	}

	memInfo := &Meminfo{}

	reader := bufio.NewReader(bytes.NewBuffer(contents))

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fields := strings.Fields(string(line))
		fieldName := fields[0]

		_, ok := want[fieldName]
		if ok && len(fields) == 3 {
			val, numerr := strconv.ParseUint(fields[1], 10, 64)
			if numerr != nil {
				continue
			}
			switch fieldName {
			case "Buffers:":
				memInfo.Buffers = val * Multi
			case "Cached:":
				memInfo.Cached = val * Multi
			case "MemTotal:":
				memInfo.MemTotal = val * Multi
			case "MemFree:":
				memInfo.MemFree = val * Multi
			case "SwapTotal:":
				memInfo.SwapTotal = val * Multi
			case "SwapFree:":
				memInfo.SwapFree = val * Multi
			}
		}
	}
	memInfo.SwapUsed = memInfo.SwapTotal - memInfo.SwapFree

	return memInfo, nil
}
