package monitor

import (
	"errors"
	"github.com/liutong19890905/util/io/file"
	"regexp"
	"strconv"
)

type Memory struct {
	MemTotal   uint64
	MemFree    uint64
	MemBuffers uint64
	MemCached  uint64
	MemUse     uint64
	MemUsage   float64
}

func GetCurrentMemory(data_file_path string) (result Memory, err error) {
	var lines []string
	lines, err = file.ReadLines(data_file_path)
	if err != nil {
		return
	}
	if len(lines) != 42 {
		err = errors.New("There has something wrong with " + data_file_path)
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
