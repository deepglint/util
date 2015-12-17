package ping

import (
	"log"
	"testing"
)

func TestPing(t *testing.T) {
	flag := Ping("115.28.62.181", 10)
	log.Println(flag)
}
