package quotereader

import (
	"testing"

	"github.com/deeper-x/fortune-cookie/lib/quotereader"
)

func TestQuotereader(t *testing.T) {
	quotereader.GenQuotes()

	var i interface{} = quotereader.QuoteSentences

	_, ok := i.(quotereader.Quotes)

	if ok != true {
		t.Errorf("error on quote reading")
	}

}
