package random

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	RandSeed()
	fmt.Printf("random int: %d\n", RandInt(0, 1000))
	fmt.Printf("random float32: %f\n", RandFloat32())
	fmt.Printf("random float64: %f\n", RandFloat64())
	fmt.Printf("random string: %s\n", RandomString(10))
}
