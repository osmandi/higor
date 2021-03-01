package dataframe

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
	csvOptionInternal := sep(separator)
	csvOptionInternal(csvResult)

	if csvResult.sep != ';' {
		t.Errorf("Sep error. Expected: ';'. But result: %v", csvResult.sep)
	}

}

func TestLazyQuotes(t *testing.T) {
	lazyQuotesBool := true
	csvResult := &CSV{}
	csvOptionInternal := lazyQuotes(lazyQuotesBool)
	csvOptionInternal(csvResult)

	if csvResult.lazyQuotes != lazyQuotesBool {
		t.Errorf("Lazy Quotes error. Expected: %v. But received: %v", lazyQuotesBool, csvResult.lazyQuotes)
	}
}

func TestRemoveNewLine(t *testing.T) {
	removeNewLineBool := true
	csvResult := &CSV{}
	csvOptionInternal := removeNewLine(removeNewLineBool)
	csvOptionInternal(csvResult)

	if csvResult.removeNewLine != removeNewLineBool {
		t.Errorf("Remove new line error. Expected: %v. But received: %v", removeNewLineBool, csvResult.removeNewLine)
	}

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
	dataResult := readCSV(tmpCSV.Name())
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
	dataResult := readCSV(tmpCSVFilename, sep('|'))
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
	dataResult := readCSV(tmpCSVFilename)
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
	dataResult := readCSV(tmpCSVFilename)
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
	exportCSV(filename, dataExpected)

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
	exportCSV(filename, dataExpected)

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
	separator := '|'

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
	exportCSV(filename, dataExpected, sep(separator))

	// Read the CSV content
	csvOpen, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer csvOpen.Close()
	csvReader := csv.NewReader(csvOpen)
	csvReader.Comma = separator
	dataResult, err := csvReader.ReadAll()
	csvCheker(dataExpected, dataResult, t)
	defer os.Remove(filename)
}

// Export - Without index
// Export - With index
// Export - Without Header
