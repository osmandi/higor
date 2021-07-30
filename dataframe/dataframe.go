package dataframe

import (
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

type Word interface {
}

// WordNaN to save NaN values
type WordNaN struct {
	Word
}

// wordString Data type for string values with support for NaN values
type WordString struct {
	Word
	value string
}

// wordBool Data type for boolean values. Not support for NaN values
type WordBool struct {
	Word
	value bool
}

// wordFloat64 Data type for numbers and float values with support for NaN values
type WordFloat64 struct {
	Word
	value float64
}

// wordInt Data type for numbers
type WordInt struct {
	Word
	value int
}

// wordDatetime To date dates with support for NaN values
type WordDatetime struct {
	Word
	value time.Time
}

// Lines It's a row
type Lines []Word

// Book save multiple lines
type Book []Lines

// DataFrame Structure for DataFrame
type DataFrame struct {
	Columns []string
	Values  Book
	Shape   [2]int // [rowsNumber, columnsNumber]
}

func WriteWordString(text string) WordString {
	wordString := WordString{value: text}
	return wordString

}

func WriteWordBool(text string) WordBool {

	parseBool, err := strconv.ParseBool(text)

	if err != nil {
		panic(err)
	}

	wordBool := WordBool{value: parseBool}

	return wordBool

}

func WriteLine(textInput []string, nanLayout, layoutDatetime string) Lines {
	line := Lines{}
	for _, v := range textInput {
		switch trans := translateWord(v, nanLayout, layoutDatetime); trans {
		case "NaN":
			line = append(line, WordNaN{})
		}
	}

	return line
}

func translateWord(textInput, nanLayout, layoutDatetime string) string {
	_, errDatetime := time.Parse(layoutDatetime, textInput)
	_, errInt := strconv.Atoi(textInput)
	_, errBool := strconv.ParseBool(textInput)
	_, errFloat64 := strconv.ParseFloat(textInput, 64)
	switch {
	case textInput == nanLayout:
		return "NaN"
	case errDatetime == nil:
		return "datetime"
	case errInt == nil:
		return "int"
	case errBool == nil:
		return "bool"
	case errFloat64 == nil:
		return "float64"
	}

	return "string"
}
