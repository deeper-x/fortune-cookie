package main

import (
	"fmt"
	"math/rand"
	"time"

	qr "github.com/deeper-x/fortune-cookie/lib/quotereader"
)

// RandNum todo doc
func RandNum() int {
	rand.Seed(time.Now().UnixNano())

	maxVal := len(qr.QuoteSentences.Quotes)

	return rand.Intn(maxVal)
}

func main() {
	qr.GenQuotes()
	index := RandNum()

	fmt.Println(qr.QuoteSentences.Quotes[index].QuoteText, qr.QuoteSentences.Quotes[index].QuoteAuthor)

	defer qr.QuoteFile.Close()
}
