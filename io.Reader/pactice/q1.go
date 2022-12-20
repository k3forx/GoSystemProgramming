package practice

import (
	"io"
	"os"
)

func CopyFile() {
	old, err := os.Open("old.txt")
	if err != nil {
		panic(err)
	}
	new, err := os.Open("new.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(new, old)
}
