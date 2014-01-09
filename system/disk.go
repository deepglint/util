package system

import (
	"syscall"
)

type Disk struct {
	Total uint64
	Used  uint64
	Free  uint64
	Usage float64
}

func GetCurrentDisk(path string) (result Disk, err error) {
	fs := syscall.Statfs_t{}
	err = syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	result.Total = uint64(fs.Blocks) * 4
	result.Free = uint64(fs.Bfree) * 4
	result.Used = result.Total - result.Free
	result.Usage = float64(result.Used) / float64(result.Total)
	return
}
