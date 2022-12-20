package practice_test

import (
	"testing"

	"github.com/k3forx/GoSystemProgramming/ioReader/practice"
)

func TestCreateZipWithStringsReader(t *testing.T) {
	const fileName string = "test.zip"

	if err := practice.CreateZipWithStringsReader(fileName); err != nil {
		t.Errorf("%s returns non-nil err: %+v\n", t.Name(), err)
		return
	}
}
