package dataframe

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

// ColumnString Column with string values
type ColumnString []string

// ColumnFloat columns with float64 values
type ColumnFloat []float64

// ColumnTime columns with time values
type ColumnTime []time.Time

// ColumnTime columns with bool values
type ColumnBool map[int]bool

// ColumnInt columns with int values
type ColumnInt []int

// Values save all columns types
type Values []interface{}

// DataFrame DatFrame struct
type DataFrame struct {
	Columns        []string
	Values         Values
	Shape          [2]int // [rowsNumber, columnsNumber]
	NaNLayout      string
	DatetimeLayout string
	Index          []uint
	ColumnIndex    map[string]int
}

// NewDataFrame Create a DataFrame with default values
func NewDataFrame(input [][]string, NaN string) DataFrame {
	if input == nil {
		return DataFrame{
			NaNLayout:      NaN,
			DatetimeLayout: "2006-01-02", // YYYY-MM-DD
			ColumnIndex:    make(map[string]int),
		}
	}
	df := DataFrame{
		NaNLayout:      NaN,
		DatetimeLayout: "2006-01-02", // YYYY-MM-DD
		ColumnIndex:    make(map[string]int),
		Columns:        input[0],
	}
	for _, v := range input[1:] {
		df.AddLine(v)
	}
	df.Shape[0] = len(df.Values)
	df.Shape[1] = len(df.Columns)
	for i, v := range df.Columns {
		df.ColumnIndex[v] = i
	}

	return df
}

