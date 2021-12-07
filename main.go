package gocsv

import (
	"bufio"
	"io"
	"strings"
)

type CSV struct {
	Rows []*CSVRow
}

type CSVRow struct {
	Cols []string
}

func Load(input io.Reader, delim rune) *CSV {
	rd := bufio.NewReader(input)
	sb := strings.Builder{}
	rw := &CSVRow{make([]string, 0, 8)}
	csv := &CSV{Rows: make([]*CSVRow, 0, 8)}
	esc := false
	for {
		ch, _, err := rd.ReadRune()
		if err != nil {
			break
		}
		if esc {
			esc = false
			sb.WriteRune(ch)
			continue
		}
		if ch == '\\' {
			esc = true
			continue
		}
		if ch == '\n' {
			rw.Cols = append(rw.Cols, sb.String())
			csv.Rows = append(csv.Rows, rw)
			rw = &CSVRow{make([]string, 0, 8)}
			sb.Reset()
			continue
		}
		if ch == delim {
			rw.Cols = append(rw.Cols, sb.String())
			sb.Reset()
			continue
		}
		sb.WriteRune(ch)
	}
	if sb.Len() > 0 {
		rw.Cols = append(rw.Cols, sb.String())
		csv.Rows = append(csv.Rows, rw)
	}
	return csv
}
