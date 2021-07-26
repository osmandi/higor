package dataframe

import (
	"fmt"
	"math"
	"reflect"
	"time"
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

func typeOnStruct(a interface{}) reflect.Type {
	return reflect.TypeOf(a)
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

// Next steps: CSVCreatorMock, CSVChecker, DataFrameChecker, DataFrameStringer
