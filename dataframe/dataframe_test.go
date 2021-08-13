package dataframe

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestWriteWordString(t *testing.T) {

	textInput := "Hello Higor"

	// Normal input
	wordStringResult := WriteWordString(textInput)
	wordStringExpected := WordString{value: textInput}

	if wordStringResult.value != wordStringExpected.value {
		t.Errorf("Error wordString. Expected: %+v Result: %+v", wordStringExpected, wordStringResult)
	}
}

func TestWriteWordBool(t *testing.T) {
	textInputFalse := "false"
	textInputTrue := "true"

	// Input true
	wordBoolResultTrue := WriteWordBool(textInputTrue)
	wordBoolExpectedTrue := WordBool{value: true}

	if !reflect.DeepEqual(wordBoolResultTrue, wordBoolExpectedTrue) {
		t.Errorf("Both words are different but equal expected. Expected: %+v, Result: %+v", wordBoolExpectedTrue, wordBoolResultTrue)
	}

	// Input false
	wordBoolResultFalse := WriteWordBool(textInputFalse)
	wordBoolExpectedFalse := WordBool{value: false}

	if !reflect.DeepEqual(wordBoolResultFalse, wordBoolExpectedFalse) {
		t.Errorf("Both words are different but equal expected. Expected: %+v, Result: %+v", wordBoolExpectedFalse, wordBoolResultFalse)
	}

}

func TestWriteLine(t *testing.T) {

	nanLayout := ""
	layoutDatetime := "2006-01-02"

	// Input with String
	var1 := "Higor"
	var2 := "Higor2"
	var3 := "Higor3"
	var4NaN := nanLayout

	inputLine := []string{var1, var2, var3, var4NaN}
	lineExpected := Lines{WordString{value: var1}, WordString{value: var2}, WordString{value: var3}, WordNaN{}}

	lineResult := WriteLine(inputLine, nanLayout, layoutDatetime)

	if !reflect.DeepEqual(lineExpected, lineResult) {
		t.Errorf("Both lines are different but equal expected.\nExpected: %v\nResult: %v", lineExpected, lineResult)
	}

	// All values
	inputLine2 := []string{"Higor", "1", "2.2", "false", "", "2020-01-01"}
	datetime2, _ := time.Parse(layoutDatetime, "2020-01-01")
	lineExpected2 := Lines{WordString{value: "Higor"}, WordFloat64{value: float64(1)}, WordFloat64{value: float64(2.2)}, WordBool{value: false}, WordNaN{}, WordDatetime{value: datetime2}}
	lineResult2 := WriteLine(inputLine2, nanLayout, layoutDatetime)

	if !reflect.DeepEqual(lineExpected2, lineResult2) {
		t.Errorf("Both lines are different but equal expected.\nExpected: %v\nResult: %v", lineExpected2, lineResult2)
	}

}

func TestTranslateWord(t *testing.T) {
	nanLayout := ""
	layoutDatetime := "2006-01-02"
	variables := []string{nanLayout, "Higor", "1", "1.2", "true", "True", "False", "false", "2020-02-01"}
	result := []string{"NaN", "string", "float64", "float64", "bool", "bool", "bool", "bool", "datetime"}

	for i, v := range variables {
		trans, _ := translateWord(v, nanLayout, layoutDatetime)

		if trans != result[i] {
			t.Errorf("Both different but equal result. Expected: %s, Result: %s", result[i], trans)
		}
	}

}

func TestAddLine(t *testing.T) {

	dfExpected := NewDataFrame()
	dfResult := NewDataFrame()

	datetime, _ := time.Parse("2006-01-02", "2020-01-01")
	inputLine := []string{"Higor", "1", "2.2", "false", "", "2020-01-01"}
	lineExpected := Lines{WordString{value: "Higor"}, WordFloat64{value: float64(1)}, WordFloat64{value: float64(2.2)}, WordBool{value: false}, WordNaN{}, WordDatetime{value: datetime}}

	dfExpected.Values = Book{lineExpected}

	dfResult.AddLine(inputLine)

	if !reflect.DeepEqual(dfExpected, dfResult) {
		t.Errorf("Dataframes different but equal expected.\nExpected: %+v\nResult: %+v", dfExpected, dfResult)
	}

}

