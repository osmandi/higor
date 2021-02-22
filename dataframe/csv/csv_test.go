package csv

import (
	"reflect"
	"testing"
)

////////////
// Header /
//////////

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
	resultHeader := CSVReadHeader(Line(csvHeaderLineNormal))

	checkerHeader(expectedHeader, resultHeader, t)

}

// Columns Normal - With another separator
func TestCSVReadHeaderWithOtherSeparator(t *testing.T) {
	csvHeaderLineNormal := "col1|col2|col3|col4"
	expectedHeader := []string{"col1", "col2", "col3", "col4"}
	resultHeader := CSVReadHeader(Line(csvHeaderLineNormal), Sep('|'))

	checkerHeader(expectedHeader, resultHeader, t)
}

// Columns - Missing columns
func TestCSVReadHeaderMissingValues(t *testing.T) {
	csvHeaderLine := "col1,col2,,col4"
	expectedHeader := []string{"col1", "col2", "", "col4"}
	resultHeader := CSVReadHeader(Line(csvHeaderLine))

	checkerHeader(expectedHeader, resultHeader, t)
}

// Columns - Columns names with numbers
func TestCSVReadHeaderNumbers(t *testing.T) {
	csvHeaderLine := "col1,2,col3,col4"
	expectedHeader := []string{"col1", "2", "col3", "col4"}
	resultHeader := CSVReadHeader(Line(csvHeaderLine))

	checkerHeader(expectedHeader, resultHeader, t)

}

// Columns - With lazy quotes
func TestCSVReadHeaderLazy(t *testing.T) {
	csvHeaderLine := "col1\",col2,col3,col4"
	expectedHeader := []string{"col1\"", "col2", "col3", "col4"}
	resultHeader := CSVReadHeader(Line(csvHeaderLine), LazyQuotes(true))

	checkerHeader(expectedHeader, resultHeader, t)

}

// Columns - Empty columns
func TestCSVReadHeaderEmpty(t *testing.T) {
	csvHeaderLine := ""
	expectedHeader := []string{}
	resultHeader := CSVReadHeader(Line(csvHeaderLine))

	checkerHeader(expectedHeader, resultHeader, t)

}

// Columns - With \n
func TestCSVReadHeaderEnter(t *testing.T) {
	csvHeaderLine := "col1,col2,col3\n,col4"
	expectedHeader := []string{"col1", "col2", "col3"}
	resultHeader := CSVReadHeader(Line(csvHeaderLine))

	checkerHeader(expectedHeader, resultHeader, t)
}

//////////
// Rows /
////////

// Rows - Normal row

// Rows - Missing values

// Rows - Lazy Quotes

// Rows - More values than expected

//////////////////////
// Matrix of values /
////////////////////

///////////////////////
// Read CSV filepath /
/////////////////////

////////////////
// Export CSV /
//////////////
