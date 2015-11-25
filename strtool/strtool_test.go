package strtool

import (
	"testing"
)

const (
	str = "hello world, i am a niubi string."
)

func TestSubStrFromTo(t *testing.T) {

	substr := SubStrFromTo(str, 2, 10)

	t.Log(substr)
}

func TestSubStrLength(t *testing.T) {

	substr := SubstrLength(str, 2, 8)

	t.Log(substr)
}
