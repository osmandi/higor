package dataframe

import "testing"

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

// Next steps:
/*
Replace wordString to WordString, etc. And translate
IsNaN function to know if a variable is NaN
Stringers for each PageType custom (include NaN values)
New function: schemaGenerator (to get dynamic schema) you can use maps and struct{} emtpy. Usar una goroutine que corrija en retroceso si llega haber un error en el schema seleccionado
ReadCSV: Implement all functions to read csvs and iterate with columns empties
Readme: Plasmar la analogía de los libros: Book, Page, Line, Word, etc
*/
