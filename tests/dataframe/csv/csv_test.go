package csv

import (
	"reflect"
	"testing"

	"github.com/osmandi/higor/dataframe/csv"
)

// Check the header
func checkerHeader(expectedHeader, resultHeader []string, t *testing.T) {
	if !reflect.DeepEqual(expectedHeader, resultHeader) {
		t.Errorf("Header with errors. Expected %s, but received: %s", expectedHeader, resultHeader)
	}

}

// Columns Normal - Columns as expected
func TestCSVReadHeaderNormal(t *testing.T) {
	csvHeaderLineNormal := "col1,col2,col3,col4"
	expectedHeader := []string{"col1", "col2", "col3", "col4"}
	resultHeader := csv.CSVReadHeader(csv.Line(csvHeaderLineNormal))

	checkerHeader(expectedHeader, resultHeader, t)

}

// Columns Normal - With another separator
func TestCSVReadHeaderWithOtherSeparator(t *testing.T) {
	csvHeaderLineNormal := "col1|col2|col3|col4"
	expectedHeader := []string{"col1", "col2", "col3", "col4"}
	resultHeader := csv.CSVReadHeader(csv.Line(csvHeaderLineNormal), csv.Sep('|'))

	checkerHeader(expectedHeader, resultHeader, t)
}

// Columns - Missing columns
func TestCSVReadHeaderMissingValues(t *testing.T) {
	csvHeaderLine := "col1,col2,,col4"
	expectedHeader := []string{"col1", "col2", "", "col4"}
	resultHeader := csv.CSVReadHeader(csv.Line(csvHeaderLine))

	checkerHeader(expectedHeader, resultHeader, t)
}

// Columns - Columns names with numbers
func TestCSVReadHeaderNumbers(t *testing.T) {
	csvHeaderLine := "col1,2,col3,col4"
	expectedHeader := []string{"col1", "2", "col3", "col4"}
	resultHeader := csv.CSVReadHeader(csv.Line(csvHeaderLine))

	checkerHeader(expectedHeader, resultHeader, t)

}

// Columns - With lazy quotes
func TestCSVReadHeaderLazy(t *testing.T) {
	csvHeaderLine := "col1\",col2,col3,col4"
	expectedHeader := []string{"col1\"", "col2", "col3", "col4"}
	resultHeader := csv.CSVReadHeader(csv.Line(csvHeaderLine), csv.LazyQuotes(true))

	checkerHeader(expectedHeader, resultHeader, t)

}

// Columns - Empty columns
func TestCSVReadHeaderEmpty(t *testing.T) {
	csvHeaderLine := ""
	expectedHeader := []string{}
	resultHeader := csv.CSVReadHeader(csv.Line(csvHeaderLine))

	checkerHeader(expectedHeader, resultHeader, t)

}

// Columns - With \n
func TestCSVReadHeaderEnter(t *testing.T) {
	csvHeaderLine := "col1,col2,col3\n,col4"
	expectedHeader := []string{"col1", "col2", "col3"}
	resultHeader := csv.CSVReadHeader(csv.Line(csvHeaderLine))

	checkerHeader(expectedHeader, resultHeader, t)
}
