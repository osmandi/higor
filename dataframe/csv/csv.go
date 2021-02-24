package csv

import (
	"encoding/csv"
	"log"
	"strings"
)

type CSV struct {
	Sep           rune
	LineString    string
	LazyQuotes    bool
	RemoveNewLine bool
}

// CSV Options

type CSVOption func(c *CSV)

// Sep CSV separator in rune type: ',', ';', '|', etc...
func Sep(separator rune) CSVOption {
	return func(c *CSV) {
		c.Sep = separator
	}
}

// Line A line
func Line(textLine string) CSVOption {
	return func(c *CSV) {
		c.LineString = textLine
	}
}

// LazyQuotes Interpreting internal quotes
func LazyQuotes(lazy bool) CSVOption {
	return func(c *CSV) {
		c.LazyQuotes = lazy
	}
}

// RemoveNewLine To remove \n on the string line
func RemoveNewLine(RemoveNewLine bool) CSVOption {
	return func(c *CSV) {
		c.RemoveNewLine = RemoveNewLine
	}
}

/////////////////////////////
// Header (Read to String) /
///////////////////////////

// CSVReadHeader Read the first column has a header
func CSVReadHeader(opts ...CSVOption) []string {
	csvInternal := &CSV{}
	csvInternal.Sep = ','
	for _, opt := range opts {
		opt(csvInternal)
	}

	if len(csvInternal.LineString) == 0 {
		return []string{}
	}

	reader := csv.NewReader(strings.NewReader(csvInternal.LineString))
	reader.Comma = csvInternal.Sep
	reader.LazyQuotes = csvInternal.LazyQuotes

	columns, err := reader.Read()

	if err != nil {
		log.Fatal(err)
	}

	return columns

}

//////////////////////////
// Row (Read to String) /
////////////////////////

// CSVReadRowNormal Read a line from CSV and convert it to []string
func CSVReadRowNormal(opts ...CSVOption) []string {
	csvInternal := &CSV{}
	csvInternal.Sep = ','

	for _, opt := range opts {
		opt(csvInternal)
	}

	if len(csvInternal.LineString) == 0 {
		return []string{}
	}

	if csvInternal.RemoveNewLine {
		csvInternal.LineString = strings.Replace(csvInternal.LineString, "\n", "", -1)
	}

	reader := csv.NewReader(strings.NewReader(csvInternal.LineString))
	reader.Comma = csvInternal.Sep
	reader.LazyQuotes = csvInternal.LazyQuotes

	columns, err := reader.Read()

	if err != nil {
		log.Fatal(err)
	}

	return columns
}
