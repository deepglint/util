package file

import (
	"log"
	"testing"
)

func TestCreateFileByString(t *testing.T) {
	err := CreateFileByString("test_string.txt", "Hello World!")
	if err != nil {
		t.Errorf("%s\n", err.Error())
		return
	}
	log.Println("Write string: Hello World!")
}

func TestCreateFileByBytes(t *testing.T) {
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!'}
	err := CreateFileByBytes("test_bytes.txt", bytes)
	if err != nil {
		t.Errorf("%s\n", err.Error())
		return
	}
	log.Println("Write bytes: Hello World!")
}
