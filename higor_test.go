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
	dfExpectedNaN.Columns = inputData[0]
	dfExpectedNaN.Shape = [2]int{4, 2}
	for _, v := range inputDataNaN[1:] {
		dfExpectedNaN.AddLine(v)
	}
	dfResult = ReadCSV("csv_examples/nan.csv")
	if !reflect.DeepEqual(dfExpectedNaN, dfResult) {
		t.Errorf("Both DataFrame are different but equal expected.\nExpected: %+v\nResult: %+v", dfExpectedNaN, dfResult)
	}

	// TODO: LazyQuotes comparation
	// TODO: Custom NaN comparation
	// TODO: Custom Datetime comparation
	// TODO: Without columns
	// TODO: More rows than columns

}
