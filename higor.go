package higor

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/osmandi/higor/dataframe"
)

var Version string = "v0.2.1"

// HelloHigor Print a simple message to check if Higor are installed correctly
// and print the version installed
func HelloHigor() string {

	helloMessage := fmt.Sprintf("Hello from Higor :) %s", Version)
	return helloMessage
}

/////////
// CSV /
///////

// ReadCSV Read a CSV file and save it as a DataFrame
func ReadCSV(filename string, opts ...dataframe.CSVOption) dataframe.DataFrame {
	csvInternal := &dataframe.CSV{}
	csvInternal.Sep = ','

	for _, opt := range opts {
		opt(csvInternal)
	}

	// Open file
	csvFile, err := os.Open(filename)
	dataframe.ErrorChecker(err)
	defer csvFile.Close()

	// Read CSV
	csvReader := csv.NewReader(csvFile)
	csvReader.Comma = csvInternal.Sep

	// Convert CSV to [][]string
	csv, err := csvReader.ReadAll()
	dataframe.ErrorChecker(err)

	df := dataframe.DataFrame{}
	df.Columns = csv[0]
	chapters := dataframe.Book{}

	for _, rowValue := range csv[1:] {
		for columnIndex, columnValue := range rowValue {
			chapters[df.Columns[columnIndex]] = append(chapters[df.Columns[columnIndex]], columnValue)
		}
	}

	df.Values = chapters

	// Get columnTypes
	df.DataType = dataframe.GetColumnTypes(df)

	return df
}

// Higor interface
// TODO: Add interface to use higor as "hg" alias - ExportCSV

// Print DataFrame section
// TODO: Print DataFrame with Index
// TODO: Print a large DataFrame
// TODO: Print head DataFrame
// TODO: Print tail dataframe

// Read DataFrame
// TODO: ReadCSV with parsing values
// TODO: ReadCSV with multiples data types
// TODO: ReadCSv with nil datatypes
// TODO: ReadCSV with an specific nan value
// TODO: ReadCSV without header
// TODO: ReadCSV with more rows than columns

// Export CSV
// TODO: Export with nils values
// TODO: Export with multiple DataTypes
// TODO: Export without header
// TODO: Export without index
