package dataframe

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"math"
	"os"
	"testing"
)

func TestErrorChecker(t *testing.T) {
	ErrorChecker(nil)
}

//////////////////
// Export CSVs //
////////////////
func TestToCSVNormal(t *testing.T) {

	filename := "higorToCSVNormalExpected.csv"
	dataExpected := [][]string{{"col1", "col2", "col3"}, {"row11", "row12", "row13"}, {"row21", "row22", "row23"}}

	dfResult := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: Book{
			"col1": PageString{"row11", "row21"},
			"col2": PageString{"row12", "row22"},
			"col3": PageString{"row13", "row23"},
		},
	}

	dfResult.ToCSV(filename)

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

// ToCSV With another separator
// ToCSV with or without header
// TOCSV with or without index

//////////////////////////
// TrasposeRows //
////////////////////////
func TestTrasposeRowsMultipleDataType(t *testing.T) {
	dataExpected := [][]string{{"col1", "col2", "col3"}, {"1", "NaN", "row13"}, {"row21", "row22", "row23"}}

	dfExpected := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: Book{
			"col1": PageString{1, "row21"},
			"col2": PageString{math.NaN(), "row22"},
			"col3": PageString{"row13", "row23"},
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
			"col1": PageString{"row11", "row21"},
			"col2": PageString{"row12", "row22"},
			"col3": PageString{"row13", "row23"},
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
			"col1": PageString{"row11", "row21"},
			"col2": PageString{"row12", "row22"},
			"col3": PageString{"row13", "row23"},
		},
	}

	dataResult := dfExpected.GetValues()

	CSVChecker(dataExpected, dataResult, t)

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
		"colInt":    PageFloat64{1, math.NaN(), 2, 3},
		"colString": PageString{"hola", "que", "hace", nil},
		"colBool":   PageBool{nil, true, false, nil},
		"colFloat":  PageFloat64{3.2, 5.4, math.NaN(), math.NaN()},
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
		"colInt2":    PageFloat64{1, math.NaN(), 2, 3},
		"colString2": PageString{"hola", "que", "hace", nil},
		"colBool2":   PageBool{nil, true, false, nil},
		"colFloat2":  PageFloat64{3.2, 5.4, math.NaN(), math.NaN()},
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

// TODO: Print normal DataFrame
// TODO: Print DataFrame with nils values
// TODO: Print DataFrame with multiple DataTypes

func TestPrintDataFrame(t *testing.T) {
	columns := []string{"col1", "col2", "col3"}
	chapters := Book{
		"col1": PageString{"row11", "row21"},
		"col2": PageString{"row12", "row22"},
		"col3": PageBool{true, false},
	}

	df := DataFrame{
		Columns: columns,
		Values:  chapters,
	}

	tableExpectedFormat := "+-------+-------+-------+\n| COL1  | COL2  | COL3  |\n+-------+-------+-------+\n| row11 | row12 | true  |\n| row21 | row22 | false |\n+-------+-------+-------+\n"

	tableResultFormat := df.String()

	if tableExpectedFormat != tableResultFormat {
		t.Errorf("Table format error.\nExpected:\n%v\nResult:\n%v", tableExpectedFormat, tableResultFormat)
	}

}

func TestPrintDataFrameWithNaN(t *testing.T) {
	columns := []string{"col1", "col2", "col3"}
	chapters := Book{
		"col1": PageString{math.NaN(), "row21"},
		"col2": PageString{"row12", "row22"},
		"col3": PageString{"row13", "row23"},
	}

	df := DataFrame{
		Columns: columns,
		Values:  chapters,
	}

	tableExpectedFormat := "+-------+-------+-------+\n| COL1  | COL2  | COL3  |\n+-------+-------+-------+\n| NaN   | row12 | row13 |\n| row21 | row22 | row23 |\n+-------+-------+-------+\n"

	tableResultFormat := df.String()

	if tableExpectedFormat != tableResultFormat {
		t.Errorf("Table format error.\nExpected:\n%v\nResult:\n%v", tableExpectedFormat, tableResultFormat)
	}

}

///////////////////////////
// DataTypes on columns //
/////////////////////////

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
	exportCSV(filename, dataExpected)

	// Read the CSV content
	csvOpen, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer csvOpen.Close()
	csvReader := csv.NewReader(csvOpen)
	dataResult, _ := csvReader.ReadAll()
	CSVChecker(dataExpected, dataResult, t)
	defer os.Remove(filename)
}

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
	dataResult, _ := csvReader.ReadAll()
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
	exportCSV(filename, dataExpected, Sep(separator))

	// Read the CSV content
	csvOpen, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer csvOpen.Close()
	csvReader := csv.NewReader(csvOpen)
	csvReader.Comma = separator
	dataResult, _ := csvReader.ReadAll()
	CSVChecker(dataExpected, dataResult, t)
	defer os.Remove(filename)
}

/////////////////////////////////////
// DataFrame Description functions /
///////////////////////////////////
