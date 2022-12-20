package practice_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/k3forx/GoSystemProgramming/ioReader/practice"
)

func TestCopyFile(t *testing.T) {
	setup(t)

	practice.CopyFile()

	file, err := os.Open("new.txt")
	if err != nil {
		t.Error(err)
		return
	}
	var buf []byte
	if _, err := file.Read(buf); err != nil {
		t.Error(err)
		return
	}
	expected := "this is content of old.txt"
	if string(buf) != expected {
		t.Errorf("content is expected as %s, but got %s", expected, string(buf))
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
