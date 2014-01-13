package exec

import (
	"bytes"
	"os/exec"
)

func Command(cmd_name string, args ...string) (results []string, err error) {
	cmd := exec.Command(cmd_name, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return
	}
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			break
		}
		results = append(results, line)
	}
	return
}
