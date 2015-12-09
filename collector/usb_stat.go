package collector

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/deepglint/util/systool"
)

type USB struct {
	Busnum int
	Devnum int
	Devid  string
	Des    string
}

func GetCurrentUSB() (results map[string]USB, err error) {
	results = make(map[string]USB)
	reg_label := regexp.MustCompile(`[0-9|a-z|A-Z|.|/|-|:|\[|\]|_|+| ]+`)
	reg := regexp.MustCompile(`[^:]+`)
	var lines []string
	out, err := systool.CmdOut("lsusb")
	if err != nil {
		return
	}
	lines = strings.Split(out, "\n")
	for i := 0; i < len(lines); i++ {
		tokens := strings.Split(lines[i], " ")
		ft := make([]string, 0)
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}
		if len(ft) < 6 {
			continue
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
		devid = reg_label.FindString(ft[5])
		if err != nil {
			continue
		}
		usb := USB{busnum, devnum, devid, des}
		results[devid] = usb
	}
	return
}
