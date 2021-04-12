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

// Book Interface to save a DataFrame
type Book map[string][]interface{}

// Letter Count how much data type.
type Letter map[string]int

// Words Count how much letter there are on a column
type Words map[string]Letter

// DataFrame Structure for DataFrame
type DataFrame struct {
	Columns  []string
	Values   Book
	DataType Words
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

// GetColumnTypes To know data type
func GetColumnTypes(df DataFrame) Words {
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
