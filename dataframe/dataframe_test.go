package dataframe

import (
	"reflect"
	"testing"
)

func TestWriteWordString(t *testing.T) {

	textInput := "Hello Higor"
	textInputNaN := ""
	nanLayout := ""

	// Normal input
	wordStringResult := WriteWordString(textInput, nanLayout)
	wordStringExpected := WordString{}
	wordStringExpected.value = textInput
	wordStringExpected.isNaN = false

	if wordStringResult.value != wordStringExpected.value || wordStringResult.isNaN != wordStringExpected.isNaN {
		t.Errorf("Error wordString. Expected: %+v Result: %+v", wordStringExpected, wordStringResult)
	}

	// NaN input
	wordStringResult = WriteWordString(textInputNaN, nanLayout)
	wordStringExpected.value = textInputNaN
	wordStringExpected.isNaN = true

	if wordStringResult.value != wordStringExpected.value || wordStringResult.isNaN != wordStringExpected.isNaN {
		t.Errorf("Error wordString. Expected: %+v Result: %+v", wordStringExpected, wordStringResult)
	}

}

func TestWriteWordBool(t *testing.T) {
	textInputFalse := "false"
	textInputTrue := "true"
	textInputNaN := ""
	nanLayout := ""

	// Input true
	wordBoolResultTrue := WriteWordBool(textInputTrue, nanLayout)
	wordBoolExpectedTrue := WordBool{}
	wordBoolExpectedTrue.value = true

	if !reflect.DeepEqual(wordBoolResultTrue, wordBoolExpectedTrue) {
		t.Errorf("Both words are different but equal expected. Expected: %+v, Result: %+v", wordBoolExpectedTrue, wordBoolResultTrue)
	}

	// Input false
	wordBoolResultFalse := WriteWordBool(textInputFalse, nanLayout)
	wordBoolExpectedFalse := WordBool{}
	wordBoolExpectedFalse.value = false

	if !reflect.DeepEqual(wordBoolResultFalse, wordBoolExpectedFalse) {
		t.Errorf("Both words are different but equal expected. Expected: %+v, Result: %+v", wordBoolExpectedFalse, wordBoolResultFalse)
	}

	// Input NaN
	wordBoolResultNaN := WriteWordBool(textInputNaN, nanLayout)
	wordBoolExpectedNaN := WordBool{}
	wordBoolExpectedNaN.isNaN = true

	if !reflect.DeepEqual(wordBoolResultNaN, wordBoolExpectedNaN) {
		t.Errorf("Both words are different but equal expected. Expected: %+v, Result: %+v", wordBoolExpectedNaN, wordBoolResultNaN)
	}

}

func TestWriteLine(t *testing.T) {

	nanLayout := ""

	// Input with String
	var1 := "Higor"
	var2 := "Higor2"
	var3 := "Higor3"
	var4NaN := nanLayout
	wordStringNaN := WordString{}
	wordStringNaN.value = var4NaN
	wordStringNaN.isNaN = true

	inputLine := []string{var1, var2, var3, var4NaN}
	lineExpected := Lines{WordString{value: var1}, WordString{value: var2}, WordString{value: var3}, wordStringNaN}

	lineResult := WriteLine(inputLine, nanLayout)

	if !reflect.DeepEqual(lineExpected, lineResult) {
		t.Errorf("Both lines are different but equal expected. Expected: %v, Result: %v", lineExpected, lineResult)
	}

	// Input with String + Bool [TODO]

}

// Next steps:
/*
Replace wordString to WordString, etc. And translate
IsNaN function to know if a variable is NaN
Stringers for each PageType custom (include NaN values)
New function: schemaGenerator (to get dynamic schema) you can use maps and struct{} emtpy. Usar una goroutine que corrija en retroceso si llega haber un error en el schema seleccionado
ReadCSV: Implement all functions to read csvs and iterate with columns empties
Readme: Plasmar la analogía de los libros: Book, Page, Line, Word, etc
*/
