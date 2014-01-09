package system

import (
	"github.com/liutong19890905/util/io"
	"regexp"
	"strconv"
	"strings"
)

type CPU struct {
	User        int64
	Nice        int64
	System      int64
	Idle        int64
	Iowait      int64
	Irq         int64
	Softirq     int64
	Stealstolen int64
	Guest_vm    int64
	Guest_lp    int64
	Usage       float64
}

func (this *CPU) Total() int64 {
	return this.User + this.Nice + this.System + this.Idle + this.Iowait +
		this.Irq + this.Softirq + this.Stealstolen + this.Guest_vm + this.Guest_lp
}

func GetCurrentCPUs(data_file_path string) (result map[string]CPU, err error) {
	var lines []string
	lines, err = io.ReadLines(data_file_path)
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
			user        int64
			nice        int64
			system      int64
			idle        int64
			iowait      int64
			irq         int64
			softirq     int64
			stealstolen int64
			guest_vm    int64
			guest_lp    int64
			label       string
		)
		user, err = strconv.ParseInt(infos[1], 10, 64)
		nice, err = strconv.ParseInt(infos[2], 10, 64)
		system, err = strconv.ParseInt(infos[3], 10, 64)
		idle, err = strconv.ParseInt(infos[4], 10, 64)
		iowait, err = strconv.ParseInt(infos[5], 10, 64)
		irq, err = strconv.ParseInt(infos[6], 10, 64)
		softirq, err = strconv.ParseInt(infos[7], 10, 64)
		stealstolen, err = strconv.ParseInt(infos[8], 10, 64)
		guest_vm, err = strconv.ParseInt(infos[9], 10, 64)
		guest_lp, err = strconv.ParseInt(infos[10], 10, 64)
		label = reg_label.FindString(infos[0])
		if err != nil {
			return
		}
		cpu := CPU{user, nice, system, idle, iowait, irq, softirq, stealstolen, guest_vm, guest_lp, 0}
		result[label] = cpu
	}
	return
}
