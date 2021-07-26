package dataframe

import (
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

func TestTypeOnStruct(t *testing.T) {
	var typeInt PageInt = 1
	var typeFloat64 PageFloat64 = 1.1
	var typeString PageString = "Higor"
	var typeBool PageBool = 0
	timeParse, _ := time.Parse("2006-01-02", "2020-01-02")
	typeDatetime := PageDatetime(timeParse)

	sliceTypes := [5]interface{}{typeInt, typeFloat64, typeString, typeDatetime, typeBool}

	for _, v := range sliceTypes {
		getTypeResult := typeOnStruct(v)
		getTypeExpected := reflect.TypeOf(v)
		if getTypeResult != getTypeExpected {
			t.Errorf("Type are different. Expected: %v, but result: %v", getTypeExpected, getTypeResult)
		}
	}

}

func TestBookGenerator(t *testing.T) {

	// Setting values
	var typeInt PageInt = 1
	var typeFloat64 PageFloat64 = 1.1
	var typeBool PageBool = 0
	var typeString PageString = "Higor"
	timeParse, _ := time.Parse("2006-01-02", "2020-01-02")
	typeDatetime := PageDatetime(timeParse)

	// Columns and Schema
	columns := []string{"ColString", "ColInt", "ColFloat64", "ColBool", "ColDatetime"}
	schema := Schema{
		columns[0]: typeOnStruct(typeString),
		columns[1]: typeOnStruct(typeInt),
		columns[2]: typeOnStruct(typeFloat64),
		columns[3]: typeOnStruct(typeBool),
		columns[4]: typeOnStruct(typeDatetime),
	}

	bookExpected := bookComplete{
		ColString:   typeString,
		ColInt:      typeInt,
		ColFloat64:  typeFloat64,
		ColBool:     typeBool,
		ColDatetime: typeDatetime,
	}

	bookResult := bookGenerator(columns, schema)
	bookResult.FieldByName(columns[0]).Set(reflect.ValueOf(typeString))
	bookResult.FieldByName(columns[1]).Set(reflect.ValueOf(typeInt))
	bookResult.FieldByName(columns[2]).Set(reflect.ValueOf(typeFloat64))
	bookResult.FieldByName(columns[3]).Set(reflect.ValueOf(typeBool))
	bookResult.FieldByName(columns[4]).Set(reflect.ValueOf(typeDatetime))

	bookComparation := isEqualBook(bookExpected, bookResult)

	if !bookComparation {
		t.Errorf("Error, both book are different but equal expected. \n%+v vs \n%+v", bookExpected, bookResult)
	}
}
