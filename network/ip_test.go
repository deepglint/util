package network

import (
	"testing"
)

func TestExternalIp(t *testing.T) {
	ip, err := ExternalIP()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ip)
}
