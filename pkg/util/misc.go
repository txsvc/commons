package util

import (
	"math/rand"
)

// Random returns a random number between 0 .. max
func Random(max int) int {
	return rand.Intn(max)
}

// RandomPlusMinus returns a random number between -max .. max
func RandomPlusMinus(max int) int {
	p := rand.Intn(10)
	if p < 6 {
		return rand.Intn(max)
	}
	return 0 - rand.Intn(max)
}

// ValueWithDefault returns the value of a default if empty
func ValueWithDefault(value, def string) string {
	if value == "" {
		return value
	}
	return def
}
