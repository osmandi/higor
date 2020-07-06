package higor

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Default values
var (
	separator rune = ','
)

type ReadCSVOptions struct {
	Separator rune
	Filename  string
	Header    bool
}

type df [][]string

// Shape
func (d df) Shape() [2]int {

	return [2]int{len(d), len(d[0])}
}

// ReadCSV load CSV and print
func ReadCSV(op ReadCSVOptions) df {

	// Default options
	switch {
	case op.Separator == 0:
		op.Separator = separator
	}

	csvFile, err := os.Open(op.Filename)

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	// Print each line
	csvLines := csv.NewReader(csvFile)

	// Set custom options
	csvLines.Comma = op.Separator

	// Read CSV
	records, err := csvLines.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	return records

}
