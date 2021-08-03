package dataframe

import (
	"strconv"
	"time"
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
	Columns        []string
	Values         Book
	Shape          [2]int // [rowsNumber, columnsNumber]
	NaNLayout      string
	DatetimeLayout string
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
		switch trans, value := translateWord(v, nanLayout, layoutDatetime); trans {
		case "NaN":
			line = append(line, WordNaN{})
		case "datetime":
			line = append(line, WordDatetime{value: value.(time.Time)})
		case "int":
			line = append(line, WordInt{value: value.(int)})
		case "bool":
			line = append(line, WordBool{value: value.(bool)})
		case "float64":
			line = append(line, WordFloat64{value: value.(float64)})
		default:
			line = append(line, WordString{value: v})
		}
	}
	return line
}

func translateWord(textInput, nanLayout, layoutDatetime string) (valueType string, value interface{}) {
	valueDatetime, errDatetime := time.Parse(layoutDatetime, textInput)
	valueInt, errInt := strconv.Atoi(textInput)
	valueBool, errBool := strconv.ParseBool(textInput)
	valueFloat64, errFloat64 := strconv.ParseFloat(textInput, 64)
	switch {
	case textInput == nanLayout:
		return "NaN", nil
	case errDatetime == nil:
		return "datetime", valueDatetime
	case errInt == nil:
		return "int", valueInt
	case errBool == nil:
		return "bool", valueBool
	case errFloat64 == nil:
		return "float64", valueFloat64
	}

	return "string", nil
}

// AddLine write a line in to book
func (df *DataFrame) AddLine(inputText []string) {
	lineTranslated := WriteLine(inputText, df.NaNLayout, df.DatetimeLayout)
	df.Values = append(df.Values, lineTranslated)
}

// NewDataFrame Create a DataFrame with default values
func NewDataFrame() DataFrame {
	df := DataFrame{NaNLayout: "", DatetimeLayout: "2006-01-02"}
	return df
}
