package higor

import (
	"os"
	"testing"

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
