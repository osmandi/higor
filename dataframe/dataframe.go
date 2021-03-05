package dataframe

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type book map[string][]interface{}

// DataFrame Structure for DataFrame
type DataFrame struct {
	Columns []string
	Values  book
}

// IsEqual to kown if two DataFrame are equal
func IsEqual(dataFrame1, dataFrame2 DataFrame) bool {
	return reflect.DeepEqual(dataFrame1, dataFrame2)

}

/////////
// CSV /
///////

// CSV type
type CSV struct {
	Sep        rune
	LineString string
	LazyQuotes bool
}

func errorChecker(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// CSVOption alternative parameters
type CSVOption func(c *CSV)

// Sep CSV separator in rune type: ',', ';', '|', etc...
func Sep(separator rune) CSVOption {
	return func(c *CSV) {
		c.Sep = separator
	}
}

// LazyQuotes Interpreting internal quotes
func LazyQuotes(lazy bool) CSVOption {
	return func(c *CSV) {
		c.LazyQuotes = lazy
	}
}

// ReadCSV Read a CSV file and save it as a DataFrame
func ReadCSV(filename string, opts ...CSVOption) DataFrame {
	csvInternal := &CSV{}
	csvInternal.Sep = ','

	for _, opt := range opts {
		opt(csvInternal)
	}

	// Open file
	csvFile, err := os.Open(filename)
	errorChecker(err)
	defer csvFile.Close()

	// Read CSV
	csvReader := csv.NewReader(csvFile)
	csvReader.Comma = csvInternal.Sep
	csvReader.LazyQuotes = csvInternal.LazyQuotes

	// Convert CSV to [][]string
	csv, err := csvReader.ReadAll()
	errorChecker(err)

	df := DataFrame{}
	df.Columns = csv[0]
	chapters := book{}

	for _, rowValue := range csv[1:] {
		for columnIndex, columnValue := range rowValue {
			chapters[df.Columns[columnIndex]] = append(chapters[df.Columns[columnIndex]], columnValue)
		}
	}

	df.Values = chapters

	return df
}

func trasposeRows(df DataFrame) [][]string {
	data := [][]string{}

	// Add []string empties
	for range df.Columns {
		data = append(data, []string{})
	}

	// Add columns names
	data[0] = df.Columns

	// Traspose row
	for _, colName := range df.Columns {
		colValues, colOk := df.Values[colName]
		if colOk {
			for rowIndex, value := range colValues {
				data[rowIndex+1] = append(data[rowIndex+1], fmt.Sprintf("%v", value))
			}
		}
	}

	return data
}

// ExportCSV To export a dataframe to CSV file
func ExportCSV(filename string, data [][]string, opts ...CSVOption) {
	csvInternal := &CSV{}
	csvInternal.Sep = ','

	for _, opt := range opts {
		opt(csvInternal)
	}

	csvFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	errorChecker(err)
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Comma = csvInternal.Sep
	err = csvWriter.WriteAll(data)
	errorChecker(err)

}

func (df DataFrame) String() string {
	tableString := &strings.Builder{}
	data := trasposeRows(df)

	table := tablewriter.NewWriter(tableString)
	table.SetHeader(df.Columns)
	table.SetFooter([]string{"", "", ""}) // Todo: Change to another method
	table.AppendBulk(data)
	table.SetBorder(false)
	table.SetCenterSeparator("|")

	table.Render()

	return tableString.String()
}
