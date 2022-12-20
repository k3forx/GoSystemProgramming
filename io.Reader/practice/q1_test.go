package practice_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/k3forx/GoSystemProgramming/ioReader/practice"
)

func TestCopyFile(t *testing.T) {
	setup(t)

	if err := practice.CopyFile(); err != nil {
		t.Errorf("err should be nil, but got %+v\n", err)
	}

	file, err := os.Open("new.txt")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()

	buf := make([]byte, 1024)
	if _, err := file.Read(buf); err != nil {
		t.Error(err)
		return
	}

	expected := "this is content of old.txt"
	if diff := cmp.Diff(expected, string(buf)); err != nil {
		t.Errorf("%s result mismatch (-want, +got):\n%s", t.Name(), diff)
	}
}

func setup(t *testing.T) {
	t.Helper()
	reader := bytes.NewBufferString("")
	new, err := os.Open("new.txt")
	if err != nil {
		t.Fatal(err)
	}
	io.Copy(new, reader)
}
