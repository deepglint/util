package monitor

import (
	"github.com/deepglint/util/io/file"
	"regexp"
	"strconv"
	"strings"
)

type CPU struct {
	User        uint64
	Nice        uint64
	System      uint64
	Idle        uint64
	Iowait      uint64
	Irq         uint64
	Softirq     uint64
	Stealstolen uint64
	Guest_vm    uint64
	Guest_lp    uint64
	Usage       float64
}

func (this *CPU) Total() uint64 {
	return this.User + this.Nice + this.System + this.Idle + this.Iowait +
		this.Irq + this.Softirq + this.Stealstolen + this.Guest_vm + this.Guest_lp
}

func GetCurrentCPUs(data_file_path string) (result map[string]CPU, err error) {
	var lines []string
	lines, err = file.ReadLines(data_file_path)
	if err != nil {
		return
	}
	reg_content := regexp.MustCompile(`[^ ]+`)
	reg_label := regexp.MustCompile(`[0-9|a-z|A-Z| ]+`)
	result = make(map[string]CPU)
	for i := 0; i < len(lines); i++ {
		if !strings.Contains(lines[i], "cpu") {
			break
		}
		infos := reg_content.FindAllString(lines[i], -1)
		var (
			user        uint64
			nice        uint64
			system      uint64
			idle        uint64
			iowait      uint64
			irq         uint64
			softirq     uint64
			stealstolen uint64
			guest_vm    uint64
			guest_lp    uint64
			label       string
		)
		user, err = strconv.ParseUint(infos[1], 10, 64)
		nice, err = strconv.ParseUint(infos[2], 10, 64)
		system, err = strconv.ParseUint(infos[3], 10, 64)
		idle, err = strconv.ParseUint(infos[4], 10, 64)
		iowait, err = strconv.ParseUint(infos[5], 10, 64)
		irq, err = strconv.ParseUint(infos[6], 10, 64)
		softirq, err = strconv.ParseUint(infos[7], 10, 64)
		stealstolen, err = strconv.ParseUint(infos[8], 10, 64)
		guest_vm, err = strconv.ParseUint(infos[9], 10, 64)
		guest_lp, err = strconv.ParseUint(infos[10], 10, 64)
		label = reg_label.FindString(infos[0])
		if err != nil {
			return
		}
		cpu := CPU{user, nice, system, idle, iowait, irq, softirq, stealstolen, guest_vm, guest_lp, 0}
		result[label] = cpu
	}
	return
}
