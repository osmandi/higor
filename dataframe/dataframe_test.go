package dataframe

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"math"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestErrorChecker(t *testing.T) {
	ErrorChecker(nil)
}

func TestErrorSchema(t *testing.T) {
	ErrorSchema("", "", nil, nil)
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
			PageString{"row11", "row21"},
			PageString{"row12", "row22"},
			PageString{"row13", "row23"},
		},
		Shape: [2]int{2, 3},
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

/////////////////////////
// TrasposeRows //
////////////////////////
func TestTrasposeRowsMultipleDataType(t *testing.T) {
	dataExpected := [][]string{{"1", "NaN", "row13"}, {"row21", "row22", "row23"}}

	dfExpected := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: Book{
			PageAny{1, "row21"},
			PageAny{math.NaN(), "row22"},
			PageString{"row13", "row23"},
		},
		Shape: [2]int{2, 3},
	}

	dataResult := trasposeRows(dfExpected)

	CSVChecker(dataExpected, dataResult, t)

}
func TestTrasposeRowsString(t *testing.T) {
	dataExpected := [][]string{{"row11", "row12", "row13"}, {"row21", "row22", "row23"}}

	dfExpected := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: Book{
			PageString{"row11", "row21"},
			PageString{"row12", "row22"},
			PageString{"row13", "row23"},
		},
		Shape: [2]int{2, 3},
	}

	dataResult := trasposeRows(dfExpected)

	CSVChecker(dataExpected, dataResult, t)

}

func TestValuesNormal(t *testing.T) {
	dataExpected := [][]string{{"row11", "row12", "row13"}, {"row21", "row22", "row23"}}

	dfExpected := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: Book{
			PageString{"row11", "row21"},
			PageString{"row12", "row22"},
			PageString{"row13", "row23"},
		},
		Shape: [2]int{2, 3},
	}

	dataResult := dfExpected.GetValues()

	CSVChecker(dataExpected, dataResult, t)

}

func TestSep(t *testing.T) {
	separator := ';'
	csvResult := &CSV{}
	csvOptionInternal := Sep(separator)
	csvOptionInternal(csvResult)

	if csvResult.Sep != separator {
		t.Errorf("Sep error. Expected: ';'. But result: %v", csvResult.Sep)
	}

}

func TestSchema(t *testing.T) {
	schema := Book{
		PageString{},
		PageBool{},
		PageFloat64{},
	}
	csvResult := &CSV{}
	csvOptionInternal := Schema(schema)
	csvOptionInternal(csvResult)

	if !reflect.DeepEqual(schema, csvResult.Schema) {
		t.Errorf("Schema error. Expected: %v. But result: %v", schema, csvResult.Sep)
	}

}

func TestNone(t *testing.T) {
	none := "none"
	csvResult := &CSV{}
	csvOptionInternal := None(none)
	csvOptionInternal(csvResult)

	if csvResult.None != none {
		t.Errorf("None error. Expected: %v - But result: %v", none, csvResult.None)
	}
}

func TestDataformat(t *testing.T) {
	dateformat := "YYYY-MM-DD"
	csvResult := &CSV{}
	csvOptionInternal := Dateformat(dateformat)
	csvOptionInternal(csvResult)

	if !reflect.DeepEqual(dateformat, csvResult.Dateformat) {
		t.Errorf("Schema error. Expected: %v. But result: %v", dateformat, csvResult.Dateformat)
	}

}

func TestEqualDataFrame(t *testing.T) {

	columns := []string{"colInt", "colString", "colBool", "colFloat"}
	chapters := Book{
		PageFloat64{1, math.NaN(), 2, 3},
		PageAny{"hola", "que", "hace", math.NaN()},
		PageAny{math.NaN(), true, false, math.NaN()},
		PageFloat64{3.2, 5.4, math.NaN(), math.NaN()},
	}
	df1 := DataFrame{
		Columns: columns,
		Values:  chapters,
	}
	df2 := DataFrame{
		Columns: columns,
		Values:  chapters,
	}

	isEqual, message := IsEqual(df1, df2)

	if !isEqual {
		t.Errorf("Error equalDataframe. df1 and df2 are different! But equal expected!: %s", message)
	}
}

