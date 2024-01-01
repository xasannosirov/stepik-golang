package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
)

func main() {
	buf := bytes.NewBuffer(nil)
	w := csv.NewWriter(buf)

	for i := 1; i <= 3; i++ {
		val1 := fmt.Sprintf("row %d col 1", i)
		val2 := fmt.Sprintf("row %d col 2", i)
		val3 := fmt.Sprintf("row %d col 3", i)
		if err := w.Write([]string{val1, val2, val3}); err != nil {
			fmt.Println(err)
			return
		}
	}
	w.Flush()

	w.WriteAll([][]string{
		{"row 4 col 1", "row 4 col 2", "row 4 col 3"},
		{"row 5 col 1", "row 5 col 2", "row 5 col 3"},
	})

	r := csv.NewReader(buf)

	for i := 1; i <= 2; i++ {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		}
		fmt.Println(row)
	}
	data, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range data {
		fmt.Println(row)
	}
}
