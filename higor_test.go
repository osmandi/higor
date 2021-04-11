package higor

import (
	"os"
	"testing"

	df "github.com/osmandi/higor/dataframe"
)

func TestHelloHigor(t *testing.T) {

	resultMessage := HelloHigor()
	expectedMessage := "Hello from Higor :) v0.2.1"

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
	csvTempFile := df.CSVCreatorMock(data, separator)
	csvTempFilename := csvTempFile.Name()
	defer os.Remove(csvTempFilename)

	typeColumnsExpected := df.Words{
		"col1": df.Letter{"s": 2},
		"col2": df.Letter{"s": 2},
		"col3": df.Letter{"s": 2},
	}

	dfExpected := df.DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: df.Book{
			"col1": {"row11", "row21"},
			"col2": {"row12", "row22"},
			"col3": {"row13", "row23"},
		},
		DataType: typeColumnsExpected,
	}

	dfResult := ReadCSV(csvTempFilename)

	df.DataFrameChecker(dfExpected, dfResult, t)

}

func TestReadCSVAnoterSeparator(t *testing.T) {
	// Mockup
	data := [][]string{{"col1", "col2", "col3"}, {"row11", "row12", "row13"}, {"row21", "row22", "row23"}}
	separator := '|'
	csvTempFile := df.CSVCreatorMock(data, separator)
	csvTempFilename := csvTempFile.Name()
	defer os.Remove(csvTempFilename)

	typeColumnsExpected := df.Words{
		"col1": df.Letter{"s": 2},
		"col2": df.Letter{"s": 2},
		"col3": df.Letter{"s": 2},
	}

	dfExpected := df.DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: df.Book{
			"col1": {"row11", "row21"},
			"col2": {"row12", "row22"},
			"col3": {"row13", "row23"},
		},
		DataType: typeColumnsExpected,
	}

	dfResult := ReadCSV(csvTempFilename, df.Sep('|'))

	df.DataFrameChecker(dfExpected, dfResult, t)

}