// ColumnType Operations by column
type ColumnType struct {
	values  []Word
	colName string
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

// NewWordDatetime To create WordDatetime
func NewWordDatetime(format, value string) WordDatetime {
	timeParsed, err := time.Parse(format, value)
	if err != nil {
		panic(err)
	}

	return WordDatetime{value: timeParsed}
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
func findIndex(columnIndex map[string]int, columnName string) int {
	index, ok := columnIndex[columnName]

	if ok != true {
		panic(fmt.Sprintf(`Column: "%s" doesn't exists`, columnName))
	}

	return index
}

// Select to select a row
func (df DataFrame) Select(columns ...string) DataFrame {
	indexs := []int{}
	for _, v := range columns {
		indexs = append(indexs, findIndex(df.ColumnIndex, v))
	}

	book := make(Book, len(df.Values))

	for _, v := range indexs {
		for j, k := range df.Values {
			book[j] = append(book[j], k[v])
		}
	}

	df.Values = book
	df.Columns = columns
	df.Shape[1] = len(df.Columns)

	df.ColumnIndex = make(map[string]int)

	for i, v := range df.Columns {
		df.ColumnIndex[v] = i
	}

	return df
}

// Column To select DataFrame with one column
func (df DataFrame) Column(columnName string) ColumnType {
	index := df.ColumnIndex[columnName]
	columnType := ColumnType{}
	for _, v := range df.Values {
		columnType.values = append(columnType.values, v[index])
	}
	columnType.colName = columnName

	return columnType
}

// TODO: Implement errors for columns not find
// Drop to delete a row
func (df *DataFrame) Drop(columns ...string) {
	indexs := []int{}
	for _, v := range columns {
		indexs = append(indexs, findIndex(df.ColumnIndex, v))
	}

	for _, v := range indexs {
		// Remove values
		for j, k := range df.Values {
			df.Values[j] = append(k[:v], k[v+1:]...)
		}

		// Remove columns
		df.Columns = append(df.Columns[:v], df.Columns[v+1:]...)
	}

	df.Shape[1] = len(df.Columns)

	for i, v := range df.Columns {
		df.ColumnIndex[v] = i
	}

	df.ColumnIndex = make(map[string]int)
	for i, v := range df.Columns {
		df.ColumnIndex[v] = i
	}
}

// Insert to add a new column with its values
func (df *DataFrame) Insert(colName string, values []Word) {
	// TODO: Warning for values len more or less than df.Values
	df.Columns = append(df.Columns, colName)
	df.Shape[1] += 1
	for i := range df.Values {
		df.Values[i] = append(df.Values[i], values[i])
	}
	df.ColumnIndex[colName] = len(df.ColumnIndex)
}

// WhereEqual To find elements with == comparator
func (df DataFrame) WhereEqual(colName string, filterValue interface{}) DataFrame {
	book := Book{}
	newIndex := []uint{}
	colIndex := findIndex(df.ColumnIndex, colName)

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
		case WordDatetime:
			if v[colIndex].(WordDatetime).value == filterValue {
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

// WhereNotEqual To find elements with != comparator
func (df DataFrame) WhereNotEqual(colName string, filterValue interface{}) DataFrame {
	book := Book{}
	newIndex := []uint{}
	colIndex := findIndex(df.ColumnIndex, colName)

	for i, v := range df.Values {
		switch v[colIndex].(type) {
		case WordBool:
			if v[colIndex].(WordBool).value != filterValue {
				book = append(book, v)
				newIndex = append(newIndex, df.Index[i])
			}
		case WordString:
			if v[colIndex].(WordString).value != filterValue {
				book = append(book, v)
				newIndex = append(newIndex, df.Index[i])
			}
		case WordFloat64:
			if v[colIndex].(WordFloat64).value != filterValue {
				book = append(book, v)
				newIndex = append(newIndex, df.Index[i])
			}
		case WordDatetime:
			if v[colIndex].(WordDatetime).value != filterValue {
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

// WhereLess To find elements with <
func (df DataFrame) WhereLess(colName string, filterValue interface{}) DataFrame {
	book := Book{}
	newIndex := []uint{}
	colIndex := findIndex(df.ColumnIndex, colName)

	for i, v := range df.Values {
		switch v[colIndex].(type) {
		case WordFloat64:
			if v[colIndex].(WordFloat64).value < filterValue.(float64) {
				book = append(book, v)
				newIndex = append(newIndex, df.Index[i])
			}
		case WordDatetime:
			if v[colIndex].(WordDatetime).value.Before(filterValue.(time.Time)) {
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

// WhereGreater To find elements with >
func (df DataFrame) WhereGreater(colName string, filterValue interface{}) DataFrame {
	book := Book{}
	newIndex := []uint{}
	colIndex := findIndex(df.ColumnIndex, colName)

	for i, v := range df.Values {
		switch v[colIndex].(type) {
		case WordFloat64:
			if v[colIndex].(WordFloat64).value > filterValue.(float64) {
				book = append(book, v)
				newIndex = append(newIndex, df.Index[i])
			}
		case WordDatetime:
			if v[colIndex].(WordDatetime).value.After(filterValue.(time.Time)) {
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

// WhereLessOrEqual To find elements with <
func (df DataFrame) WhereOrEqual(colName string, filterValue interface{}) DataFrame {
	book := Book{}
	newIndex := []uint{}
	colIndex := findIndex(df.ColumnIndex, colName)

	for i, v := range df.Values {
		switch v[colIndex].(type) {
		case WordFloat64:
			if v[colIndex].(WordFloat64).value <= filterValue.(float64) {
				book = append(book, v)
				newIndex = append(newIndex, df.Index[i])
			}
		case WordDatetime:
			if v[colIndex].(WordDatetime).value.Before(filterValue.(time.Time)) || v[colIndex].(WordDatetime).value == filterValue.(time.Time) {
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

// WhereGreaterOrEqual To find elements with >
func (df DataFrame) WhereGreaterOrEqual(colName string, filterValue interface{}) DataFrame {
	book := Book{}
	newIndex := []uint{}
	colIndex := findIndex(df.ColumnIndex, colName)

	for i, v := range df.Values {
		switch v[colIndex].(type) {
		case WordFloat64:
			if v[colIndex].(WordFloat64).value >= filterValue.(float64) {
				book = append(book, v)
				newIndex = append(newIndex, df.Index[i])
			}
		case WordDatetime:
			if v[colIndex].(WordDatetime).value.After(filterValue.(time.Time)) || v[colIndex].(WordDatetime).value == filterValue.(time.Time) {
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

// Add to add elements
func (ct ColumnType) Add(valueInput interface{}) []Word {
	for i, v := range ct.values {
		switch v.(type) {
		case WordString:
			ct.values[i] = NewWordString(v.(WordString).value + valueInput.(string))
		case WordFloat64:
			ct.values[i] = NewWordFloat64(v.(WordFloat64).value + valueInput.(float64))
		}
	}

	return ct.values
}
