package system

import (
	"github.com/liutong19890905/util/io"
	"regexp"
	"strconv"
)

type CurrentCoretempState struct {
	Cores map[string]Coretemp
}

type Coretemp struct {
	CurTemp  int
	MaxTemp  int
	CritTemp int
}

func GetCurrentCoretempState(data_file_path []string) (coretemp_state CurrentCoretempState, err error) {
	var lines []string
	reg := regexp.MustCompile(`[0-9]+`)
	coretemp_state.Cores = make(map[string]Coretemp)
	for i := 0; i < len(data_file_path); i++ {
		var (
			cur_temp  int
			max_temp  int
			crit_temp int
			label     string
		)
		lines, err = io.ReadLines(data_file_path[i] + "_input")
		cur_temp, err = strconv.Atoi(reg.FindString(lines[0]))
		lines, err = io.ReadLines(data_file_path[i] + "_max")
		max_temp, err = strconv.Atoi(reg.FindString(lines[0]))
		lines, err = io.ReadLines(data_file_path[i] + "_crit")
		crit_temp, err = strconv.Atoi(reg.FindString(lines[0]))
		reg = regexp.MustCompile(`[0-9|a-z|A-Z| ]+`)
		lines, err = io.ReadLines(data_file_path[i] + "_label")
		label = reg.FindString(lines[0])
		if err != nil {
			return
		}
		coretemp := Coretemp{cur_temp, max_temp, crit_temp}
		coretemp_state.Cores[label] = coretemp
		//fmt.Printf("label:%s, CurTemp: %d, MaxTemp:%d, CritTemp:%d\n", label, cur_temp, max_temp, crit_temp)
	}
	return
}
