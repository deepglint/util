package pad

import (
	"testing"
)

const (
	str = "1234567890"
)

func TestparseString(t *testing.T) {
	var s string = "1"
	t.Log(parseString(s))
	var i32 int32 = 1
	t.Log(parseString(i32))
	var ui64 uint64 = 1
	t.Log(parseString(ui64))
}

func TestLeftPad(t *testing.T) {
	ret := LeftPad(str, "0", 15)
	t.Log(ret)
}

func TestRightPad(t *testing.T) {
	ret := RightPad(str, "1", 15)
	t.Log(ret)
}

func TestLeftPad2Len(t *testing.T) {
	ret := LeftPad2Len(str, "#", 15)
	t.Log(ret)
}

func TestRightPad2Len(t *testing.T) {
	ret := RightPad2Len(str, "*", 15)
	t.Log(ret)
}
