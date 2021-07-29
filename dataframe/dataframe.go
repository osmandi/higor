package dataframe

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
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

// wordString Data type for string values with support for NaN values
type wordString string

// wordBool Data type for boolean values. Not support for NaN values
type wordBool uint8

// wordFloat64 Data type for numbers and float values with support for NaN values
type wordFloat64 float64

// wordInt Data type for numbers
type wordInt int

// wordDatetime To date dates with support for NaN values
type wordDatetime time.Time

// Words It's a value
type Word reflect.Value

// Lines It's a row
type Lines []Word

// Book save multiple lines
type Book []Lines

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

func parseBool(v wordBool) interface{} {

	parse := map[wordBool]interface{}{0: false, 1: true, 2: math.NaN()}

	return parse[v]
}

func typeString() reflect.Type {
	return reflect.TypeOf(wordString(LibraryName))
}

func typeInt() reflect.Type {
	return reflect.TypeOf(wordInt(VersionGlobal))
}

func typeFloat64() reflect.Type {
	return reflect.TypeOf(wordFloat64(VersionSub))
}

func typeBool() reflect.Type {
	return reflect.TypeOf(wordBool(uint8(VersionGlobal)))
}

func typeDatetime() reflect.Type {
	timeParse, _ := time.Parse("2006-01-02", FirstCommit)
	return reflect.TypeOf(wordDatetime(timeParse))
}

func translateWord(text string, typeValue reflect.Type) (Word, error) {

	nanValueInput := ""
	datetimeLayout := "2006-01-02"

	switch typeValue {
	case typeString():
		return Word(reflect.ValueOf(wordString(text))), nil
	case typeInt():
		if text == nanValueInput {
			return Word{}, fmt.Errorf("Error parsing Int to NaN")
		}
		result, err := strconv.Atoi(text)
		return Word(reflect.ValueOf(wordInt(result))), err
	case typeFloat64():
		if text == nanValueInput {
			return Word(reflect.ValueOf(wordFloat64(math.NaN()))), nil
		}
		result, err := strconv.ParseFloat(text, 64)
		return Word(reflect.ValueOf(wordFloat64(result))), err
	case typeBool():
		if text == nanValueInput {
			return Word(reflect.ValueOf(wordBool(2))), nil
		}
		result, err := strconv.ParseBool(text)
		if result {
			return Word(reflect.ValueOf(wordBool(uint8(1)))), err
		} else {
			return Word(reflect.ValueOf(wordBool(uint8(0)))), err
		}
	case typeDatetime():
		if text == nanValueInput {
			valueDatetimeNaN := time.Date(0001, 1, 1, 0, 0, 0, 0, time.UTC)
			return Word(reflect.ValueOf(wordDatetime(valueDatetimeNaN))), nil
		}
		result, err := time.Parse(datetimeLayout, text)
		return Word(reflect.ValueOf(wordDatetime(result))), err
	}
	return Word{}, fmt.Errorf("Error to translate the word: %s", text)
}
