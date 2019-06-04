package utils

import (
	"testing"

	"github.com/deeper-x/fortune-cookie/lib/utils"
)

func TestRand(t *testing.T) {
	a := utils.RandNum(5)
	var emptyInterf interface{} = a
	_, ok := emptyInterf.(int)

	if ok != true {
		t.Errorf("Random number must return an integer")
	}
}
