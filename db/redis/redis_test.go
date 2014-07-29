package redis

import (
	// "github.com/deepglint/util/db/redis"
	"log"
	"testing"
)

func TestRedis(t *testing.T) {
	var err error

	//Connect redis server
	conn, err := NewConn("tcp", 6379)
	if err != nil {
		t.Errorf("%s\n", err.Error())
		return
	}
	t.Logf("Connect test passed.\n")

	//Set a key-value
	err = conn.SetString("TestKey", "TestValue")
	if err != nil {
		t.Errorf("%s\n", err.Error())
		return
	}
	t.Logf("Set test passed.\n")

	//Append key-value
	err = conn.AppendString("TestKey", "~^^~")
	if err != nil {
		t.Errorf("%s\n", err.Error())
		return
	}
	t.Logf("Append test passed.\n")

	//Get key-value
	result, err := conn.GetString("TestKey")
	if err != nil {
		t.Errorf("%s\n", err.Error())
		return
	}
	log.Println(result)
	t.Logf("Get test passed.\n")

	conn.Close()
}
