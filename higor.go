package higor

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

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
	csvInternal.None = ""
	dateNaN := time.Date(0001, 1, 1, 0, 0, 0, 0, time.UTC)

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
	df.Values = dataframe.Book{}
	layout := "2006-01-02" // Dafault: YYYY-MM-DD

	if csvInternal.Dateformat != "" {
		value := strings.Replace(csvInternal.Dateformat, "YYYY", "2006", 1)
		value = strings.Replace(value, "MM", "01", 1)
		value = strings.Replace(value, "DD", "02", 1)
		layout = value
	}

	// If schema is set
	if len(csvInternal.Schema) > 0 {
		df.Values = csvInternal.Schema

		for _, rowValue := range csv[1:] {
			for columnIndex, columnValue := range rowValue {

				switch df.Values[columnIndex].(type) {
				case dataframe.PageString:
					df.Values[columnIndex] = append(df.Values[columnIndex].(dataframe.PageString), columnValue)
				case dataframe.PageFloat64:
					if columnValue != csvInternal.None {
						valueFloat64, err := strconv.ParseFloat(columnValue, 64)
						dataframe.ErrorChecker(err)
						df.Values[columnIndex] = append(df.Values[columnIndex].(dataframe.PageFloat64), valueFloat64)
					} else {
						df.Values[columnIndex] = append(df.Values[columnIndex].(dataframe.PageFloat64), math.NaN())
					}
				case dataframe.PageBool:
					valueBool, err := strconv.ParseBool(columnValue)
					dataframe.ErrorChecker(err)
					df.Values[columnIndex] = append(df.Values[columnIndex].(dataframe.PageBool), valueBool)
				case dataframe.PageDatetime:
					if columnValue != csvInternal.None {
						dateValue, err := time.Parse(layout, columnValue)
						dataframe.ErrorChecker(err)
						df.Values[columnIndex] = append(df.Values[columnIndex].(dataframe.PageDatetime), dateValue)
					} else {
						df.Values[columnIndex] = append(df.Values[columnIndex].(dataframe.PageDatetime), dateNaN)
					}
				case dataframe.PageAny:
					if columnValue != csvInternal.None {
						df.Values[columnIndex] = append(df.Values[columnIndex].(dataframe.PageAny), columnValue)
					} else {
						df.Values[columnIndex] = append(df.Values[columnIndex].(dataframe.PageAny), math.NaN())
					}

				}

			}

		}
	}

	return df
}
