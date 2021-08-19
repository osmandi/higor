package higor

import (
	"os/exec"
	"reflect"
	"strings"
	"testing"

	"github.com/osmandi/higor/csv"
	"github.com/osmandi/higor/dataframe"
)

func TestVersion(t *testing.T) {
	c, b := exec.Command("git", "branch", "--show-current"), new(strings.Builder)
	c.Stdout = b
	c.Run()
	versionExpected := strings.TrimRight(b.String(), "\n")
	versionResult := Version

	if versionExpected != versionResult {
		t.Errorf("Version different. Expected: %s, Result: %s", versionExpected, versionResult)
	}
}

func TestReadCSV(t *testing.T) {
	// TODO: Add ColumnIndex

	// Normal comparation
	inputData := [][]string{{"name", "age"}, {"pepito", "21"}, {"juanito", "22"}, {"pepita", "23"}, {"juanita", "24"}}
	dfExpected := dataframe.NewDataFrame(inputData)

	dfResult := ReadCSV("csv_examples/simple.csv")
	if !reflect.DeepEqual(dfExpected, dfResult) {
		t.Errorf("Both DataFrame are different but equal expected.\nExpected: %+v\nResult: %+v", dfExpected, dfResult)
	}

	// Custom separator
	dfResult = ReadCSV("csv_examples/sep.csv", csv.Sep(';'))
	if !reflect.DeepEqual(dfExpected, dfResult) {
		t.Errorf("Both DataFrame are different but equal expected.\nExpected: %+v\nResult: %+v", dfExpected, dfResult)
	}

	// Normal NaN
	inputDataNaN := [][]string{{"name", "age"}, {"pepito", "21"}, {"", "22"}, {"pepita", ""}, {"juanita", "24"}}
	dfExpectedNaN := dataframe.NewDataFrame(inputDataNaN)

	dfResult = ReadCSV("csv_examples/nan.csv")

	if !reflect.DeepEqual(dfExpectedNaN, dfResult) {
		t.Errorf("Both DataFrame are different but equal expected.\nExpected: %+v\nResult: %+v", dfExpectedNaN, dfResult)
	}

	// NaN custom
	inputDataNaNCustom := [][]string{{"name", "age"}, {"pepito", "21"}, {"None", "22"}, {"pepita", "None"}, {"juanita", "24"}}
	dfExpectedNaNCustom := dataframe.NewDataFrame(inputDataNaNCustom)

	dfResult = ReadCSV("csv_examples/nan_custom.csv", csv.NaNLayout("None"))

	if !reflect.DeepEqual(dfExpectedNaNCustom, dfResult) {
		t.Errorf("Both DataFrame are different but equal expected.\nExpected: %+v\nResult: %+v", dfExpectedNaNCustom, dfResult)
	}

	// Without columns
	inputDataWithoutColumns := [][]string{{"pepito", "21"}, {"juanito", "22"}, {"pepita", "23"}, {"juanita", "24"}}
	dfExpectedWithoutColumns := dataframe.NewDataFrame(inputDataWithoutColumns)

	dfResult = ReadCSV("csv_examples/without_columns.csv")

	if !reflect.DeepEqual(dfExpectedWithoutColumns, dfResult) {
		t.Errorf("Both DataFrame are different but equal expected.\nExpected: %+v\nResult: %+v", dfExpectedWithoutColumns, dfResult)
	}

	// More rows than columns
	// TODO: Create errors
	inputDataMoreRowsThanColumns := [][]string{{"name", "age"}, {"pepito", "21"}, {"juanito", "22"}, {"pepita", "23"}, {"juanita", "24"}}
	dfExpectedMoreRowsThanColumns := dataframe.NewDataFrame(inputDataMoreRowsThanColumns)

	dfResult = ReadCSV("csv_examples/more_columns_than_rows.csv")

	if !reflect.DeepEqual(dfExpectedMoreRowsThanColumns, dfResult) {
		t.Errorf("Both DataFrame are different but equal expected.\nExpected: %+v\nResult: %+v", dfExpectedMoreRowsThanColumns, dfResult)
	}
}

// TODO: Implement concurrency and test
