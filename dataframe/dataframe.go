package dataframe

import (
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

// DataFrame Structure for DataFrame
type DataFrame struct {
	Columns []string
	Values  Book
	Shape   [2]int // [rowsNumber, columnsNumber]
}
