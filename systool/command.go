package systool

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"time"

	log "github.com/deepglint/util/logtool"
	"github.com/deepglint/util/strtool"
)

func CmdOut(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func CmdOutBytes(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.Bytes(), err
}

func CmdOutNoLn(name string, arg ...string) (out string, err error) {
	out, err = CmdOut(name, arg...)
	if err != nil {
		return
	}

	return strtool.TrimRightSpace(string(out)), nil
}

func CmdRunWithTimeout(timeout time.Duration, name string, arg ...string) (string, error, bool) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Start()
	if err != nil {
		return "", err, false
	}

	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-time.After(timeout):
		//timeout
		if err = cmd.Process.Kill(); err != nil {
			log.Error("failed to kill: %s, error: %s", cmd.Path, err)
		}
		go func() {
			<-done // allow goroutine to exit
		}()
		msg := fmt.Sprintf("process:%s killed because of timeout", cmd.Path)
		err = errors.New(msg)
		return "", err, true
	case err = <-done:
		return out.String(), err, false
	}
}