func TestNewDataFrame(t *testing.T) {
	dfExpected := DataFrame{}
	dfExpected.NaNLayout = ""
	dfExpected.DatetimeLayout = "2006-01-02"

	dfResult := NewDataFrame()

	if !reflect.DeepEqual(dfExpected, dfResult) {
		t.Errorf("Dataframes different but equal expected.\nExpected: %+v\nResult: %+v", dfExpected, dfResult)
	}

}

func TestString(t *testing.T) {
	df := NewDataFrame()
	input := [][]string{{"NAME", "AGE"}, {"pepito", "21"}, {"juanito", "22"}, {"pepita", "2.3"}, {"juanita", ""}}
	df.Columns = input[0]
	df.Shape = [2]int{4, 2}
	for _, v := range input[1:] {
		df.AddLine(v)
	}

	expected := `+---------+-----+
|  NAME   | AGE |
+---------+-----+
| pepito  | 21  |
| juanito | 22  |
| pepita  | 2.3 |
| juanita | NaN |
+---------+-----+
`
	result := fmt.Sprint(df)

	if expected != result {
		t.Errorf("Dataframe Print Different.\nExpected:\n%s\nResult:\n%s", expected, result)
	}
}

func TestHead(t *testing.T) {
	df := NewDataFrame()
	input := [][]string{
		{"name", "age"},
		{"pepito", "21"},
		{"juanito", "22"},
		{"pepita", "2.3"},
		{"juanita", ""},
		{"pepito", "21"},
		{"juanito", "22"},
		{"pepita", "2.3"},
		{"juanita", ""},
		{"pepita", "2.3"},
		{"juanita", ""},
		{"pepito", "21"},
		{"juanito", "22"},
		{"pepita", "2.3"},
		{"juanita", ""},
	}
	df.Columns = input[0]
	df.Shape = [2]int{14, 2}
	for _, v := range input[1:] {
		df.AddLine(v)
	}

	dfHead := df.Head()
	dfHeadExpected := df
	dfHeadExpected.Values = df.Values[:10]
	dfHeadExpected.Shape[0] = 10

	if !reflect.DeepEqual(dfHeadExpected, dfHead) {
		t.Errorf("DataFrame different but equal expected.")
	}

	dfHead5 := df.Head(5)
	dfHeadExpected5 := df
	dfHeadExpected5.Values = df.Values[:5]
	dfHeadExpected5.Shape[0] = 5

	if !reflect.DeepEqual(dfHead5, dfHeadExpected5) {
		t.Errorf("DataFrame different but equal expected.\nExpected:%+v\nResult:\n%+v", dfHeadExpected5, dfHead5)
	}

	// Less than 10 rows
	dfLess := NewDataFrame()
	input = [][]string{
		{"name", "age"},
		{"pepito", "21"},
		{"juanito", "22"},
		{"pepita", "2.3"},
		{"juanita", ""},
	}
	dfLess.Columns = input[0]
	dfLess.Shape = [2]int{4, 2}
	for _, v := range input[1:] {
		dfLess.AddLine(v)
	}

	dfLessHead := dfLess.Head()
	if !reflect.DeepEqual(dfLess, dfLessHead) {
		t.Errorf("Dataframes different but equal expected")
	}
}

