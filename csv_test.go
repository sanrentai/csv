package csv

import (
	"testing"

	"fmt"
)

func TestCsv(t *testing.T) {
	csv, err := New("test.csv", "GBK")
	if err != nil {
		t.Error(err)
	}
	s, err := csv.Find(0, 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(s)
	fmt.Println(csv.String(1, 0))
	fmt.Println(csv.Float(1, 1))
	fmt.Println(csv.Int(1, 2))
}
