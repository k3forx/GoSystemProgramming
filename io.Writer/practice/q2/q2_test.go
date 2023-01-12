package q2_test

import (
	"testing"

	"github.com/k3forx/GoSystemProgramming/io.Writer/practice/q2"
)

func TestCreateCSV(t *testing.T) {
	const fileName = "test.csv"
	records := [][]string{
		{"userId", "name", "age"},
		{"1", "john", "12"},
		{"2", "mike", "34"},
	}

	if err := q2.CreateCSV(fileName, records); err != nil {
		t.Errorf("CreateCSV returns non-nil err: %+v\n", err)
		return
	}
}
