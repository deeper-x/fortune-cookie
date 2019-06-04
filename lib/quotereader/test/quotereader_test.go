package quotereader

import (
	"testing"

	"github.com/deeper-x/fortune-cookie/lib/quotereader"
)

func TestQuotereader(t *testing.T) {
	quotereader.LoadData()

	var i interface{} = quotereader.Data

	_, ok := i.(quotereader.QContainer)

	if ok != true {
		t.Errorf("error on quote reading")
	}

}
