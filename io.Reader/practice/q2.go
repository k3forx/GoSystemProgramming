package practice

import (
	"crypto/rand"
	"os"
)

func CreateRandFile(fileName string) error {
	buf := make([]byte, 1024)
	if _, err := rand.Reader.Read(buf); err != nil {
		return err
	}

	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(buf); err != nil {
		return err
	}
	return nil
}
