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
	dfExpected := NewDataFrame(nil, "")
	dfResult := NewDataFrame(nil, "")

	datetime, _ := time.Parse("2006-01-02", "2020-01-01")
	inputLine := []string{"Higor", "1", "2.2", "false", "", "2020-01-01"}
	lineExpected := Lines{WordString{value: "Higor"}, WordFloat64{value: float64(1)}, WordFloat64{value: float64(2.2)}, WordBool{value: false}, WordNaN{}, WordDatetime{value: datetime}}

	dfExpected.Values = Book{lineExpected}
	dfExpected.Index = []uint{0}

	dfResult.AddLine(inputLine)

	if !reflect.DeepEqual(dfExpected, dfResult) {
		t.Errorf("Dataframes different but equal expected.\nExpected: %#v\nResult: %#v", dfExpected, dfResult)
	}

}

func TestNewDataFrame(t *testing.T) {
	dfExpected := DataFrame{}
	dfExpected.NaNLayout = ""
	dfExpected.DatetimeLayout = "2006-01-02"
	dfExpected.ColumnIndex = make(map[string]int)
	dfResult := NewDataFrame(nil, "")

	if !reflect.DeepEqual(dfExpected, dfResult) {
		t.Errorf("Dataframes different but equal expected.\nExpected: %#v\nResult: %#v", dfExpected, dfResult)
	}

}

func TestString(t *testing.T) {
	input := [][]string{{"NAME", "AGE"}, {"pepito", "21"}, {"juanito", "22"}, {"pepita", "2.3"}, {"juanita", ""}}
	df := NewDataFrame(input, "")

	expected := `+---+---------+-----+
|   |  NAME   | AGE |
+---+---------+-----+
| 0 | pepito  | 21  |
| 1 | juanito | 22  |
| 2 | pepita  | 2.3 |
| 3 | juanita | NaN |
+---+---------+-----+
`
	result := fmt.Sprint(df)

	if expected != result {
		t.Errorf("Dataframe Print Different.\nExpected:\n%s\nResult:\n%s", expected, result)
	}
}

func TestHead(t *testing.T) {
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
	df := NewDataFrame(input, "")

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
	input = [][]string{
		{"name", "age"},
		{"pepito", "21"},
		{"juanito", "22"},
		{"pepita", "2.3"},
		{"juanita", ""},
	}
	dfLess := NewDataFrame(input, "")

	dfLessHead := dfLess.Head()
	if !reflect.DeepEqual(dfLess, dfLessHead) {
		t.Errorf("Dataframes different but equal expected")
	}
}

