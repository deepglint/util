package collector

import (
	"testing"
)

var mount [][3]string

func TestGetCurrentDisk(t *testing.T) {
	disk, err := GetCurrentDisk()
	if err != nil {
		t.Log(err)
	} else {
		t.Log(disk)
	}
}

func TestListMountPoint(t *testing.T) {
	var err error
	mount, err = ListMountPoint()
	if err != nil {
		t.Log(err)
	} else {
		t.Log(mount)
	}
}

func TestBuildDeviceUsage(t *testing.T) {
	for _, value := range mount {
		du, err := BuildDeviceUsage(value)
		if err != nil {
			t.Log(err)
		} else {
			t.Log(du)
		}
	}
}
