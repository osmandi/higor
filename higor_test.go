package higor

import (
	"os"
	"testing"
	"time"

	"github.com/osmandi/higor/dataframe"
)

func TestHelloHigor(t *testing.T) {

	resultMessage := HelloHigor()
	expectedMessage := "Hello from Higor :) v0.3.0"

	if resultMessage != expectedMessage {
		t.Errorf("Message expected: '%s' but received: '%s'", expectedMessage, resultMessage)
	}

}

/////////////
// ReadCSV /
///////////

func TestReadCSVNormal(t *testing.T) {
	// Mockup
	data := [][]string{{"col1", "col2", "col3"}, {"row11", "row12", "row13"}, {"row21", "row22", "row23"}}
	separator := ','
	csvTempFile := dataframe.CSVCreatorMock(data, separator)
	csvTempFilename := csvTempFile.Name()
	defer os.Remove(csvTempFilename)

	dfExpected := dataframe.DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: dataframe.Book{
			dataframe.PageString{"row11", "row21"},
			dataframe.PageString{"row12", "row22"},
			dataframe.PageString{"row13", "row23"},
		},
	}
	schema := dataframe.Book{
		dataframe.PageString{},
		dataframe.PageString{},
		dataframe.PageString{},
	}
	dfResult := ReadCSV(csvTempFilename, dataframe.Schema(schema))

	dataframe.DataFrameChecker(dfExpected, dfResult, t)

}

func TestReadCSVAnoterSeparator(t *testing.T) {
	// Mockup
	data := [][]string{{"col1", "col2", "col3"}, {"row11", "row12", "row13"}, {"row21", "row22", "row23"}}
	separator := '|'
	csvTempFile := dataframe.CSVCreatorMock(data, separator)
	csvTempFilename := csvTempFile.Name()
	defer os.Remove(csvTempFilename)

	dfExpected := dataframe.DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: dataframe.Book{
			dataframe.PageString{"row11", "row21"},
			dataframe.PageString{"row12", "row22"},
			dataframe.PageString{"row13", "row23"},
		},
	}
	schema := dataframe.Book{
		dataframe.PageString{},
		dataframe.PageString{},
		dataframe.PageString{},
	}
	dfResult := ReadCSV(csvTempFilename, dataframe.Sep('|'), dataframe.Schema(schema))

	dataframe.DataFrameChecker(dfExpected, dfResult, t)

}

func TestReadCSVMultipleDataTypes(t *testing.T) {
	// Mockup
	data := [][]string{{"colString", "colBool", "colFloat64", "colAny"}, {"rowString", "true", "1", "uno"}, {"rowString", "false", "3.2", "false"}}
	separator := '|'
	csvTempFile := dataframe.CSVCreatorMock(data, separator)
	csvTempFilename := csvTempFile.Name()
	defer os.Remove(csvTempFilename)

	dfExpected := dataframe.DataFrame{
		Columns: []string{"colString", "colBool", "colFloat64", "colAny"},
		Values: dataframe.Book{
			dataframe.PageString{"rowString", "rowString"},
			dataframe.PageBool{true, false},
			dataframe.PageFloat64{1, 3.2},
			dataframe.PageAny{"uno", "false"},
		},
	}
	schema := dataframe.Book{
		dataframe.PageString{},
		dataframe.PageBool{},
		dataframe.PageFloat64{},
		dataframe.PageAny{},
	}
	dfResult := ReadCSV(csvTempFilename, dataframe.Sep('|'), dataframe.Schema(schema))

	dataframe.DataFrameChecker(dfExpected, dfResult, t)

}

func TestReadCSVWithDatatimeType(t *testing.T) {
	// Mockup
	data := [][]string{{"colDatetime"}, {"2020-01-02"}, {"2020-01-30"}}
	separator := '|'
	csvTempFile := dataframe.CSVCreatorMock(data, separator)
	csvTempFilename := csvTempFile.Name()
	defer os.Remove(csvTempFilename)
	layout := "2006-01-02"
	date1, _ := time.Parse(layout, "2020-01-02")
	date2, _ := time.Parse(layout, "2020-01-30")
	dfExpected := dataframe.DataFrame{
		Columns: []string{"colDatetime"},
		Values: dataframe.Book{
			dataframe.PageDatetime{date1, date2},
		},
	}
	schema := dataframe.Book{
		dataframe.PageDatetime{},
	}
	dfResult := ReadCSV(csvTempFilename, dataframe.Sep('|'), dataframe.Schema(schema))

	dataframe.DataFrameChecker(dfExpected, dfResult, t)
}

// TestReadCSVWithTimestamp
// TestReadCSVWithMultipleDatetime on Multiple columns
