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

// Check the Row
func checkerRow(expectedRow, resultRow []string, t *testing.T) {
	if !reflect.DeepEqual(expectedRow, resultRow) {
		t.Errorf("Header with errors. Expected %s, but received: %s", expectedRow, resultRow)
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
// Split row to []string without parsing
// Row example: Name (string), age (int), live(bool), salary(float)

// Rows - Normal row
func TestCSVReadRowNormal(t *testing.T) {
	csvRowLine := "Harry Potter,21,true,1000.3"
	expectedRow := []string{"Harry Potter", "21", "true", "1000.3"}
	resultRow := CSVReadRowNormal(Line(csvRowLine))

	checkerRow(expectedRow, resultRow, t)
}

// Rows - Missing values
func TestCSVReadRowMissingValues(t *testing.T) {
	csvRowLine := "Harry Potter,21,,1000.3"
	expectedRow := []string{"Harry Potter", "21", "", "1000.3"}
	resultRow := CSVReadRowNormal(Line(csvRowLine))

	checkerRow(expectedRow, resultRow, t)
}

// Rows - Another separator
func TestCSVReadRowAnotherSeparator(t *testing.T) {
	csvRowLine := "Harry Potter|21|true|1000.3"
	expectedRow := []string{"Harry Potter", "21", "true", "1000.3"}
	resultRow := CSVReadRowNormal(Line(csvRowLine), Sep('|'))

	checkerRow(expectedRow, resultRow, t)
}

// Rows - Lazy Quotes
func TestCSVReadRowLazyQuotes(t *testing.T) {
	//csvHeaderLine := "col1\",col2,col3,col4"
	csvRowLine := "Harry Potter\",21,true,1000.3"
	expectedRow := []string{"Harry Potter\"", "21", "true", "1000.3"}
	resultRow := CSVReadRowNormal(Line(csvRowLine), LazyQuotes(true))

	checkerRow(expectedRow, resultRow, t)
}

// Rows - With \n (Not removing)
func TestCSVReadRowNotNewLine(t *testing.T) {
	csvRowLine := "Harry Potter,21,true\n,1000.3"
	expectedRow := []string{"Harry Potter", "21", "true"}
	resultRow := CSVReadRowNormal(Line(csvRowLine), LazyQuotes(true))

	checkerRow(expectedRow, resultRow, t)

}

// Rows - With \n (Yes removing)
func TestCSVReadRowRemoveNewLine(t *testing.T) {
	csvRowLine := "Harry Potter,21,true\n,1000.3"
	expectedRow := []string{"Harry Potter", "21", "true", "1000.3"}
	resultRow := CSVReadRowNormal(Line(csvRowLine), LazyQuotes(true), RemoveNewLine(true))

	checkerRow(expectedRow, resultRow, t)

}

// Rows - Empty row
func TestCSVReadRowEmpty(t *testing.T) {
	csvRowLine := ""
	expectedRow := []string{}
	resultRow := CSVReadRowNormal(Line(csvRowLine), LazyQuotes(true), RemoveNewLine(true))

	checkerRow(expectedRow, resultRow, t)

}

///////////////////////////////////////////////////
// Matrix of values exported lines to [][]string /
/////////////////////////////////////////////////

// Matrix - Header and Rows sush as expected
// Matrix - Header empty
// Matrix - Rows empty
// Matrix - Header most columns than expected
// Matrix - Header less columns than expected
// Matrix - Row most columns than expected
// Matrix - Row less columns than expected

///////////////////////
// Read CSV filepath /
/////////////////////

////////////////
// Export CSV /
//////////////
