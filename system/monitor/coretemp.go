package monitor

import (
	"github.com/deepglint/util/io/file"
	"regexp"
	"strconv"
)

type Coretemp struct {
	CurTemp  int
	MaxTemp  int
	CritTemp int
}

func GetCurrentCoretemps(data_file_path []string) (result map[string]Coretemp, err error) {
	var lines []string
	reg_content := regexp.MustCompile(`[0-9]+`)
	reg_label := regexp.MustCompile(`[0-9|a-z|A-Z| ]+`)
	result = make(map[string]Coretemp)
	for i := 0; i < len(data_file_path); i++ {
		var (
			cur_temp  int
			max_temp  int
			crit_temp int
			label     string
		)
		lines, err = file.ReadLines(data_file_path[i] + "_input")
		cur_temp, err = strconv.Atoi(reg_content.FindString(lines[0]))
		lines, err = file.ReadLines(data_file_path[i] + "_max")
		max_temp, err = strconv.Atoi(reg_content.FindString(lines[0]))
		lines, err = file.ReadLines(data_file_path[i] + "_crit")
		crit_temp, err = strconv.Atoi(reg_content.FindString(lines[0]))
		lines, err = file.ReadLines(data_file_path[i] + "_label")
		label = reg_label.FindString(lines[0])
		if err != nil {
			return
		}
		coretemp := Coretemp{cur_temp, max_temp, crit_temp}
		result[label] = coretemp
	}
	return
}
