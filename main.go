package main

import (
	"fmt"

	qr "github.com/deeper-x/fortune-cookie/lib/quotereader"
	"github.com/deeper-x/fortune-cookie/lib/utils"
)

func main() {
	qr.LoadData()

	totIn := len(qr.Data.Records)
	index := utils.GenRandom(totIn)

	quote := qr.Data.Records[index].QuoteText
	author := qr.Data.Records[index].QuoteAuthor

	fmt.Println(quote, author)
}
