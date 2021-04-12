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
	ErrorChecker(nil)
}

//////////////////////////
// Utilities DataFrame //
////////////////////////

func TestTrasposeRowsMultipleDataType(t *testing.T) {
	dataExpected := [][]string{{"col1", "col2", "col3"}, {"1", "", "row13"}, {"row21", "row22", "row23"}}

	dfExpected := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: Book{
			"col1": {1, "row21"},
			"col2": {nil, "row22"},
			"col3": {"row13", "row23"},
		},
	}

	dataResult := trasposeRows(dfExpected)

	CSVChecker(dataExpected, dataResult, t)

}

func TestTrasposeRowsString(t *testing.T) {
	dataExpected := [][]string{{"col1", "col2", "col3"}, {"row11", "row12", "row13"}, {"row21", "row22", "row23"}}

	dfExpected := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: Book{
			"col1": {"row11", "row21"},
			"col2": {"row12", "row22"},
			"col3": {"row13", "row23"},
		},
	}

	dataResult := trasposeRows(dfExpected)

	CSVChecker(dataExpected, dataResult, t)

}

func TestValuesNormal(t *testing.T) {
	dataExpected := [][]string{{"row11", "row12", "row13"}, {"row21", "row22", "row23"}}

	dfExpected := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: Book{
			"col1": {"row11", "row21"},
			"col2": {"row12", "row22"},
			"col3": {"row13", "row23"},
		},
	}

	dataResult := dfExpected.GetValues()

	CSVChecker(dataExpected, dataResult, t)

}

func TestGetColumnTypesOnlyString(t *testing.T) {
	df := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: Book{
			"col1": {"row11", "row21"},
			"col2": {"row12", "row22"},
			"col3": {"row13", "row23"},
		},
	}

	typeColumnsExpected := Words{
		"col1": Letter{"s": 2},
		"col2": Letter{"s": 2},
		"col3": Letter{"s": 2},
	}

	typeColmnsResult := GetColumnTypes(df)

	if !reflect.DeepEqual(typeColumnsExpected, typeColmnsResult) {
		t.Errorf("Header with errors. \nExpected: \n%v. \nReceived: \n%v", typeColumnsExpected, typeColmnsResult)
	}

}

func TestGetColumnTypesMultipleDataType(t *testing.T) {
	df := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: Book{
			"col1": {"row11", 1.2},
			"col2": {1, "row22"},
			"col3": {"row13", "row23"},
		},
	}

	typeColumnsExpected := Words{
		"col1": Letter{"s": 1, "f": 1},
		"col2": Letter{"s": 1, "i": 1},
		"col3": Letter{"s": 2},
	}

	typeColmnsResult := GetColumnTypes(df)

	if !reflect.DeepEqual(typeColumnsExpected, typeColmnsResult) {
		t.Errorf("Header with errors. \nExpected: \n%v. \nReceived: \n%v", typeColumnsExpected, typeColmnsResult)
	}

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

func TestEqualDataFrame(t *testing.T) {

	columns := []string{"colInt", "colString", "colBool", "colFloat"}
	chapters := Book{
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
	chapters := Book{
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

/////////////////////
// Print DataFrame /
///////////////////

func TestPrintDataFrame(t *testing.T) {
	columns := []string{"col1", "col2", "col3"}
	chapters := Book{
		"col1": {"row11", "row21"},
		"col2": {"row12", "row22"},
		"col3": {"row13", "row23"},
	}

	typeColumnsExpected := Words{
		"col1": Letter{"s": 2},
		"col2": Letter{"s": 2},
		"col3": Letter{"s": 2},
	}

	df := DataFrame{
		Columns:  columns,
		Values:   chapters,
		DataType: typeColumnsExpected,
	}

	tableExpectedFormat := "  COL1  | COL2  | COL3   \n--------|-------|--------\n  row11 | row12 | row13  \n  row21 | row22 | row23  \n--------|-------|--------\n    S   |   S   |   S    \n--------|-------|--------\n"

	tableResultFormat := df.String()

	if tableExpectedFormat != tableResultFormat {
		t.Errorf("Table format error.\nExpected:\n%v\nResult:\n%v", tableExpectedFormat, tableResultFormat)
	}

}

func TestPrintDataFrameMultipleDataType(t *testing.T) {
	columns := []string{"col1", "col2", "col3"}
	chapters := Book{
		"col1": {1, "row21"},
		"col2": {2.3, "row22"},
		"col3": {"row13", "row23"},
	}

	typeColumnsExpected := Words{
		"col1": Letter{"s": 1, "i": 1},
		"col2": Letter{"s": 1, "f": 1},
		"col3": Letter{"s": 2},
	}

	df := DataFrame{
		Columns:  columns,
		Values:   chapters,
		DataType: typeColumnsExpected,
	}

	tableExpectedFormat := "  COL1  | COL2  | COL3   \n--------|-------|--------\n      1 |   2.3 | row13  \n  row21 | row22 | row23  \n--------|-------|--------\n   S,I  |  S,F  |   S    \n--------|-------|--------\n"

	tableResultFormat := df.String()

	if tableExpectedFormat != tableResultFormat {
		t.Errorf("Table format error.\nExpected:\n%v\nResult:\n%v", tableExpectedFormat, tableResultFormat)
	}

}

func TestPrintDataFrameWithNils(t *testing.T) {
	columns := []string{"col1", "col2", "col3"}
	chapters := Book{
		"col1": {nil, "row21"},
		"col2": {"row12", "row22"},
		"col3": {"row13", "row23"},
	}

	typeColumnsExpected := Words{
		"col1": Letter{"s": 1, "n": 1},
		"col2": Letter{"s": 2},
		"col3": Letter{"s": 2},
	}

	df := DataFrame{
		Columns:  columns,
		Values:   chapters,
		DataType: typeColumnsExpected,
	}

	tableExpectedFormat := "  COL1  | COL2  | COL3   \n--------|-------|--------\n        | row12 | row13  \n  row21 | row22 | row23  \n--------|-------|--------\n   S,N  |   S   |   S    \n--------|-------|--------\n"

	tableResultFormat := df.String()

	if tableExpectedFormat != tableResultFormat {
		t.Errorf("Table format error.\nExpected:\n%v\nResult:\n%v", tableExpectedFormat, tableResultFormat)
	}

}

///////////////
// ExportCSV /
/////////////

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
	CSVChecker(dataExpected, dataResult, t)
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
	CSVChecker(dataExpected, dataResult, t)

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
	CSVChecker(dataExpected, dataResult, t)
	defer os.Remove(filename)
}

/////////////////////////////////////
// DataFrame Description functions /
///////////////////////////////////
