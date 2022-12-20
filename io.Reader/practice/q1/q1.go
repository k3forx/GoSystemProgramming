package q1

import (
	"io"
	"os"
)

func CopyFile() error {
	old, err := os.Open("old.txt")
	if err != nil {
		return err
	}
	defer old.Close()

	new, err := os.Create("new.txt")
	if err != nil {
		return err
	}
	defer new.Close()

	if _, err := io.Copy(new, old); err != nil {
		return err
	}

	return nil
}
