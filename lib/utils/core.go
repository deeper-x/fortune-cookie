package utils

import (
	"math/rand"
	"time"

	"github.com/deeper-x/fortune-cookie/lib/quotereader"
)

// RandNum todo doc
func RandNum(toCount []quotereader.Quote) int {
	rand.Seed(time.Now().UnixNano())

	maxVal := len(toCount)

	return rand.Intn(maxVal)
}
