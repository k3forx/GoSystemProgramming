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
	defer old.Close()
	new, err := os.Open("new.txt")
	if err != nil {
		panic(err)
	}
	defer new.Close()
	io.Copy(new, old)
}
