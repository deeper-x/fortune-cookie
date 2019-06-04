package utils

import (
	"math/rand"
	"time"
)

// RandNum todo doc
func RandNum(totToCount int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(totToCount)
}
