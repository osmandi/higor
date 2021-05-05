package higor

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/osmandi/higor/dataframe"
)

var Version string = "v0.3.0"

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
	df.Values = csvInternal.Schema
	//	chapters := csvInternal.Schema
	valueMap := make(map[string][]interface{})

	for _, rowValue := range csv[1:] {
		for columnIndex, columnValue := range rowValue {
			//chapters[df.Columns[columnIndex]] = append(chapters[df.Columns[columnIndex]], columnValue)
			valueMap[df.Columns[columnIndex]] = append(valueMap[df.Columns[columnIndex]], columnValue)
			//df.Values[df.Columns[columnIndex]] = append(df.Values[df.Columns[columnIndex]], columnValue)
		}

	}

	for key, value := range valueMap {
		df.Values[key] = value
	}

	//df.Values = chapters

	return df
}
