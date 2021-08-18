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

	// Normal comparation
	inputData := [][]string{{"name", "age"}, {"pepito", "21"}, {"juanito", "22"}, {"pepita", "23"}, {"juanita", "24"}}
	dfExpected := dataframe.NewDataFrame()
	dfExpected.Columns = inputData[0]
	dfExpected.Shape = [2]int{4, 2}

	for _, v := range inputData[1:] {
		dfExpected.AddLine(v)
	}

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
	dfExpectedNaN := dataframe.NewDataFrame()
	dfExpectedNaN.Columns = inputDataNaN[0]
	dfExpectedNaN.Shape = [2]int{4, 2}
	for _, v := range inputDataNaN[1:] {
		dfExpectedNaN.AddLine(v)
	}

	dfResult = ReadCSV("csv_examples/nan.csv")

	if !reflect.DeepEqual(dfExpectedNaN, dfResult) {
		t.Errorf("Both DataFrame are different but equal expected.\nExpected: %+v\nResult: %+v", dfExpectedNaN, dfResult)
	}

	// NaN custom
	inputDataNaNCustom := [][]string{{"name", "age"}, {"pepito", "21"}, {"None", "22"}, {"pepita", "None"}, {"juanita", "24"}}
	dfExpectedNaNCustom := dataframe.NewDataFrame()
	dfExpectedNaNCustom.Columns = inputDataNaNCustom[0]
	dfExpectedNaNCustom.Shape = [2]int{4, 2}
	dfExpectedNaNCustom.NaNLayout = "None"
	for _, v := range inputDataNaNCustom[1:] {
		dfExpectedNaNCustom.AddLine(v)
	}

	dfResult = ReadCSV("csv_examples/nan_custom.csv", csv.NaNLayout("None"))

	if !reflect.DeepEqual(dfExpectedNaNCustom, dfResult) {
		t.Errorf("Both DataFrame are different but equal expected.\nExpected: %+v\nResult: %+v", dfExpectedNaNCustom, dfResult)
	}

	// Without columns
	inputDataWithoutColumns := [][]string{{"pepito", "21"}, {"juanito", "22"}, {"pepita", "23"}, {"juanita", "24"}}
	dfExpectedWithoutColumns := dataframe.NewDataFrame()
	dfExpectedWithoutColumns.Columns = inputDataWithoutColumns[0]
	dfExpectedWithoutColumns.Shape = [2]int{3, 2}
	for _, v := range inputDataWithoutColumns[1:] {
		dfExpectedWithoutColumns.AddLine(v)
	}

	dfResult = ReadCSV("csv_examples/without_columns.csv")

	if !reflect.DeepEqual(dfExpectedWithoutColumns, dfResult) {
		t.Errorf("Both DataFrame are different but equal expected.\nExpected: %+v\nResult: %+v", dfExpectedWithoutColumns, dfResult)
	}

	// More rows than columns
	// TODO: Create errors
	inputDataMoreRowsThanColumns := [][]string{{"name", "age"}, {"pepito", "21"}, {"juanito", "22"}, {"pepita", "23"}, {"juanita", "24"}}
	dfExpectedMoreRowsThanColumns := dataframe.NewDataFrame()
	dfExpectedMoreRowsThanColumns.Columns = inputDataMoreRowsThanColumns[0]
	dfExpectedMoreRowsThanColumns.Shape = [2]int{4, 2}
	for _, v := range inputDataMoreRowsThanColumns[1:] {
		dfExpectedMoreRowsThanColumns.AddLine(v)
	}

	dfResult = ReadCSV("csv_examples/more_columns_than_rows.csv")

	if !reflect.DeepEqual(dfExpectedMoreRowsThanColumns, dfResult) {
		t.Errorf("Both DataFrame are different but equal expected.\nExpected: %+v\nResult: %+v", dfExpectedMoreRowsThanColumns, dfResult)
	}
}

// TODO: Implement concurrency and test
