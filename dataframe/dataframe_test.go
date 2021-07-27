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
	ColString   PageString
	ColInt      PageInt
	ColFloat64  PageFloat64
	ColBool     PageBool
	ColDatetime PageDatetime
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
	expected := reflect.TypeOf(PageString(LibraryName))

	if result != expected {
		t.Errorf("Type are different. Expected: %v, but result: %v", expected, result)
	}
}

func TestTypeInt(t *testing.T) {
	result := typeInt()
	expected := reflect.TypeOf(PageInt(VersionGlobal))

	if result != expected {
		t.Errorf("Type are different. Expected: %v, but result: %v", expected, result)
	}

}

func TestTypeFloat64(t *testing.T) {
	result := typeFloat64()
	expected := reflect.TypeOf(PageFloat64(VersionSub))

	if result != expected {
		t.Errorf("Type are different. Expected: %v, but result: %v", expected, result)
	}
}

func TestTypeBool(t *testing.T) {
	result := typeBool()
	expected := reflect.TypeOf(PageBool(uint8(VersionGlobal)))

	if result != expected {
		t.Errorf("Type are different. Expected: %v, but result: %v", expected, result)
	}
}

func TestTypeDatetime(t *testing.T) {
	result := typeDatetime()
	timeParse, _ := time.Parse("2006-01-02", FirstCommit)
	expected := reflect.TypeOf(PageDatetime(timeParse))

	if result != expected {
		t.Errorf("Type are different. Expected: %v, but result: %v", expected, result)
	}

}

func TestParseBool(t *testing.T) {

	// Equal
	slicesIterator := []PageBool{0, 1, 2}
	resultExpected := []interface{}{false, true, math.NaN()}

	for i, v := range slicesIterator {
		result := parseBool(v)
		if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", resultExpected[i]) {
			t.Errorf("Result: %v but expected: %v", result, resultExpected[i])
		}
	}

	// Different
	slicesIterator = []PageBool{1, 2, 0}
	resultExpected = []interface{}{false, true, math.NaN()}

	for i, v := range slicesIterator {
		result := parseBool(v)
		if fmt.Sprintf("%v", result) == fmt.Sprintf("%v", resultExpected[i]) {
			t.Errorf("Result: %v but expected: %v", result, resultExpected[i])
		}
	}

}

func TestBookGenerator(t *testing.T) {

	// Setting values
	var valueString PageString = "Higor"
	var valueInt PageInt = 1
	var valueFloat64 PageFloat64 = 1.1
	var valueBool PageBool = 0
	timeParse, _ := time.Parse("2006-01-02", "2020-01-02")
	valueDatetime := PageDatetime(timeParse)

	// Columns and Schema
	columns := []string{"ColString", "ColInt", "ColFloat64", "ColBool", "ColDatetime"}
	schema := Schema{
		columns[0]: typeString(),
		columns[1]: typeInt(),
		columns[2]: typeFloat64(),
		columns[3]: typeBool(),
		columns[4]: typeDatetime(),
	}

	bookExpected := bookComplete{
		ColString:   valueString,
		ColInt:      valueInt,
		ColFloat64:  valueFloat64,
		ColBool:     valueBool,
		ColDatetime: valueDatetime,
	}

	bookResult := bookGenerator(columns, schema)
	bookResult.FieldByName(columns[0]).Set(reflect.ValueOf(valueString))
	bookResult.FieldByName(columns[1]).Set(reflect.ValueOf(valueInt))
	bookResult.FieldByName(columns[2]).Set(reflect.ValueOf(valueFloat64))
	bookResult.FieldByName(columns[3]).Set(reflect.ValueOf(valueBool))
	bookResult.FieldByName(columns[4]).Set(reflect.ValueOf(valueDatetime))

	bookComparation := isEqualBook(bookExpected, bookResult)

	if !bookComparation {
		t.Errorf("Error, both book are different but equal expected. \n%+v vs \n%+v", bookExpected, bookResult)
	}
}

func TestWriteLine(t *testing.T) {

	// Setting values
	var valueString PageString = "Higor"
	var valueInt PageInt = 1
	var valueFloat64 PageFloat64 = 1.1
	var valueBool PageBool = 0
	timeParse, _ := time.Parse("2006-01-02", "2020-01-02")
	valueDatetime := PageDatetime(timeParse)

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

	// Expected book values
	bookExpected.FieldByName(columns[0]).Set(reflect.ValueOf(valueString))
	bookExpected.FieldByName(columns[1]).Set(reflect.ValueOf(valueInt))
	bookExpected.FieldByName(columns[2]).Set(reflect.ValueOf(valueFloat64))
	bookExpected.FieldByName(columns[3]).Set(reflect.ValueOf(valueBool))
	bookExpected.FieldByName(columns[4]).Set(reflect.ValueOf(valueDatetime))

	values := Words{valueString, valueInt, valueFloat64, valueBool, valueDatetime}

	bookResult := writeLine(book, values)

	// Equal
	bookComparation := isEqualBook(bookExpected, bookResult)

	if !bookComparation {
		t.Errorf("Error, both book are different but equal expected. \n%+v vs \n%+v", bookExpected, bookResult)
	}

	// Different
	bookExpected.FieldByName(columns[0]).Set(reflect.ValueOf(PageString("Hello Higor")))
	bookComparation = isEqualBook(bookExpected, bookResult)

	if bookComparation {
		t.Errorf("Error, both book are equal but different expected. \n%+v vs \n%+v", bookExpected, bookResult)
	}

}
