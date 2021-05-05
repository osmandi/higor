package dataframe

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/olekukonko/tablewriter"
)

// PageString Data type for string values
type PageString []interface{}

// PageBool Data type for boolean values
type PageBool []interface{}

// PageFloat64 Data type for numbers and float values
type PageFloat64 []float64

// Page values saved on a map
type Page interface{}

// Book Interface to save a DataFrame
type Book map[string]Page

// DataFrame Structure for DataFrame
type DataFrame struct {
	Columns []string
	Values  Book
}

// IsEqual to kown if two DataFrame are equal
func IsEqual(dataFrame1, dataFrame2 DataFrame) bool {
	return reflect.DeepEqual(dataFrame1, dataFrame2)

}

/////////
// CSV /

// CSVCreatorMock sample csv to tests
func CSVCreatorMock(data [][]string, separator rune) *os.File {
	// Temp file
	tmpFile, err := ioutil.TempFile(os.TempDir(), "higorCSVTest-*.csv")
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(tmpFile)
	writer.Comma = separator
	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			log.Fatal(err)
		}
	}

	writer.Flush()

	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}
	return tmpFile

}

// CSVChecker To Check csv result
func CSVChecker(dataExpected, dataResult [][]string, t *testing.T) {
	if !reflect.DeepEqual(dataExpected, dataResult) {
		t.Errorf("Header with errors. \nExpected: \n%s. \nReceived: \n%s", dataExpected, dataResult)
	}
}

// DataFrameChecker To check if two DataFrame are equal
func DataFrameChecker(dfExpected, dfResult DataFrame, t *testing.T) {
	isEqual := IsEqual(dfExpected, dfResult)
	if !isEqual {
		t.Errorf("dfExpected and dfResult are distinct.\ndfExpected: \n%v \ndfResult: \n%v", dfExpected, dfResult)
	}

}

//////

// CSV type
type CSV struct {
	Sep        rune
	LineString string
	LazyQuotes bool
}

// ErrorChecker to kown if there are error
func ErrorChecker(err error) {
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

// GetValues get all values
func (df DataFrame) GetValues() [][]string {
	return trasposeRows(df)[1:]
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
		valuesIterate := []interface{}{}
		values := reflect.ValueOf(colValues)
		for i := 0; i < values.Len(); i++ {
			valuesIterate = append(valuesIterate, values.Index(i))
		}
		if colOk {
			for rowIndex, value := range valuesIterate {
				var v interface{}
				v = value
				if value == nil {
					v = ""
				}
				data[rowIndex+1] = append(data[rowIndex+1], fmt.Sprintf("%v", v))
			}
		}
	}

	return data
}

// exportCSV To export a dataframe to CSV file
func exportCSV(filename string, data [][]string, opts ...CSVOption) {
	csvInternal := &CSV{}
	csvInternal.Sep = ','

	for _, opt := range opts {
		opt(csvInternal)
	}

	csvFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	ErrorChecker(err)
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Comma = csvInternal.Sep
	err = csvWriter.WriteAll(data)
	ErrorChecker(err)

}

func (df DataFrame) String() string {
	tableString := &strings.Builder{}
	data := trasposeRows(df)
	table := tablewriter.NewWriter(tableString)
	table.SetHeader(df.Columns)
	table.AppendBulk(data[1:])
	table.SetBorder(true)
	table.SetCenterSeparator("+")

	table.Render()

	return tableString.String()
}

// ToCSV Export DataFrame to CSV
func (df DataFrame) ToCSV(filename string) {
	data := [][]string{}
	data = append(data, df.Columns)

	dfInternal := DataFrame{
		Columns: df.Columns,
		Values:  df.Values,
	}

	data = append(data, dfInternal.GetValues()...)

	exportCSV(filename, data)
}
