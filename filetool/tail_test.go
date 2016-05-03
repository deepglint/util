package filetool

import (
	"testing"
)

var file = "/Users/Cheng/100m.txt"
var n int64 = 20

func TestNewFileTail(t *testing.T) {
	fileTail, err := NewFileTail(file, n)
	if err != nil {
		t.Log(err)
	}
	defer fileTail.Close()
	t.Log(fileTail)
	for line := range fileTail.Tail.Lines {
		t.Log(line.Text)
	}

	go func() {
		for {
			select {
			case line := <-fileTail.Tail.Lines:
				t.Log(line.Text)
			}
		}
	}()

	select {}
}
