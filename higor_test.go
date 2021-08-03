package higor

import (
	"os/exec"
	"reflect"
	"strings"
	"testing"

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
	inputData := [][]string{[]string{"name", "age"}, []string{"pepito", "21"}, []string{"juanito", "22"}, []string{"pepita", "23"}, []string{"juanita", "24"}}
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

	// TODO: Custom Sep comparation
	// TODO: LazyQuotes comparation
	// TODO: Custom NaN comparation
	// TODO: Custom Datetime comparation
	// TODO: Without columns
	// TODO: More rows than columns

}
