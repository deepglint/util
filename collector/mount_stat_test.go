package collector

import (
	"testing"
)

var mps map[string]string

func TestListMountPartations(t *testing.T) {
	var err error
	mps, err = ListMountPartations()
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(mps)
}

func TestIsMounted(t *testing.T) {
	flag := IsMounted(mps, "/")

	t.Log(flag)
}

func TestListMountPartationsFromFile(t *testing.T) {
	content, err := ListMountPartationsFromFile("/proc/self/mountinfo")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(content)
	}
}
