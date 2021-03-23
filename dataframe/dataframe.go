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

// Letter Count how much data type.
type Letter map[string]int

// Words Count how much letter there are on a column
type Words map[string]Letter

// DataFrame Structure for DataFrame
type DataFrame struct {
	Columns  []string
	Values   book
	DataType Words
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

	// Get columnTypes
	df.DataType = getColumnTypes(df)

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

func getColumnTypes(df DataFrame) Words {
	/*
		s = String
		f = Float
		i = int
		b = bool
		n = nil
	*/
	//m := make(map[string]float64)

	myWords := make(Words)

	for key := range df.Values {
		myLetter := make(Letter)
		for _, v := range df.Values[key] {
			switch v.(type) {
			case int:
				myLetter["i"] = myLetter["i"] + 1
			case string:
				myLetter["s"] = myLetter["s"] + 1
			case float64:
				myLetter["f"] = myLetter["f"] + 1
			case bool:
				myLetter["b"] = myLetter["b"] + 1
			case nil:
				myLetter["n"] = myLetter["n"] + 1
			}
		}
		myWords[key] = myLetter

	}

	return myWords
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
	footer := []string{}
	for _, colName := range df.Columns {
		keys := reflect.ValueOf(df.DataType[colName]).MapKeys()
		keysString := []string{}
		for _, v := range keys {
			keysString = append(keysString, fmt.Sprintf("%v", v))
		}

		footer = append(footer, strings.Join(keysString, ","))
	}

	table := tablewriter.NewWriter(tableString)
	table.SetHeader(df.Columns)
	table.SetFooter(footer) // Todo: Change to another method
	table.AppendBulk(data[1:])
	table.SetBorder(false)
	table.SetCenterSeparator("|")

	table.Render()

	return tableString.String()
}
