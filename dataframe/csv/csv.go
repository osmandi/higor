package csv

import (
	"encoding/csv"
	"log"
	"strings"
)

type CSV struct {
	Sep        rune
	LineString string
	LazyQuotes bool
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