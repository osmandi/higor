package dataframe

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
)

func TestErrorChecker(t *testing.T) {
	errorChecker(nil)
}

//////////////////////////
// Dataframe functions //
////////////////////////

// Utilities DataFrame
func csvCheker(dataExpected, dataResult [][]string, t *testing.T) {
	if !reflect.DeepEqual(dataExpected, dataResult) {
		t.Errorf("Header with errors. \nExpected: \n%s. \nReceived: \n%s", dataExpected, dataResult)
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

	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}
	return tmpFile

}

func dataFrameChecker(dfExpected, dfResult DataFrame, t *testing.T) {
	isEqual := IsEqual(dfExpected, dfResult)
	if !isEqual {
		t.Errorf("dfExpected and dfResult are distinct.\ndfExpected: %v \ndfResult: %v", dfExpected, dfResult)
	}

}

func TestPrintDataFrame(t *testing.T) {
	columns := []string{"col1", "col2", "col3"}
	chapters := book{
		"col1": {"row11", "row21"},
		"col2": {"row12", "row22"},
		"col3": {"row13", "row23"},
	}

	df := DataFrame{
		Columns: columns,
		Values:  chapters,
	}

	tableExpectedFormat := "   COL1  |  COL2  |  COL3   \n---------|--------|---------\n  row11  | row12  | row13   \n  row21  | row22  | row23   \n---------|--------|---------\n  STRING | STRING | STRING  \n---------|--------|---------\n"

	tableResultFormat := df.String()

	if tableExpectedFormat != tableResultFormat {
		t.Errorf("Table format error.\nExpected:\n%v\nResult:\n%v", tableExpectedFormat, tableResultFormat)
	}

}

func TestTrasposeRows(t *testing.T) {
	dataExpected := [][]string{{"col1", "col2", "col3"}, {"row11", "row12", "row13"}, {"row21", "row22", "row23"}}

	dfExpected := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: book{
			"col1": {"row11", "row21"},
			"col2": {"row12", "row22"},
			"col3": {"row13", "row23"},
		},
	}

	dataResult := trasposeRows(dfExpected)

	csvCheker(dataExpected, dataResult, t)

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

// ReadCSV

func TestLazyQuotes(t *testing.T) {
	lazyQuotesBool := true
	csvResult := &CSV{}
	csvOptionInternal := LazyQuotes(lazyQuotesBool)
	csvOptionInternal(csvResult)

	if csvResult.LazyQuotes != lazyQuotesBool {
		t.Errorf("Lazy Quotes error. Expected: %v. But received: %v", lazyQuotesBool, csvResult.LazyQuotes)
	}
}

func TestReadCSVNormal(t *testing.T) {
	// Mockup
	data := [][]string{{"col1", "col2", "col3"}, {"row11", "row12", "row13"}, {"row21", "row22", "row23"}}
	separator := ','
	csvTempFile := csvCreatorMock(data, separator)
	csvTempFilename := csvTempFile.Name()
	defer os.Remove(csvTempFilename)

	dfExpected := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: book{
			"col1": {"row11", "row21"},
			"col2": {"row12", "row22"},
			"col3": {"row13", "row23"},
		},
	}

	dfResult := ReadCSV(csvTempFilename)

	dataFrameChecker(dfExpected, dfResult, t)

}

func TestReadCSVAnoterSeparator(t *testing.T) {
	// Mockup
	data := [][]string{{"col1", "col2", "col3"}, {"row11", "row12", "row13"}, {"row21", "row22", "row23"}}
	separator := '|'
	csvTempFile := csvCreatorMock(data, separator)
	csvTempFilename := csvTempFile.Name()
	defer os.Remove(csvTempFilename)

	dfExpected := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: book{
			"col1": {"row11", "row21"},
			"col2": {"row12", "row22"},
			"col3": {"row13", "row23"},
		},
	}

	dfResult := ReadCSV(csvTempFilename, Sep('|'))

	dataFrameChecker(dfExpected, dfResult, t)

}

func TestReadCSVWithLazyQuotes(t *testing.T) {
	// Mockup
	data := [][]string{{"col1\"", "col2", "col3"}, {"row11\"", "row12", "row13"}, {"row21", "row22", "row23"}}
	separator := '|'
	csvTempFile := csvCreatorMock(data, separator)
	csvTempFilename := csvTempFile.Name()
	defer os.Remove(csvTempFilename)

	dfExpected := DataFrame{
		Columns: []string{"col1\"", "col2", "col3"},
		Values: book{
			"col1\"": {"row11\"", "row21"},
			"col2":   {"row12", "row22"},
			"col3":   {"row13", "row23"},
		},
	}

	dfResult := ReadCSV(csvTempFilename, Sep('|'), LazyQuotes(true))

	dataFrameChecker(dfExpected, dfResult, t)

}

// ExportCSV
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
	ExportCSV(filename, dataExpected, Sep(separator))

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

// Equal DataFrame
func TestEqualDataFrame(t *testing.T) {

	columns := []string{"colInt", "colString", "colBool", "colFloat"}
	chapters := book{
		"colInt":    {1, nil, 2, 3},
		"colString": {"hola", "que", "hace", nil},
		"colBool":   {nil, true, false, nil},
		"colFloat":  {3.2, 5.4, nil, nil},
	}

	df1 := DataFrame{
		Columns: columns,
		Values:  chapters,
	}
	df2 := DataFrame{
		Columns: columns,
		Values:  chapters,
	}

	isEqual := IsEqual(df1, df2)

	if !isEqual {
		t.Errorf("Error equalDataframe. df1 and df2 are different! But equal expected!")
	}
}

// Diferent DataFrame
func TestDifferentlDataFrame(t *testing.T) {
	columns := []string{"colInt", "colString", "colBool", "colFloat"}
	chapters := book{
		"colInt2":    {1, nil, 2, 3},
		"colString2": {"hola", "que", "hace", nil},
		"colBool2":   {nil, true, false, nil},
		"colFloat2":  {3.2, 5.4, nil, nil},
	}

	df1 := DataFrame{
		Columns: columns,
		Values:  chapters,
	}
	df2 := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values:  chapters,
	}

	isEqual := IsEqual(df1, df2)

	if isEqual {
		t.Errorf("Error differentDataframe. df1 and df2 are eual! But different expected!")
	}
}

// Print DataFrame with Index (TODO)
// Print a DataFrame

// Print a large DataFrame

// Print header DataFrame

// Print tail dataframe

// Print DataFrame with unknown format

// More test cases to ReadCSV
// More test cases to ExportCSV