// Diferent DataFrame
func TestDifferentlDataFrame(t *testing.T) {
	columns := []string{"colInt", "colString", "colBool", "colFloat"}
	chapters := Book{
		PageFloat64{1, math.NaN(), 2, 3},
		PageAny{"hola", "que", "hace", math.NaN()},
		PageAny{math.NaN(), true, false, math.NaN()},
		PageFloat64{3.2, 5.4, math.NaN(), math.NaN()},
	}

	df1 := DataFrame{
		Columns: columns,
		Values:  chapters,
	}
	df2 := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values:  chapters,
	}

	isEqual, _ := IsEqual(df1, df2)

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
		PageString{"row11", "row21"},
		PageString{"row12", "row22"},
		PageBool{true, false},
	}

	df := DataFrame{
		Columns: columns,
		Values:  chapters,
		Shape:   [2]int{2, 3},
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
		PageAny{math.NaN(), "row21"},
		PageString{"row12", "row22"},
		PageString{"row13", "row23"},
	}

	df := DataFrame{
		Columns: columns,
		Values:  chapters,
		Shape:   [2]int{2, 3},
	}

	tableExpectedFormat := "+-------+-------+-------+\n| COL1  | COL2  | COL3  |\n+-------+-------+-------+\n| NaN   | row12 | row13 |\n| row21 | row22 | row23 |\n+-------+-------+-------+\n"

	tableResultFormat := df.String()

	if tableExpectedFormat != tableResultFormat {
		t.Errorf("Table format error.\nExpected:\n%v\nResult:\n%v", tableExpectedFormat, tableResultFormat)
	}

}

func TestIsNaNPageDatetimeFalse(t *testing.T) {
	layout := "2006-01-02"
	dateNaNFalse, _ := time.Parse(layout, "2020-01-02")
	//dateNaNTrue := time.Date(0001, 1, 1, 0, 0, 0, 0, time.UTC)
	resultExpectedFalse := IsNaN(dateNaNFalse)

	if resultExpectedFalse != false {
		t.Errorf("IsNaN Error. Expected: False | but Result: %t", resultExpectedFalse)
	}

}

func TestIsNaNPageDatetimeTrue(t *testing.T) {
	dateNaNTrue := time.Date(0001, 1, 1, 0, 0, 0, 0, time.UTC)
	resultExpectedTrue := IsNaN(dateNaNTrue)

	if resultExpectedTrue != true {
		t.Errorf("IsNaN Error. Expected: True | but Result: %t", resultExpectedTrue)
	}

}

func TestIsNaNPageStringFalse(t *testing.T) {
	stringNaNFalse := "isFalse"
	resultExpectedFalse := IsNaN(stringNaNFalse)
	if resultExpectedFalse != false {
		t.Errorf("IsNaN Error. Expected: False | but Result: %t", resultExpectedFalse)
	}
}

func TestIsNaNPageStringTrue(t *testing.T) {
	stringNaNTrue := ""
	resultExpectedTrue := IsNaN(stringNaNTrue)
	if resultExpectedTrue != true {
		t.Errorf("IsNaN Error. Expected: true | but Result: %t", resultExpectedTrue)
	}
}

func TestPrintDataFrameDateWithNaN(t *testing.T) {
	columns := []string{"colDate", "colString1", "colString2"}
	layout := "2006-01-02"
	date1, _ := time.Parse(layout, "2020-01-02")
	dateNaN := time.Date(0001, 1, 1, 0, 0, 0, 0, time.UTC)
	chapters := Book{
		PageDatetime{date1, dateNaN},
		PageString{"hola1", "hola2"},
		PageString{"hola1", "hola2"},
	}

	df := DataFrame{
		Columns: columns,
		Values:  chapters,
		Shape:   [2]int{2, 3},
	}
	tableExpectedFormat := "+-------------------------------+------------+------------+\n|            COLDATE            | COLSTRING1 | COLSTRING2 |\n+-------------------------------+------------+------------+\n| 2020-01-02 00:00:00 +0000 UTC | hola1      | hola1      |\n| NaN                           | hola2      | hola2      |\n+-------------------------------+------------+------------+\n"

	tableResultFormat := df.String()

	if tableExpectedFormat != tableResultFormat {
		t.Errorf("Table format error.\nExpected:\n%v\nResult:\n%v", tableExpectedFormat, tableResultFormat)
	}

}

func TestPrintPageStringWithNaN(t *testing.T) {
	columns := []string{"colString1", "colString2", "colstring3"}
	chapters := Book{
		PageString{"hola1", ""},
		PageString{"", "hola2"},
		PageString{"", ""},
	}

	df := DataFrame{
		Columns: columns,
		Values:  chapters,
		Shape:   [2]int{2, 3},
	}
	tableExpectedFormat := "+------------+------------+------------+\n| COLSTRING1 | COLSTRING2 | COLSTRING3 |\n+------------+------------+------------+\n| hola1      | NaN        | NaN        |\n| NaN        | hola2      | NaN        |\n+------------+------------+------------+\n"

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
