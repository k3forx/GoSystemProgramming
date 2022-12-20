package practice

import (
	"archive/zip"
	"log"
	"os"
	"strings"
)

func CreateZipWithStringsReader(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	reader := strings.NewReader("Create zip file with strings.Reader!")
	zipReader, err := zip.NewReader(reader, int64(1024))
	if err != nil {
		log.Print(err)
		return err
	}
	if err := writer.Copy(zipReader.File[0]); err != nil {
		return err
	}

	return nil
}
