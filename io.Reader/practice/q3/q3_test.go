package q3_test

import (
	"testing"

	"github.com/k3forx/GoSystemProgramming/ioReader/practice/q3"
)

func TestCreateZipWithStringsReader(t *testing.T) {
	const fileName string = "test.zip"

	if err := q3.CreateZipWithStringsReader(fileName); err != nil {
		t.Errorf("%s returns non-nil err: %+v\n", t.Name(), err)
		return
	}
}
