package dataframe

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

// Higor description
const (
	LibraryName   = "Higor"
	FirstCommit   = "2020-01-02"
	Version       = "0.4.0"
	VersionGlobal = 0
	VersionSub    = 0.4
	StableVersion = false
)

// PageString Data type for string values with support for NaN values
type PageString string

// PageBool Data type for boolean values. Not support for NaN values
type PageBool uint8

// PageFloat64 Data type for numbers and float values with support for NaN values
type PageFloat64 float64

// PageInt Data type for numbers
type PageInt int

// PageDatetime To date dates with support for NaN values
type PageDatetime time.Time

// Words Each value before to insert
type Words interface{}

// Book Interface to save a DataFrame
type Book []reflect.Value

// Schema Map to set the schema
type Schema map[string]reflect.Type

// DataFrame Structure for DataFrame
type DataFrame struct {
	Columns []string
	Values  Book
	Shape   [2]int // [rowsNumber, columnsNumber]
}

func isEqualBook(a, b interface{}) bool {
	valueA := fmt.Sprintf("%+v", a)
	valueB := fmt.Sprintf("%+v", b)
	if valueA == valueB {
		return true
	}
	return false

}

func bookGenerator(columns []string, schema Schema) reflect.Value {

	rsfs := []reflect.StructField{}

	for i, v := range columns {
		rsf := reflect.StructField{
			Name: columns[i],
			Type: schema[v],
		}

		rsfs = append(rsfs, rsf)
	}

	internalBook := reflect.StructOf(rsfs)
	return reflect.New(internalBook).Elem()
}

func parseBool(v PageBool) interface{} {

	parse := map[PageBool]interface{}{0: false, 1: true, 2: math.NaN()}

	return parse[v]
}

func writeLine(book reflect.Value, words []Words) reflect.Value {

	for i, v := range words {
		book.Field(i).Set(reflect.ValueOf(v))
	}

	return book
}

func typeString() reflect.Type {
	return reflect.TypeOf(PageString(LibraryName))
}

func typeInt() reflect.Type {
	return reflect.TypeOf(PageInt(VersionGlobal))
}

func typeFloat64() reflect.Type {
	return reflect.TypeOf(PageFloat64(VersionSub))
}

func typeBool() reflect.Type {
	return reflect.TypeOf(PageBool(uint8(VersionGlobal)))
}

func typeDatetime() reflect.Type {
	timeParse, _ := time.Parse("2006-01-02", FirstCommit)
	return reflect.TypeOf(PageDatetime(timeParse))
}

// Next steps: writeBook with CSVvalues, custom String type, custom NaN values
