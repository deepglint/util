package collector

import (
	"strings"

	// "github.com/golang/glog"
	"github.com/deepglint/util/filetool"
	"github.com/deepglint/util/system/exec"
)

// only used for host
func ListMountPartations() (map[string]string, error) {

	ret := make(map[string]string)

	mps, err := exec.Command("findmnt")
	if err != nil {
		return ret, err
	}
	for _, mp := range mps {

		mpvs := strings.Fields(mp)

		if len(mpvs) == 4 && strings.HasPrefix(mpvs[1], "/") {
			ret[mpvs[0]] = mpvs[1]
		}
	}

	return ret, nil
}

// used mountinfo for docker
func ListMountPartationsFromFile(path string) (mountpoints, mountformats map[string]string, err error) {
	//
	mountpoints = make(map[string]string)
	mountformats = make(map[string]string)

	content, err := filetool.ReadFileToStringNoLn(path)
	if err != nil {
		return
	}

	ars := strings.Fields(content)
	for i := 0; i < len(ars)/10; i++ {
		mountpoints[ars[10*i+4]] = ars[10*i+8]
		mountformats[ars[10*i+8]] = ars[10*i+7]
	}

	return
}

// used mounts for docker
func ListMountPartationsFromMounts(path string) (mountpoints map[string]string, err error) {
	//
	mountpoints = make(map[string]string)

	content, err := filetool.ReadFileToStringNoLn(path)
	if err != nil {
		return
	}

	ars := strings.Fields(content)
	for i := 0; i < len(ars)/6; i++ {
		mountpoints[ars[6*i+1]] = ars[6*i]
	}

	return
}