func TestTail(t *testing.T) {
	df := NewDataFrame()
	input := [][]string{
		{"name", "age"},
		{"pepito", "21"},
		{"juanito", "22"},
		{"pepita", "2.3"},
		{"juanita", ""},
		{"pepito", "21"},
		{"juanito", "22"},
		{"pepita", "2.3"},
		{"juanita", ""},
		{"pepita", "2.3"},
		{"juanita", ""},
		{"pepito", "21"},
		{"juanito", "22"},
		{"pepita", "2.3"},
		{"juanita", ""},
	}
	df.Columns = input[0]
	df.Shape = [2]int{4, 2}
	for _, v := range input[1:] {
		df.AddLine(v)
	}

	dfTail := df.Tail()
	dfTailExpected := df
	dfTailExpected.Values = df.Values[4:]
	dfTailExpected.Shape[0] = 10

	if !reflect.DeepEqual(dfTail, dfTailExpected) {
		t.Errorf("DataFrame different but equal expected.\nExpected:\n%+v\nResult:\n%+v", dfTailExpected, dfTail)
	}

	dfTail5 := df.Tail(5)
	dfTailExpected5 := df
	dfTailExpected5.Values = df.Values[9:]
	dfTailExpected5.Shape[0] = 5

	if !reflect.DeepEqual(dfTail5, dfTailExpected5) {
		t.Errorf("DataFrame different but equal expected.\nExpected:\n%+v\nResult:\n%+v", dfTailExpected5, dfTail5)
	}

	// Less than 10 rows
	dfLess := NewDataFrame()
	input = [][]string{
		{"name", "age"},
		{"pepito", "21"},
		{"juanito", "22"},
		{"pepita", "2.3"},
		{"juanita", ""},
	}

	dfLess.Columns = input[0]
	dfLess.Shape = [2]int{4, 2}
	for _, v := range input[1:] {
		dfLess.AddLine(v)
	}

	dfLessTail := dfLess.Tail()
	if !reflect.DeepEqual(dfLess, dfLessTail) {
		t.Errorf("Dataframes different but equal expected")
	}
}

// TODO: Select columns
func TestSelect(t *testing.T) {
	// Base
	df := NewDataFrame()
	input := [][]string{
		{"name", "age", "data"},
		{"pepito", "21", "true"},
		{"juanito", "22", "false"},
		{"pepita", "2.3", "true"},
		{"juanita", "", "false"},
	}
	df.Columns = input[0]
	df.Shape = [2]int{4, 3}
	for _, v := range input[1:] {
		df.AddLine(v)
	}

	// Select two columns
	dfSelect := NewDataFrame()
	input = [][]string{
		{"name", "age"},
		{"pepito", "21"},
		{"juanito", "22"},
		{"pepita", "2.3"},
		{"juanita", ""},
	}
	dfSelect.Columns = input[0]
	dfSelect.Shape = [2]int{4, 2}
	for _, v := range input[1:] {
		dfSelect.AddLine(v)
	}

	dfSelected2 := df.Select("name", "age")
	dfSelected2.Shape[1] = 2

	if !reflect.DeepEqual(dfSelected2, dfSelect) {
		t.Errorf("Dataframes different but equal expected.\nExpected:\n%+v\nResult:\n%+v", dfSelect, dfSelected2)
	}

	// Select one columns
	dfSelect1 := NewDataFrame()
	input = [][]string{
		{"name"},
		{"pepito"},
		{"juanito"},
		{"pepita"},
		{"juanita"},
	}
	dfSelect1.Columns = input[0]
	dfSelect1.Shape = [2]int{4, 1}
	for _, v := range input[1:] {
		dfSelect1.AddLine(v)
	}

	dfSelected1 := df.Select("name")
	dfSelected1.Shape[1] = 1

	if !reflect.DeepEqual(dfSelected1, dfSelect1) {
		t.Errorf("Dataframe different but equal expected")
	}

}

func TestFindIndex(t *testing.T) {
	listBase := []string{"col1", "col2", "col3"}
	listFind := []string{"col1", "col3", "col2"}
	expected := []int{0, 2, 1}
	result := findIndex(listBase, listFind)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Index Failed.\nExpected:\n%v\nResult:\n%v", expected, result)
	}

}

// TODO: Drop columns
// TODO: Filter columns
// TODO: Create columns
