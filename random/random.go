package random

import (
	"math/rand"
	"time"
)

func RandSeed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func RandInt64(min int64, max int64) int64 {
	return min + rand.Int63n(max-min)
}

func RandFloat32() float32 {
	var base int
	base = 10000000
	return float32(RandInt(0, base)) / float32(base)
}

func RandFloat64() float64 {
	var base int
	base = 10000000
	return float64(RandInt(0, base)) / float64(base)
}

func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}

func RandomKey(l int) string {
	bytes := make([]byte, l)
	var tmp int
	for i := 0; i < l; i++ {
		for {
			tmp = RandInt(48, 122)
			if (tmp > 57 && tmp < 65) || (tmp > 90 && tmp < 97) {
				continue
			}
			break
		}
		bytes[i] = byte(tmp)
	}
	return string(bytes)
}
