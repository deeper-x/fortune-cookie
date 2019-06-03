package quotereader

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var QuoteSentences Quotes
var quoteFile *os.File

// Quotes todo doc
type Quotes struct {
	Quotes []Quote `json:"quotes"`
}

// Quote doc doc
type Quote struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
}

// GenQuotes todo doc
func GenQuotes() {
	quoteFile, _ = os.Open("assets/quotes.json")

	byteVal, _ := ioutil.ReadAll(quoteFile)
	json.Unmarshal(byteVal, &QuoteSentences)
	defer quoteFile.Close()
}
