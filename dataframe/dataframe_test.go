package dataframe

import (
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
	lineExpected2 := Lines{WordString{value: "Higor"}, WordInt{value: int(1)}, WordFloat64{value: float64(2.2)}, WordBool{value: false}, WordNaN{}, WordDatetime{value: datetime2}}
	lineResult2 := WriteLine(inputLine2, nanLayout, layoutDatetime)

	if !reflect.DeepEqual(lineExpected2, lineResult2) {
		t.Errorf("Both lines are different but equal expected.\nExpected: %v\nResult: %v", lineExpected2, lineResult2)
	}

}

func TestTranslateWord(t *testing.T) {
	nanLayout := ""
	layoutDatetime := "2006-01-02"
	variables := []string{nanLayout, "Higor", "1", "1.2", "true", "True", "False", "false", "2020-02-01"}
	result := []string{"NaN", "string", "int", "float64", "bool", "bool", "bool", "bool", "datetime"}

	for i, v := range variables {
		trans, _ := translateWord(v, nanLayout, layoutDatetime)

		if trans != result[i] {
			t.Errorf("Both different but equal result. Expected: %s, Result: %s", result[i], trans)
		}
	}

}

func TestAddLine(t *testing.T) {
	nanLayout := ""
	layoutDatetime := "2006-01-02"

	dfExpected := DataFrame{}
	dfResult := DataFrame{}

	dfExpected.NaNLayout = nanLayout
	dfExpected.DatetimeLayout = layoutDatetime
	dfResult.NaNLayout = nanLayout
	dfResult.DatetimeLayout = layoutDatetime

	datetime, _ := time.Parse(layoutDatetime, "2020-01-01")
	inputLine := []string{"Higor", "1", "2.2", "false", "", "2020-01-01"}
	lineExpected := Lines{WordString{value: "Higor"}, WordInt{value: int(1)}, WordFloat64{value: float64(2.2)}, WordBool{value: false}, WordNaN{}, WordDatetime{value: datetime}}

	dfExpected.Values = Book{lineExpected}

	dfResult.AddLine(inputLine)

	if !reflect.DeepEqual(dfExpected, dfResult) {
		t.Errorf("Dataframes different but equal expected.\nExpected: %+v\nResult: %+v", dfExpected, dfResult)
	}

}

// TODO: DataFrame.String()
// TODO: higor.ReadCSV()

// Next steps:
/*
Replace wordString to WordString, etc. And translate
IsNaN function to know if a variable is NaN
Stringers for each PageType custom (include NaN values)
New function: schemaGenerator (to get dynamic schema) you can use maps and struct{} emtpy. Usar una goroutine que corrija en retroceso si llega haber un error en el schema seleccionado
ReadCSV: Implement all functions to read csvs and iterate with columns empties
Readme: Plasmar la analogía de los libros: Book, Page, Line, Word, etc
*/
