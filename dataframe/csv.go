package dataframe

import (
	"encoding/csv"
	"log"
	"os"
)

// CSV type
type CSV struct {
	sep           rune
	lineString    string
	lazyQuotes    bool
	removeNewLine bool
}

func errorChecker(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CSVOption alternative parameters
type csvOption func(c *CSV)

// Sep CSV separator in rune type: ',', ';', '|', etc...
func sep(separator rune) csvOption {
	return func(c *CSV) {
		c.sep = separator
	}
}

// LazyQuotes Interpreting internal quotes
func lazyQuotes(lazy bool) csvOption {
	return func(c *CSV) {
		c.lazyQuotes = lazy
	}
}

// RemoveNewLine To remove \n on the string line
func removeNewLine(RemoveNewLine bool) csvOption {
	return func(c *CSV) {
		c.removeNewLine = RemoveNewLine
	}
}

/////////
// CSV /
///////

// ReadCSV to read CSV files and convert it to [][]string
func readCSV(filename string, opts ...csvOption) [][]string {

	csvInternal := &CSV{}
	csvInternal.sep = ','

	for _, opt := range opts {
		opt(csvInternal)
	}

	// Open file
	csvFile, err := os.Open(filename)
	errorChecker(err)
	defer csvFile.Close()

	// Read CSV
	csvReader := csv.NewReader(csvFile)
	csvReader.Comma = csvInternal.sep

	// Convert CSV to [][]string
	csvLines, err := csvReader.ReadAll()
	errorChecker(err)

	return csvLines
}

// ExportCSV To export a dataframe to CSV file
func exportCSV(filename string, data [][]string, opts ...csvOption) {
	csvInternal := &CSV{}
	csvInternal.sep = ','

	for _, opt := range opts {
		opt(csvInternal)
	}

	csvFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	errorChecker(err)
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Comma = csvInternal.sep
	err = csvWriter.WriteAll(data)
	errorChecker(err)

}