func TestTail(t *testing.T) {
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
	df := NewDataFrame(input, "")

	dfTail := df.Tail()
	dfTailExpected := df
	dfTailExpected.Values = df.Values[4:]
	dfTailExpected.Shape[0] = 10
	dfTailExpected.Index = []uint{4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	if !reflect.DeepEqual(dfTail, dfTailExpected) {
		t.Errorf("DataFrame different but equal expected.\nExpected:\n%+v\nResult:\n%+v", dfTailExpected, dfTail)
	}

	dfTail5 := df.Tail(5)
	dfTailExpected5 := df
	dfTailExpected5.Values = df.Values[9:]
	dfTailExpected5.Shape[0] = 5
	dfTailExpected5.Index = []uint{9, 10, 11, 12, 13}

	if !reflect.DeepEqual(dfTail5, dfTailExpected5) {
		t.Errorf("DataFrame different but equal expected.\nExpected:\n%+v\nResult:\n%+v", dfTailExpected5, dfTail5)
	}

	// Less than 10 rows
	input = [][]string{
		{"name", "age"},
		{"pepito", "21"},
		{"juanito", "22"},
		{"pepita", "2.3"},
		{"juanita", ""},
	}
	dfLess := NewDataFrame(input, "")

	dfLessTail := dfLess.Tail()
	if !reflect.DeepEqual(dfLess, dfLessTail) {
		t.Errorf("Dataframes different but equal expected")
	}
}

func TestSelect(t *testing.T) {
	// Base
	input := [][]string{
		{"name", "age", "data"},
		{"pepito", "21", "true"},
		{"juanito", "22", "false"},
		{"pepita", "2.3", "true"},
		{"juanita", "", "false"},
	}
	df := NewDataFrame(input, "")

	// Select two columns
	input = [][]string{
		{"name", "age"},
		{"pepito", "21"},
		{"juanito", "22"},
		{"pepita", "2.3"},
		{"juanita", ""},
	}
	dfSelect := NewDataFrame(input, "")

	dfSelected2 := df.Select("name", "age")
	dfSelected2.Shape[1] = 2

	if !reflect.DeepEqual(dfSelected2, dfSelect) {
		t.Errorf("Dataframes different but equal expected.\nExpected:\n%+v\nResult:\n%+v", dfSelect, dfSelected2)
	}

	// Select one columns
	input = [][]string{
		{"name"},
		{"pepito"},
		{"juanito"},
		{"pepita"},
		{"juanita"},
	}
	dfSelect1 := NewDataFrame(input, "")

	dfSelected1 := df.Select("name")
	dfSelected1.Shape[1] = 1

	if !reflect.DeepEqual(dfSelected1, dfSelect1) {
		t.Errorf("Dataframe different but equal expected.\nExpected:\n%#v\nResult:\n%#v", dfSelect1, dfSelected1)
	}

}

func TestFindIndex(t *testing.T) {
	//	listBase := []string{"col1", "col2", "col3"}
	listFind := []string{"col1", "col3", "col2"}
	expected := []int{0, 2, 1}
	result := []int{}
	columnIndex := make(map[string]int)
	columnIndex["col1"] = 0
	columnIndex["col3"] = 2
	columnIndex["col2"] = 1
	for _, v := range listFind {
		result = append(result, findIndex(columnIndex, v))
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Index Failed.\nExpected:\n%v\nResult:\n%v", expected, result)
	}

}

func TestDrop(t *testing.T) {
	// Base
	input := [][]string{
		{"name", "age", "data"},
		{"pepito", "21", "true"},
		{"juanito", "22", "false"},
		{"pepita", "2.3", "true"},
		{"juanita", "", "false"},
	}
	df := NewDataFrame(input, "")

	// Drop column "data"
	input = [][]string{
		{"name", "age"},
		{"pepito", "21"},
		{"juanito", "22"},
		{"pepita", "2.3"},
		{"juanita", ""},
	}
	dfDrop1 := NewDataFrame(input, "")

	df.Drop("data")

	if !reflect.DeepEqual(df, dfDrop1) {
		t.Errorf("Dataframes different but equal expected.\nExpected:\n%+v\nResult:\n%+v", dfDrop1, df)
	}

	// Drop columns "data" and "age"
	input = [][]string{
		{"name", "age", "data"},
		{"pepito", "21", "true"},
		{"juanito", "22", "false"},
		{"pepita", "2.3", "true"},
		{"juanita", "", "false"},
	}
	df = NewDataFrame(input, "")

	df.Drop("data", "age")

	input = [][]string{
		{"name"},
		{"pepito"},
		{"juanito"},
		{"pepita"},
		{"juanita"},
	}
	dfDrop2 := NewDataFrame(input, "")

	if !reflect.DeepEqual(dfDrop2, df) {
		t.Errorf("Dataframes different but equal expected.\nExpected:\n%#v\nResult:\n%#v", dfDrop2, df)
	}

}

func TestInsert(t *testing.T) {

	input := [][]string{
		{"name", "age", "data", "last_name", "year_experience", "birthdate"},
		{"pepito", "21", "true", "pepote", "2", "2020-01-02"},
		{"juanito", "22", "false", "susano", "3", "2021-02-04"},
		{"pepita", "2.3", "true", "mulano", "8", "2019-04-02"},
		{"juanita", "", "false", "pentano", "100", "2018-12-30"},
	}
	dfExpected := NewDataFrame(input, "")

	input = [][]string{
		{"name", "age"},
		{"pepito", "21"},
		{"juanito", "22"},
		{"pepita", "2.3"},
		{"juanita", ""},
	}
	dfBase := NewDataFrame(input, "")

	// Bool inserts
	dfBase.Insert("data", []Word{NewWordBool(true), NewWordBool(false), NewWordBool(true), NewWordBool(false)})

	// String inserts
	dfBase.Insert("last_name", []Word{NewWordString("pepote"), NewWordString("susano"), NewWordString("mulano"), NewWordString("pentano")})

	// Float64 inserts
	dfBase.Insert("year_experience", []Word{NewWordFloat64(float64(2)), NewWordFloat64(float64(3)), NewWordFloat64(float64(8)), NewWordFloat64(float64(100))})

	// Datetime inserts
	dfBase.Insert("birthdate", []Word{NewWordDatetime(dfBase.DatetimeLayout, "2020-01-02"), NewWordDatetime(dfBase.DatetimeLayout, "2021-02-04"), NewWordDatetime(dfBase.DatetimeLayout, "2019-04-02"), NewWordDatetime(dfBase.DatetimeLayout, "2018-12-30")})

	if !reflect.DeepEqual(dfExpected, dfBase) {
		t.Errorf("Dataframes different but equal expected.\nExpected:\n%#v\nResult:\n%#v", dfExpected, dfBase)
	}

}

func TestWhereEqual(t *testing.T) {
	input := [][]string{
		{"name", "age", "data"},
		{"pepito", "21", "true"},
		{"juanito", "22", "false"},
		{"pepita", "2.3", "true"},
		{"juanita", "", "false"},
	}
	dfBase := NewDataFrame(input, "")

	// Where equal Bool
	input = [][]string{
		{"name", "age", "data"},
		{"pepito", "21", "true"},
		{"pepita", "2.3", "true"},
	}
	dfBaseWhereDataTrue := NewDataFrame(input, "")
	dfBaseWhereDataTrue.Index = []uint{0, 2}
	dfResult := dfBase.WhereEqual("data", true)
	if !reflect.DeepEqual(dfBaseWhereDataTrue, dfResult) {
		t.Errorf("Dataframes different but equal expected.\nExpected:\n%+v\nResult:\n%+v", dfBaseWhereDataTrue, dfResult)
	}

	// Where equal String
	input = [][]string{
		{"name", "age", "data"},
		{"pepito", "21", "true"},
	}
	dfBaseWhereString := NewDataFrame(input, "")
	dfBaseWhereString.Index = []uint{0}
	dfResult = dfBase.WhereEqual("name", "pepito")
	if !reflect.DeepEqual(dfBaseWhereString, dfResult) {
		t.Errorf("Dataframes different but equal expected.\nExpected:\n%+v\nResult:\n%+v", dfBaseWhereDataTrue, dfResult)
	}

	// Where equal float64
	input = [][]string{
		{"name", "age", "data"},
		{"pepito", "21", "true"},
	}
	dfBaseFloat64 := NewDataFrame(input, "")
	dfBaseFloat64.Index = []uint{0}
	dfResult = dfBase.WhereEqual("age", float64(21))
	if !reflect.DeepEqual(dfBaseFloat64, dfResult) {
		t.Errorf("Dataframes different but equal expected.\nExpected:\n%+v\nResult:\n%+v", dfBaseFloat64, dfResult)
	}

}

func TestWhereNotEqual(t *testing.T) {
	input := [][]string{
		{"name", "age", "data"},
		{"pepito", "21", "true"},
		{"juanito", "22", "false"},
		{"pepita", "2.3", "true"},
		{"juanita", "", "false"},
	}
	dfBase := NewDataFrame(input, "")

	input = [][]string{
		{"name", "age", "data"},
		{"juanito", "22", "false"},
		{"juanita", "", "false"},
	}
	dfWhereNotEqualBool := NewDataFrame(input, "")
	dfWhereNotEqualBool.Index = []uint{1, 3}
	dfResult := dfBase.WhereNotEqual("data", true)
	if !reflect.DeepEqual(dfWhereNotEqualBool, dfResult) {
		t.Errorf("Dataframes different but equal expected.\nExpected:\n%+v\nResult:\n%+v", dfWhereNotEqualBool, dfResult)
	}

}

func TestWhereLess(t *testing.T) {
	input := [][]string{
		{"name", "age", "data"},
		{"pepito", "21", "true"},
		{"juanito", "22", "false"},
		{"pepita", "2.3", "true"},
		{"juanita", "", "false"},
	}
	dfBase := NewDataFrame(input, "")

	// Float64 comparison
	input = [][]string{
		{"name", "age", "data"},
		{"pepita", "2.3", "true"},
	}
	dfWhereLessExpected := NewDataFrame(input, "")
	dfWhereLessExpected.Index = []uint{2}
	dfResult := dfBase.WhereLess("age", float64(3))
	if !reflect.DeepEqual(dfWhereLessExpected, dfResult) {
		t.Errorf("Dataframes different but equal expected.\nExpected:\n%+v\nResult:\n%+v", dfWhereLessExpected, dfResult)
	}

	// TODO: Datetime comparison Test

}

func TestWhereGreater(t *testing.T) {
	input := [][]string{
		{"name", "age", "data"},
		{"pepito", "21", "true"},
		{"juanito", "22", "false"},
		{"pepita", "2.3", "true"},
		{"juanita", "", "false"},
	}
	dfBase := NewDataFrame(input, "")

	// Float64 comparison
	input = [][]string{
		{"name", "age", "data"},
		{"pepito", "21", "true"},
		{"juanito", "22", "false"},
	}
	dfWhereGreaterExpected := NewDataFrame(input, "")
	dfWhereGreaterExpected.Index = []uint{0, 1}

	dfResult := dfBase.WhereGreater("age", float64(3))
	if !reflect.DeepEqual(dfWhereGreaterExpected, dfResult) {
		t.Errorf("Dataframes different but equal expected.\nExpected:\n%+v\nResult:\n%+v", dfWhereGreaterExpected, dfResult)
	}

}

func TestNewWordBool(t *testing.T) {
	// Word True
	wordTrue := NewWordBool(true)

	if wordTrue.value == false {
		t.Errorf("Expected True but result: %t", wordTrue)
	}

	wordFalse := NewWordBool(false)

	if wordFalse.value == true {
		t.Errorf("Expected False but result: %t", wordFalse)
	}
}

func TestNewWordString(t *testing.T) {
	value := "Hello"
	wordString := NewWordString(value)

	if wordString.value != value {
		t.Errorf("Error on string creation. Expected: %s. But result: %s", value, wordString.value)
	}
}

func TestNewWordFloat64(t *testing.T) {
	value := float64(4)
	wordFloat64 := NewWordFloat64(value)

	if wordFloat64.value != value {
		t.Errorf("Error on Float64 creation. Expected: %v. But result: %v", value, wordFloat64)
	}
}

func TestNewDatetime(t *testing.T) {
	df := DataFrame{}
	df.DatetimeLayout = "2006-01-02"
	timeValue := "2020-01-02"
	value, _ := time.Parse(df.DatetimeLayout, timeValue)
	wordDatetime := NewWordDatetime(df.DatetimeLayout, timeValue)

	if value != wordDatetime.value {
		t.Errorf("Error on Datetime. Expected: %v. But result: %v", value, wordDatetime.value)
	}

}

func TestAdd(t *testing.T) {
	// dfExpectedString
	dfExpectedString := ColumnType{
		colName: "name",
		values:  []Word{NewWordString("pepito2"), NewWordString("juanito2"), NewWordString("pepita2"), NewWordString("juanita2")},
	}

	// dfExpectedFloat
	dfExpectedFloat := ColumnType{
		colName: "age",
		values:  []Word{NewWordFloat64(float64(23)), NewWordFloat64(float64(24)), NewWordFloat64(float64(4.3)), WordNaN{}},
	}

	// Base
	input := [][]string{
		{"name", "age", "data"},
		{"pepito", "21", "true"},
		{"juanito", "22", "false"},
		{"pepita", "2.3", "true"},
		{"juanita", "", "false"},
	}
	dfBase := NewDataFrame(input, "")

	// Add String
	dfAddString := dfBase.Column("name").Add("2")
	if !reflect.DeepEqual(dfExpectedString.values, dfAddString) {
		t.Errorf("Add function error.\nExpected:\n%v\nResult:\n%v", dfExpectedString, dfAddString)
	}

	// Add float64
	dfAddFloat := dfBase.Column("age").Add(float64(2))
	if !reflect.DeepEqual(dfExpectedFloat.values, dfAddFloat) {
		t.Errorf("Add function error.\nEpected:\n%v\nResult:\n%v", dfExpectedFloat, dfAddFloat)
	}

}
