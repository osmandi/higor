package dataframe

import (
	"os"
	"testing"
)

//////////////////////////
// Dataframe functions //
////////////////////////

// Equal DataFrame
func TestEqualDataFrame(t *testing.T) {

	columns := []string{"colInt", "colString", "colBool", "colFloat"}
	chapters := book{
		"colInt":    {1, nil, 2, 3},
		"colString": {"hola", "que", "hace", nil},
		"colBool":   {nil, true, false, nil},
		"colFloat":  {3.2, 5.4, nil, nil},
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

// Print a prety DataFrame

// Diferent DataFrame
func TestDifferentlDataFrame(t *testing.T) {
	columns := []string{"colInt", "colString", "colBool", "colFloat"}
	chapters := book{
		"colInt2":    {1, nil, 2, 3},
		"colString2": {"hola", "que", "hace", nil},
		"colBool2":   {nil, true, false, nil},
		"colFloat2":  {3.2, 5.4, nil, nil},
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

/////////
// CSV /
///////

func dataFrameChecker(dfExpected, dfResult DataFrame, t *testing.T) {
	isEqual := IsEqual(dfExpected, dfResult)
	if !isEqual {
		t.Errorf("dfExpected and dfResult are distinct.\ndfExpected: %v \ndfResult: %v", dfExpected, dfResult)
	}

}

// Read a normal DataFrame

func TestReadCSVNormal(t *testing.T) {
	// Mockup
	data := [][]string{{"col1", "col2", "col3"}, {"row11", "row12", "row13"}, {"row21", "row22", "row23"}}
	separator := ','
	csvTempFile := csvCreatorMock(data, separator)
	csvTempFilename := csvTempFile.Name()
	defer os.Remove(csvTempFilename)

	dfExpected := DataFrame{
		Columns: []string{"col1", "col2", "col3"},
		Values: book{
			"col1": {"row11", "row21"},
			"col2": {"row12", "row22"},
			"col3": {"row13", "row23"},
		},
	}

	dfResult := ReadCSV(csvTempFilename)

	dataFrameChecker(dfExpected, dfResult, t)

}

// Read a DataFrame with another separator

// Read a DataFrame that doesn't exists

// Read a DataFrame with LazyQuotes

// Read a DataFrame with new lines

// Autocomplete the header to header empty
