package gocsv

import (
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	r := strings.NewReader("a,b,c\n1,2,3,a\\,b")
	csv := Load(r, ',')

	firstRow := csv.Rows[0]
	if len(firstRow.Cols) < 3 {
		t.Fatal("First row len should be 3: ", len(firstRow.Cols))
	}

	for i := range firstRow.Cols {
		t.Log(firstRow.Cols[i])
	}

	if len(csv.Rows) != 2 {
		t.Fatal("CSV Rows must be 2: ", len(csv.Rows))
	}
	if csv.Rows[0].Cols[0] != "a" {
		t.Fatal("0,0 Should be 'a': ", csv.Rows[0].Cols[0])
	}
	if csv.Rows[1].Cols[0] != "1" {
		t.Fatal("1,0 Should be '1': ", csv.Rows[1].Cols[0])
	}
	if csv.Rows[1].Cols[3] != "a,b" {
		t.Fatal("1,3 Should be 'a,b': ", csv.Rows[1].Cols[3])
	}
}
