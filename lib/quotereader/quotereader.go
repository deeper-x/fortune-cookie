package quotereader

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Data is the container of records: we write on it everything we read in JSON
var Data QContainer
var quoteFile *os.File

// QContainer defines a root container, with all quotes+authors
type QContainer struct {
	Records []Quote `json:"records"`
}

// Quote define each node in root json
type Quote struct {
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
}

// LoadData generates the records' list
func LoadData() {
	quoteFile, _ = os.Open("assets/quotes.json")

	byteVal, _ := ioutil.ReadAll(quoteFile)
	json.Unmarshal(byteVal, &Data)

	defer quoteFile.Close()
}
