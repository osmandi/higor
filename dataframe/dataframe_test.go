package dataframe

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"time"
)

type book1 struct {
	name string
}

type book2 struct {
	name string
}

type book3 struct {
	noName string
}

type bookComplete struct {
	ColString   wordString
	ColInt      wordInt
	ColFloat64  wordFloat64
	ColBool     wordBool
	ColDatetime wordDatetime
}

func TestIsEqualBookEqual(t *testing.T) {
	internalName := "Higor"

	book1Example := book1{
		name: internalName,
	}

	book2Example := book2{
		name: internalName,
	}

	book3Example := book3{
		noName: internalName,
	}

	bookComparation := isEqualBook(book1Example, book2Example)

	if !bookComparation {
		t.Errorf("Error, both DataFrame are different but equal expected. %+v vs %+v", book1Example, book2Example)
	}

	bookComparation = isEqualBook(book1Example, book3Example)

	if bookComparation {
		t.Errorf("Error, both DataFrame are equal but different expected. %+v vs %+v", book1Example, book3Example)
	}
}

func TestTypeString(t *testing.T) {
	result := typeString()
	expected := reflect.TypeOf(wordString(LibraryName))

	if result != expected {
		t.Errorf("Type are different. Expected: %v, but result: %v", expected, result)
	}
}

func TestTypeInt(t *testing.T) {
	result := typeInt()
	expected := reflect.TypeOf(wordInt(VersionGlobal))

	if result != expected {
		t.Errorf("Type are different. Expected: %v, but result: %v", expected, result)
	}

}

func TestTypeFloat64(t *testing.T) {
	result := typeFloat64()
	expected := reflect.TypeOf(wordFloat64(VersionSub))

	if result != expected {
		t.Errorf("Type are different. Expected: %v, but result: %v", expected, result)
	}
}

func TestTypeBool(t *testing.T) {
	result := typeBool()
	expected := reflect.TypeOf(wordBool(uint8(VersionGlobal)))

	if result != expected {
		t.Errorf("Type are different. Expected: %v, but result: %v", expected, result)
	}
}

func TestTypeDatetime(t *testing.T) {
	result := typeDatetime()
	timeParse, _ := time.Parse("2006-01-02", FirstCommit)
	expected := reflect.TypeOf(wordDatetime(timeParse))

	if result != expected {
		t.Errorf("Type are different. Expected: %v, but result: %v", expected, result)
	}

}

func TestParseBool(t *testing.T) {

	// Equal
	slicesIterator := []wordBool{0, 1, 2}
	resultExpected := []interface{}{false, true, math.NaN()}

	for i, v := range slicesIterator {
		result := parseBool(v)
		if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", resultExpected[i]) {
			t.Errorf("Result: %v but expected: %v", result, resultExpected[i])
		}
	}

	// Different
	slicesIterator = []wordBool{1, 2, 0}
	resultExpected = []interface{}{false, true, math.NaN()}

	for i, v := range slicesIterator {
		result := parseBool(v)
		if fmt.Sprintf("%v", result) == fmt.Sprintf("%v", resultExpected[i]) {
			t.Errorf("Result: %v but expected: %v", result, resultExpected[i])
		}
	}

}

/*
func TestWriteLine(t *testing.T) {

	timeParse, _ := time.Parse("2006-01-02", "2020-01-02")
	linesExpected := Lines{Word(reflect.ValueOf(wordString("Higor"))), Word(reflect.ValueOf(wordInt(1))), Word(reflect.ValueOf(wordFloat64(1.1))), Word(reflect.ValueOf(wordBool(0))), Word(reflect.ValueOf(wordDatetime(timeParse)))}
	linesResult := Lines{}

	// Columns and Schema
	columns := []string{"ColString", "ColInt", "ColFloat64", "ColBool", "ColDatetime"}
	schema := Schema{
		columns[0]: typeString(),
		columns[1]: typeInt(),
		columns[2]: typeFloat64(),
		columns[3]: typeBool(),
		columns[4]: typeDatetime(),
	}

	// Book generate
	book := bookGenerator(columns, schema)
	bookExpected := bookGenerator(columns, schema)

	for i := range columns {
		bookExpected.Field(i).Set(reflect.ValueOf(lines[i]))
	}

	bookResult := writeLine(book, words)

	// Equal
	bookComparation := isEqualBook(bookExpected, bookResult)

	if !bookComparation {
		t.Errorf("Error, both book are different but equal expected. \n%+v vs \n%+v", bookExpected, bookResult)
	}

	// Different
	bookExpected.FieldByName(columns[0]).Set(reflect.ValueOf(wordString("Hello Higor")))
	bookComparation = isEqualBook(bookExpected, bookResult)

	if bookComparation {
		t.Errorf("Error, both book are equal but different expected. \n%+v vs \n%+v", bookExpected, bookResult)
	}
}
*/

/*
func TestTranslateWords(t *testing.T) {

	datetimeLayout := "2006-01-02"
	timeParse, _ := time.Parse(datetimeLayout, "2020-01-02")
	valueDatetimeNaN := time.Date(0001, 1, 1, 0, 0, 0, 0, time.UTC)

	textInput := []string{"Higor", "0", "0.4", "false", "2020-01-02"}
	textExpected := []Word{wordString("Higor"), wordInt(0), wordFloat64(0.4), wordBool(0), wordDatetime(timeParse)}
	textInputNaN := []string{"", "1", "", "", ""}
	textExpectedNaN := []Word{wordString(""), wordInt(1), wordFloat64(math.NaN()), wordBool(2), wordDatetime(valueDatetimeNaN)}

	columns := []string{"ColString", "ColInt", "ColFloat64", "ColBool", "ColDatetime"}
	schema := Schema{
		columns[0]: typeString(),
		columns[1]: typeInt(),
		columns[2]: typeFloat64(),
		columns[3]: typeBool(),
		columns[4]: typeDatetime(),
	}

	// With normal values
	for i, v := range textInput {
		result, err := translateWord(v, schema[columns[i]])
		if err != nil {
			panic(err)
		}
		if result != textExpected[i] {
			t.Errorf("Different values but equal expected. Expected: %v, Result: %v", textExpected[i], result)
		}
	}

	// With NaN values
	for i, v := range textInputNaN {
		result, err := translateWord(v, schema[columns[i]])
		if err != nil {
			panic(err)
		}
		if fmt.Sprint(result) != fmt.Sprint(textExpectedNaN[i]) {
			t.Errorf("Different values but equal expected. Expected: %v, Result: %v", textExpectedNaN[i], result)
		}
	}

}
*/

// Next steps:
/*
Replace wordString to WordString, etc. And translate
IsNaN function to know if a variable is NaN
Stringers for each PageType custom (include NaN values)
New function: schemaGenerator (to get dynamic schema) you can use maps and struct{} emtpy. Usar una goroutine que corrija en retroceso si llega haber un error en el schema seleccionado
ReadCSV: Implement all functions to read csvs and iterate with columns empties
Readme: Plasmar la analogía de los libros: Book, Page, Line, Word, etc
*/
