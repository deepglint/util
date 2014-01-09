package system

import (
	"errors"
	"github.com/liutong19890905/util/io"
	"regexp"
	"strconv"
)

type Memory struct {
	MemTotal   int64
	MemFree    int64
	MemBuffers int64
	MemCached  int64
	MemUse     int64
	MemUsage   float64
}

func GetCurrentMemory(data_file_path string) (result Memory, err error) {
	var lines []string
	lines, err = io.ReadLines(data_file_path)
	if err != nil {
		return
	}
	if len(lines) != 42 {
		err = errors.New("There has something wrong with " + data_file_path)
	}
	reg_content := regexp.MustCompile(`[0-9]+`)
	var (
		mem_total   int64
		mem_free    int64
		mem_buffers int64
		mem_cached  int64
		mem_use     int64
		mem_usage   float64
	)
	mem_total, err = strconv.ParseInt(reg_content.FindString(lines[0]), 10, 64)
	mem_free, err = strconv.ParseInt(reg_content.FindString(lines[1]), 10, 64)
	mem_buffers, err = strconv.ParseInt(reg_content.FindString(lines[2]), 10, 64)
	mem_cached, err = strconv.ParseInt(reg_content.FindString(lines[3]), 10, 64)
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
