package higor

import (
	"reflect"
	"testing"

	"github.com/osmandi/higor/csv"
	"github.com/osmandi/higor/dataframe"
)

func TestVersion(t *testing.T) {
	versionExpected := "v0.6.0"
	versionResult := Version

	if versionExpected != versionResult {
		t.Errorf("Version different. Expected: %s, Result: %s", versionExpected, versionResult)
	}
}

func TestReadCSV(t *testing.T) {
	// TODO: Add ColumnIndex

	// Normal comparation
	inputData := [][]string{{"name", "age"}, {"pepito", "21"}, {"juanito", "22"}, {"pepita", "23"}, {"juanita", "24"}}
	schema := make(dataframe.Schema)
	schema[0] = dataframe.StructString{}
	schema[1] = dataframe.StructInt{}
	dfExpected := dataframe.NewDataFrame(inputData[1:], inputData[0], schema, "")

	dfResult := ReadCSV("csv_examples/simple.csv", schema)
	if !reflect.DeepEqual(dfExpected, dfResult) {
		t.Errorf("Both DataFrame are different but equal expected.\nExpected: %#v\nResult: %#v", dfExpected, dfResult)
	}

	// Custom separator
	dfResult = ReadCSV("csv_examples/sep.csv", schema, csv.Sep(';'))
	if !reflect.DeepEqual(dfExpected, dfResult) {
		t.Errorf("Both DataFrame are different but equal expected.\nExpected: %#v\nResult: %#v", dfExpected, dfResult)
	}

	// Normal NaN
	// TODO: Create test NaN values for each type

	// NaN custom
	// TODO: Create test NaN custom values

	// Without columns
	inputDataWithoutColumns := [][]string{{"pepito", "21"}, {"juanito", "22"}, {"pepita", "23"}, {"juanita", "24"}}
	dfExpectedWithoutColumns := dataframe.NewDataFrame(inputDataWithoutColumns[1:], inputDataWithoutColumns[0], schema, "")

	dfResult = ReadCSV("csv_examples/without_columns.csv", schema)

	if !reflect.DeepEqual(dfExpectedWithoutColumns, dfResult) {
		t.Errorf("Both DataFrame are different but equal expected.\nExpected: %#v\nResult: %#v", dfExpectedWithoutColumns, dfResult)
	}

	// More rows than columns
	// TODO: Create test for more rows than columns
}

// TODO: Implement concurrency and test
// TODO: Create errors
