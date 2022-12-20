package q2

import (
	"encoding/csv"
	"os"
)

func CreateCSV(fileName string, records [][]string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	if err := csvWriter.WriteAll(records); err != nil {
		return err
	}
	csvWriter.Flush()
	return nil
}
