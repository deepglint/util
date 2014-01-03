package security

import (
	"encoding/base64"
)

var coder = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

func EncodeBase64(src string) string {
	return coder.EncodeToString([]byte(src))
}

func DecodeBase64(src string) (result string, err error) {
	var result_bytes []byte
	result_bytes, err = coder.DecodeString(src)
	result = string(result_bytes)
	return
}
