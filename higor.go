package higor

import (
	"encoding/csv"
	"fmt"
	"log"
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
	layout := "2006-01-02" // YYYY-MM-DD

	if csvInternal.Dateformat != "" {
		value := strings.Replace(csvInternal.Dateformat, "YYYY", "2006", 1)
		value = strings.Replace(value, "MM", "01", 1)
		value = strings.Replace(value, "DD", "02", 1)
		layout = value
	}

	// If schema is set
	if len(csvInternal.Schema) > 0 {
		/*
			for i := range csvInternal.Schema {
				df.Values = append(df.Values, csvInternal.Schema[i])
			}
		*/
		df.Values = csvInternal.Schema

		//		df.Values = csvInternal.Schema

		//		chapters := []interface{}{}
		//		valueMap := make(map[string][]interface{})

		for _, rowValue := range csv[1:] {
			for columnIndex, columnValue := range rowValue {
				//				chapters[columnIndex] = append(chapters[columnIndex], columnValue)
				//				valueMap[df.Columns[columnIndex]] = append(valueMap[df.Columns[columnIndex]], columnValue)

				switch df.Values[columnIndex].(type) {
				case dataframe.PageString:
					df.Values[columnIndex] = append(df.Values[columnIndex].(dataframe.PageString), columnValue)
				case dataframe.PageFloat64:
					valueFloat64, err := strconv.ParseFloat(columnValue, 64)
					if err != nil {
						log.Fatal(err)
					}
					df.Values[columnIndex] = append(df.Values[columnIndex].(dataframe.PageFloat64), valueFloat64)

				case dataframe.PageBool:
					valueBool, err := strconv.ParseBool(columnValue)
					if err != nil {
						log.Fatal(err)
					}
					df.Values[columnIndex] = append(df.Values[columnIndex].(dataframe.PageBool), valueBool)
				case dataframe.PageDatetime:
					dateValue, err := time.Parse(layout, columnValue)
					if err != nil {
						log.Fatal(err)
					}
					df.Values[columnIndex] = append(df.Values[columnIndex].(dataframe.PageDatetime), dateValue)
				case dataframe.PageAny:
					df.Values[columnIndex] = append(df.Values[columnIndex].(dataframe.PageAny), columnValue)

				}

				//fmt.Println(df.Values[columnIndex], columnValue)
				//fmt.Println(reflect.TypeOf(df.Values[columnIndex]))
			}

		}
	}
	//df.Values = chapters

	return df
}
