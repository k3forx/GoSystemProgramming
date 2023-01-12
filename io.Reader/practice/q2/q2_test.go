package q2_test

import (
	"os"
	"testing"

	"github.com/k3forx/GoSystemProgramming/io.Reader/practice/q2"
)

func TestCreateRandFile(t *testing.T) {
	const fileName string = "rand.txt"

	if err := q2.CreateRandFile(fileName); err != nil {
		t.Errorf("err should be nil, but got %+v\n", err)
		return
	}

	f, err := os.Open(fileName)
	if err != nil {
		t.Errorf("err should be nil, but got %+v\n", err)
		return
	}
	defer f.Close()

	stats, err := f.Stat()
	if err != nil {
		t.Errorf("failed to get file info: %+v\n", err)
		return
	}

	expectedSize := 1024
	if stats.Size() != int64(expectedSize) {
		t.Errorf("file size is wrong. want: %d, but got: %d\n", expectedSize, stats.Size())
	}
}
