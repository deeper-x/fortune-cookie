package utils

import (
	"math/rand"
	"time"
)

// GenRandom todo doc
func GenRandom(totToCount int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(totToCount)
}
