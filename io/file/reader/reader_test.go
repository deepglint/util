package file

import (
	"fmt"
	"log"
	"testing"
)

func TestReadLines(t *testing.T) {
	results, err := ReadLines("readme.txt")
	if err != nil {
		t.Errorf("%s\n", err.Error())
		return
	}
	for i, result := range results {
		log.Println("[" + fmt.Sprintf("%d", i) + "] " + result)
	}
}
