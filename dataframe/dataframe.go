package dataframe

import (
	"reflect"
)

type book map[string][]interface{}

// DataFrame Structure for DataFrame
type DataFrame struct {
	Columns []string
	Values  book
}

// IsEqual to kown if two DataFrame are equal
func IsEqual(dataFrame1, dataFrame2 DataFrame) bool {
	return reflect.DeepEqual(dataFrame1, dataFrame2)

}

// ReadCSV Read a CSV file and save it as a DataFrame
func ReadCSV(filename string) DataFrame {
	df := DataFrame{}
	csv := readCSV(filename)

	df.Columns = csv[0]
	chapters := book{}

	for _, rowValue := range csv[1:] {
		for columnIndex, columnValue := range rowValue {
			chapters[df.Columns[columnIndex]] = append(chapters[df.Columns[columnIndex]], columnValue)
		}
	}

	df.Values = chapters

	return df
}
