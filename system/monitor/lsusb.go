package monitor

import (
	"github.com/liutong19890905/util/system/exec"
	"regexp"
	"strconv"
	"strings"
)

type USB struct {
	Busnum int
	Devnum int
	Devid  string
	Des    string
}

func GetCurrentUSB() (results []USB, err error) {
	reg := regexp.MustCompile(`[^:]+`)
	var lines []string
	lines, err = exec.Command("lsusb")
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
			busnum int
			devnum int
			devid  string
			des    string
		)
		var tep_des string
		for i := 6; i < len(ft); i++ {
			if i == len(ft) {
				tep_des += ft[i]
			} else {
				tep_des += ft[i] + " "
			}
		}
		des = tep_des
		busnum, err = strconv.Atoi(ft[1])
		devnum, err = strconv.Atoi(reg.FindString(ft[3]))
		devid = ft[5]
		if err != nil {
			continue
		}
		usb := USB{busnum, devnum, devid, des}
		results = append(results, usb)
	}
	return
}
