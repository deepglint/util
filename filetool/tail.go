package filetool

import (
	"os"

	"github.com/hpcloud/tail"
)

type FileTail struct {
	file string
	Tail *tail.Tail
}

func NewFileTail(file string, n int64) (*FileTail, error) {
	var err error
	fileTail := &FileTail{
		file: file,
	}

	config := tail.Config{Follow: true}
	config.Location = &tail.SeekInfo{-n, os.SEEK_END}

	fileTail.Tail, err = tail.TailFile(file, config)

	return fileTail, err
}

func (this *FileTail) Close() {
	this.Tail.Cleanup()
	this.Tail.Stop()
}
