package systool

import (
	"testing"
	"time"
)

func TestCmdOut(t *testing.T) {
	ret, err := CmdOut("cmd.sh", "1")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(ret)

	ret, err = CmdOut("cmd.sh")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(err)
}

func TestCmdOutNoLn(t *testing.T) {
	ret, err := CmdOutNoLn("cmd.sh", "1")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(ret)
}

func TestCmdRunWithTimeout(t *testing.T) {
	out, err, done := CmdRunWithTimeout(time.Duration(1)*time.Second, "cmd.sh", "1")
	if err != nil {
		t.Log(err)
	}
	t.Log(done)
	t.Log(out)

	out, err, done = CmdRunWithTimeout(time.Duration(15)*time.Second, "cmd_1.sh")
	if err != nil {
		t.Log(err)
	}
	t.Log(done)
	t.Log(out)

	out, err, done = CmdRunWithTimeout(time.Duration(15)*time.Second, "cmd.sh", "1")
	if err != nil {
		t.Log(err)
	}
	t.Log(done)
	t.Log(out)
}
