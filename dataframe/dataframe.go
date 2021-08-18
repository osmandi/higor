package dataframe

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
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
	Index          []uint
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
		return "float64", float64(valueInt)
	case errBool == nil:
		return "bool", valueBool
	case errFloat64 == nil || errInt == nil:
		return "float64", valueFloat64
	}

	return "string", nil
}

// AddLine write a line in to book
func (df *DataFrame) AddLine(inputText []string) {
	lineTranslated := WriteLine(inputText, df.NaNLayout, df.DatetimeLayout)
	df.Values = append(df.Values, lineTranslated)
	totalIndex := len(df.Index)
	if totalIndex == 0 {
		df.Index = []uint{0}
	} else {
		df.Index = append(df.Index, df.Index[len(df.Index)-1]+1)
	}

}

// NewDataFrame Create a DataFrame with default values
func NewDataFrame() DataFrame {
	df := DataFrame{NaNLayout: "", DatetimeLayout: "2006-01-02"}
	return df
}

// Stringer
func (df DataFrame) String() string {
	data := [][]string{}
	for i, v := range df.Values {
		dataInternal := []string{}
		dataInternal = append(dataInternal, fmt.Sprint(df.Index[i]))
		for _, j := range v {
			dataInternal = append(dataInternal, fmt.Sprint(j))
		}

		data = append(data, dataInternal)
	}
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	header := []string{""}
	header = append(header, df.Columns...)
	table.SetHeader(header)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.AppendBulk(data)
	table.SetBorder(true)
	table.SetCenterSeparator("+")
	table.SetAutoFormatHeaders(false)

	table.Render()

	return tableString.String()
}

func (w WordNaN) String() string {
	return "NaN"
}

func (w WordString) String() string {
	return w.value
}

func (w WordBool) String() string {
	return strconv.FormatBool(w.value)
}

func (w WordFloat64) String() string {
	return fmt.Sprintf("%g", w.value)
}

func (w WordDatetime) String() string {
	return fmt.Sprintf("%v", w.value)
}

// NewWordBool To create a WordBool
func NewWordBool(value bool) WordBool {
	return WordBool{value: value}
}

// NewWordString To create a WordString
func NewWordString(value string) WordString {
	return WordString{value: value}
}

// NewWordFloat64 To create a WordFloat64
func NewWordFloat64(value float64) WordFloat64 {
	return WordFloat64{value: value}
}

// Head Save first 10 dataframe rows
func (df DataFrame) Head(rowsLimit ...int) DataFrame {
	// Return 10 first rows
	if len(rowsLimit) == 0 {
		if len(df.Values) >= 10 {
			df.Values = df.Values[:10]
		}
		df.Shape[0] = len(df.Values)
	} else {
		if len(df.Values) >= rowsLimit[0] {
			df.Values = df.Values[:rowsLimit[0]]
		}
		df.Shape[0] = len(df.Values)
	}

	return df
}

// Tail Save the last 10 dataframe rows
func (df DataFrame) Tail(rowsLimit ...int) DataFrame {
	if len(rowsLimit) == 0 {
		if len(df.Values) >= 10 {
			df.Values = df.Values[len(df.Values)-10:]
			df.Index = df.Index[len(df.Index)-10:]
		}
		df.Shape[0] = len(df.Values)
	} else {
		if len(df.Values) >= rowsLimit[0] {
			df.Values = df.Values[len(df.Values)-rowsLimit[0]:]
			df.Index = df.Index[len(df.Index)-rowsLimit[0]:]
		}
		df.Shape[0] = len(df.Values)
	}

	return df

}

// TODO: Apply concurrency and implement errors for keys not find
// findIndex to find index
func findIndex(base, find []string) []int {
	index := []int{}
	for _, column := range find {
		for i, v := range base {
			if column == v {
				index = append(index, i)
			}
		}
	}

	return index
}

// Select to select a row
func (df DataFrame) Select(columns ...string) DataFrame {
	index := findIndex(df.Columns, columns)
	book := make(Book, len(df.Values))

	for _, v := range index {
		for j, k := range df.Values {
			book[j] = append(book[j], k[v])
		}
	}

	df.Values = book
	df.Columns = columns
	df.Shape[1] = len(df.Columns)

	return df
}

// TODO: Implement errors for columns not find
// Drop to delete a row
func (df *DataFrame) Drop(columns ...string) {
	index := findIndex(df.Columns, columns)
	for _, v := range index {
		// Remove values
		for j, k := range df.Values {
			df.Values[j] = append(k[:v], k[v+1:]...)
		}

		// Remove columns
		df.Columns = append(df.Columns[:v], df.Columns[v+1:]...)
	}

	df.Shape[1] = len(df.Columns)

}

// Insert to add a new column with its values
func (df *DataFrame) Insert(colName string, values []Word) {
	// TODO: Warning for values len more or less than df.Values
	df.Columns = append(df.Columns, colName)
	df.Shape[1] += 1
	for i := range df.Values {
		df.Values[i] = append(df.Values[i], values[i])
	}
}

// WhereEqual To find elements with == comparator
func (df DataFrame) WhereEqual(colName string, filterValue interface{}) DataFrame {
	book := Book{}
	newIndex := []uint{}
	colIndex := findIndex(df.Columns, []string{colName})[0]

	for i, v := range df.Values {
		switch v[colIndex].(type) {
		case WordBool:
			if v[colIndex].(WordBool).value == filterValue {
				book = append(book, v)
				newIndex = append(newIndex, df.Index[i])
			}
		case WordString:
			if v[colIndex].(WordString).value == filterValue {
				book = append(book, v)
				newIndex = append(newIndex, df.Index[i])
			}
		case WordFloat64:
			if v[colIndex].(WordFloat64).value == filterValue {
				book = append(book, v)
				newIndex = append(newIndex, df.Index[i])
			}

		}
	}
	df.Values = book
	df.Shape[0] = len(df.Values)
	df.Index = newIndex

	return df
}
