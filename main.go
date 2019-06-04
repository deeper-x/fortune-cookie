package main

import (
	"fmt"

	qr "github.com/deeper-x/fortune-cookie/lib/quotereader"
	"github.com/deeper-x/fortune-cookie/lib/utils"
)

func main() {
	qr.GenQuotes()
	totIn := len(qr.QuoteSentences.Quotes)

	index := utils.RandNum(totIn)

	fmt.Println(qr.QuoteSentences.Quotes[index].QuoteText,
		qr.QuoteSentences.Quotes[index].QuoteAuthor)

}
