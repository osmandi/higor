package csv

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
)

/////////////////////////////////////////////////
// CSV Options - To create optional parameters /
///////////////////////////////////////////////
func TestErrorChecker(t *testing.T) {
	errorChecker(nil)
}

func TestSep(t *testing.T) {
	separator := ';'
	csvResult := &CSV{}
	csvOptionInternal := Sep(separator)
	csvOptionInternal(csvResult)

	if csvResult.Sep != ';' {
		t.Errorf("Sep error. Expected: ';'. But result: %v", csvResult.Sep)
	}

}

func TestLine(t *testing.T) {
	line := "textLine"
	csvResult := &CSV{}
	csvOptionInternal := Line(line)
	csvOptionInternal(csvResult)

	if csvResult.LineString != line {
		t.Errorf("Line error. Expected %s. But received: %v", line, csvResult.LineString)
	}
}

func TestLazyQuotes(t *testing.T) {
	lazyQuotes := true
	csvResult := &CSV{}
	csvOptionInternal := LazyQuotes(true)
	csvOptionInternal(csvResult)

	if csvResult.LazyQuotes != lazyQuotes {
		t.Errorf("Lazy Quotes error. Expected: %v. But received: %v", lazyQuotes, csvResult.LazyQuotes)
	}
}

func TestRemoveNewLine(t *testing.T) {
	removeNewLine := true
	csvResult := &CSV{}
	csvOptionInternal := RemoveNewLine(removeNewLine)
	csvOptionInternal(csvResult)

	if csvResult.RemoveNewLine != removeNewLine {
		t.Errorf("Remove new line error. Expected: %v. But received: %v", removeNewLine, csvResult.RemoveNewLine)
	}

}

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

///////////////////////////////////////////
// Read CSV filepath - Return [][]string /
/////////////////////////////////////////
func csvCheker(dataExpected, dataResult [][]string, t *testing.T) {
	if !reflect.DeepEqual(dataExpected, dataResult) {
		t.Errorf("Header with errors. Expected %s, but received: %s", dataExpected, dataResult)
	}
}

func csvCreatorMock(data [][]string, separator rune) *os.File {
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

	//defer os.Remove((tmpFile.Name()))

	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}
	return tmpFile

}

// Normal CSV
func TestReadCSV(t *testing.T) {
	// Mock data
	dataExpected := [][]string{{"col1", "col2", "col3"}, {"row11", "row12", "row13"}, {"row21", "row22", "row23"}}
	tmpCSV := csvCreatorMock(dataExpected, ',')
	defer os.Remove((tmpCSV.Name()))

	// Test
	dataResult := ReadCSV(tmpCSV.Name())
	csvCheker(dataExpected, dataResult, t)

}

// Normal CSV with another separator ('|')
func TestReadCSVAnotherSeparator(t *testing.T) {
	// Mock data
	dataExpected := [][]string{{"col1", "col2", "col3"}, {"row11", "row12", "row13"}, {"row21", "row22", "row23"}}
	tmpCSV := csvCreatorMock(dataExpected, '|')
	tmpCSVFilename := tmpCSV.Name()
	defer os.Remove(tmpCSVFilename)

	// Test
	dataResult := ReadCSV(tmpCSVFilename, Sep('|'))
	csvCheker(dataExpected, dataResult, t)

}

// CSV with new line on rows
func TestReadCSVNewLine(t *testing.T) {
	// Mock data
	dataExpected := [][]string{{"col1", "col2\n", "col3"}, {"row11\n", "row12", "row13"}, {"row21", "row22", "row23"}}
	tmpCSV := csvCreatorMock(dataExpected, ',')
	tmpCSVFilename := tmpCSV.Name()
	defer os.Remove(tmpCSVFilename)

	// Test
	dataResult := ReadCSV(tmpCSVFilename)
	csvCheker(dataExpected, dataResult, t)
}

// CSV with lazy quotes on row
func TestReadCSVLazyQuotes(t *testing.T) {
	// Mock data
	dataExpected := [][]string{{"col1", "col2\"", "col3"}, {"row11\"", "row12", "row13"}, {"row21", "row22", "row23"}}
	tmpCSV := csvCreatorMock(dataExpected, ',')
	tmpCSVFilename := tmpCSV.Name()
	defer os.Remove(tmpCSVFilename)

	// Test
	dataResult := ReadCSV(tmpCSVFilename)
	csvCheker(dataExpected, dataResult, t)
}

////////////////
// Export CSV /
//////////////
// Export - Normal CSV
func TestExportCSVFileExists(t *testing.T) {
	// Temp file

	tmpFile, err := ioutil.TempFile(os.TempDir(), "higorCSVTestExport-*.csv")
	if err != nil {
		log.Fatal(err)
	}
	filename := tmpFile.Name()

	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}

	// Export to CSV
	dataExpected := [][]string{{"col1", "col2", "col3"}, {"row11", "row12", "row13"}, {"row21", "row22", "row23"}}
	ExportCSV(filename, dataExpected)

	// Read the CSV content
	csvOpen, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer csvOpen.Close()
	csvReader := csv.NewReader(csvOpen)
	dataResult, err := csvReader.ReadAll()
	csvCheker(dataExpected, dataResult, t)
	defer os.Remove(filename)
}

// Export - File doesn't exists
func TestExportCSVDoesNotExists(t *testing.T) {
	filename := "higorCSVTestExport-DoesNotExists.csv"

	// Export to CSV
	dataExpected := [][]string{{"col1", "col2", "col3"}, {"row11", "row12", "row13"}, {"row21", "row22", "row23"}}
	ExportCSV(filename, dataExpected)

	// Read the CSV content
	csvOpen, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer csvOpen.Close()
	csvReader := csv.NewReader(csvOpen)
	dataResult, err := csvReader.ReadAll()
	csvCheker(dataExpected, dataResult, t)

	// Delete file created
	defer os.Remove(filename)

}

// Export - With another separator
func TestExportCSVAnotherSeparator(t *testing.T) {
	// Separator
	sep := '|'

	// Temp file
	tmpFile, err := ioutil.TempFile(os.TempDir(), "higorCSVTestExport-*.csv")
	if err != nil {
		log.Fatal(err)
	}
	filename := tmpFile.Name()

	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}

	// Export to CSV
	dataExpected := [][]string{{"col1", "col2", "col3"}, {"row11", "row12", "row13"}, {"row21", "row22", "row23"}}
	ExportCSV(filename, dataExpected, Sep(sep))

	// Read the CSV content
	csvOpen, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer csvOpen.Close()
	csvReader := csv.NewReader(csvOpen)
	csvReader.Comma = sep
	dataResult, err := csvReader.ReadAll()
	csvCheker(dataExpected, dataResult, t)
	defer os.Remove(filename)
}

// Export - Without index
// Export - With index
// Export - Without Header
