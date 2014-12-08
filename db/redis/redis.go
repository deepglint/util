package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	conn redis.Conn
}

func NewConn(protocal string, port int) (this Redis, err error) {
	this.conn, err = redis.Dial(protocal, fmt.Sprintf(":%d", port))
	if err != nil {
		return
	}
	return
}

func (this *Redis) GetString(key string) (result string, err error) {
	n, err := this.conn.Do("GET", key)
	if err != nil {
		return
	}
	result = fmt.Sprintf("%s", n)
	return
}

func (this *Redis) SetString(key string, value string) (err error) {
	_, err = this.conn.Do("GETSET", key, value)
	if err != nil {
		return
	}
	return
}

func (this *Redis) AppendString(key string, value string) (err error) {
	_, err = this.conn.Do("APPEND", key, value)
	if err != nil {
		return
	}
	return
}

func (this *Redis) Close() {
	defer this.conn.Close()
}
