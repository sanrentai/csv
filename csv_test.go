package csv

import (
	"testing"

	"fmt"
)

func TestCsv(t *testing.T) {
	csv, err := New("F:\\test\\test.csv", "GBK")
	if err != nil {
		t.Error(err)
	}
	s, err := csv.Find(0, 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(s)
}
