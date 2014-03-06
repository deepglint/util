package security

import (
	"math/rand"
	"strconv"
	"time"
)

func SessionID() string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	rndNum := rand.Int63()
	return Md5(Md5(strconv.FormatInt(nano, 10)) + Md5(strconv.FormatInt(rndNum, 10)))
}
