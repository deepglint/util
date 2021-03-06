package collector

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"strings"
	"syscall"
)

var FSTYPE_IGNORE = map[string]bool{
	"cgroup":     true,
	"debugfs":    true,
	"devtmpfs":   true,
	"rpc_pipefs": true,
	"rootfs":     true,
}

type DeviceUsageStruct struct {
	FsSpec            string
	FsFile            string
	FsVfstype         string
	BlockSize         uint64 `json:"omitempty"`
	BlocksAll         uint64
	BlocksUsed        uint64
	BlocksFree        uint64
	BlocksUsedPercent float64
	BlocksFreePercent float64
	InodesAll         uint64
	InodesUsed        uint64
	InodesFree        uint64
	InodesUsedPercent float64
	InodesFreePercent float64
}

type Disk struct {
	Total uint64
	Used  uint64
	Free  uint64
	Usage float64
}

func GetCurrentDisk() (result Disk, err error) {
	fs := syscall.Statfs_t{}
	err = syscall.Statfs("/", &fs)
	if err != nil {
		return
	}
	result.Total = uint64(fs.Blocks) * 4
	result.Free = uint64(fs.Bfree) * 4
	result.Used = result.Total - result.Free
	result.Usage = float64(result.Used) / float64(result.Total)
	return
}

func BuildDeviceUsage(arr [3]string) (*DeviceUsageStruct, error) {
	// too long to show
	if len(arr[0]) > 15 {
		arr[0] = "/dev/xvda1"
	}
	ret := &DeviceUsageStruct{FsSpec: arr[0], FsFile: arr[1], FsVfstype: arr[2]}

	fs := syscall.Statfs_t{}
	err := syscall.Statfs(arr[1], &fs)
	if err != nil {
		return nil, err
	}

	// blocks
	ret.BlockSize = uint64(fs.Bsize)
	used := fs.Blocks - fs.Bfree
	// TODO：
	// different fs.Frsize in 'linux' and 'mac'
	// type syscall.Statfs_t has no field or method Frsize
	// ret.BlocksAll = uint64(fs.Frsize) * fs.Blocks
	// ret.BlocksUsed = uint64(fs.Frsize) * used
	// ret.BlocksFree = uint64(fs.Frsize) * fs.Bfree

	ret.BlocksAll = fs.Blocks
	ret.BlocksUsed = used
	ret.BlocksFree = fs.Bfree
	if fs.Blocks == 0 {
		ret.BlocksUsedPercent = 100.0
	} else {
		ret.BlocksUsedPercent = float64(float64(used) * 100.0 / float64(fs.Blocks))
	}
	ret.BlocksFreePercent = 100.0 - ret.BlocksUsedPercent

	// inodes
	ret.InodesAll = fs.Files
	ret.InodesFree = fs.Ffree
	ret.InodesUsed = fs.Files - fs.Ffree
	if fs.Files == 0 {
		ret.InodesUsedPercent = 100.0
	} else {
		ret.InodesUsedPercent = float64(float64(ret.InodesUsed) * 100.0 / float64(ret.InodesAll))
	}
	ret.InodesFreePercent = 100.0 - ret.InodesUsedPercent

	return ret, nil
}

func ListMountPointFromMountInfo() ([][3]string, error) {
	contents, err := ioutil.ReadFile("/proc/self/mountinfo")
	if err != nil {
		return nil, err
	}

	ret := make([][3]string, 0)

	reader := bufio.NewReader(bytes.NewBuffer(contents))
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fields := strings.Fields(string(line))

		fs_spec := fields[8]
		fs_file := fields[4]
		fs_vfstype := fields[7]

		if fs_spec == "none" || fs_spec == "nodev" {
			continue
		}
		if FSTYPE_IGNORE[fs_vfstype] || strings.HasPrefix(fs_vfstype, "fuse") {
			continue
		}

		if strings.HasPrefix(fs_file, "/dev") ||
			strings.HasPrefix(fs_file, "/sys") ||
			strings.HasPrefix(fs_file, "/net") ||
			strings.HasPrefix(fs_file, "/misc") ||
			strings.HasPrefix(fs_file, "/proc") ||
			strings.HasPrefix(fs_file, "/lib") {
			continue
		}
		// keep /dev/xxx device with shorter fs_file (remove mount binds)
		if strings.HasPrefix(fs_spec, "/dev") {
			deviceFound := false
			for idx := range ret {
				if ret[idx][0] == fs_spec {
					deviceFound = true
					if len(fs_file) < len(ret[idx][1]) {
						ret[idx][1] = fs_file
					}
					break
				}
			}
			if !deviceFound {
				ret = append(ret, [3]string{fs_spec, fs_file, fs_vfstype})
			}
		} else {
			ret = append(ret, [3]string{fs_spec, fs_file, fs_vfstype})
		}
	}
	return ret, nil
}

func ListMountPoint() ([][3]string, error) {
	contents, err := ioutil.ReadFile("/proc/mounts")
	if err != nil {
		return nil, err
	}

	ret := make([][3]string, 0)

	reader := bufio.NewReader(bytes.NewBuffer(contents))
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fields := strings.Fields(string(line))
		// Docs come from the fstab(5)
		// fs_spec     # Mounted block special device or remote filesystem e.g. /dev/sda1
		// fs_file     # Mount point e.g. /data
		// fs_vfstype  # File system type e.g. ext4
		// fs_mntops   # Mount options
		// fs_freq     # Dump(8) utility flags
		// fs_passno   # Order in which filesystem checks are done at reboot time

		fs_spec := fields[0]
		fs_file := fields[1]
		fs_vfstype := fields[2]

		if fs_spec == "none" || fs_spec == "nodev" {
			continue
		}

		if FSTYPE_IGNORE[fs_vfstype] || strings.HasPrefix(fs_vfstype, "fuse") {
			continue
		}

		if strings.HasPrefix(fs_file, "/dev") ||
			strings.HasPrefix(fs_file, "/sys") ||
			strings.HasPrefix(fs_file, "/net") ||
			strings.HasPrefix(fs_file, "/misc") ||
			strings.HasPrefix(fs_file, "/proc") ||
			strings.HasPrefix(fs_file, "/lib") {
			continue
		}

		// keep /dev/xxx device with shorter fs_file (remove mount binds)
		if strings.HasPrefix(fs_spec, "/dev") {
			deviceFound := false
			for idx := range ret {
				if ret[idx][0] == fs_spec {
					deviceFound = true
					if len(fs_file) < len(ret[idx][1]) {
						ret[idx][1] = fs_file
					}
					break
				}
			}
			if !deviceFound {
				ret = append(ret, [3]string{fs_spec, fs_file, fs_vfstype})
			}
		} else {
			ret = append(ret, [3]string{fs_spec, fs_file, fs_vfstype})
		}
	}
	return ret, nil
}
