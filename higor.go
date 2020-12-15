package higor

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"gonum.org/v1/gonum/stat"
)

type bookmine []interface{}

type book map[string]bookmine

func (b bookmine) Mean() float64 {
	var valuesFloat []float64
	for _, v := range b {
		valuesFloat = append(valuesFloat, v.(float64))
	}
	// Calculate the mean
	mean := stat.Mean(valuesFloat, nil)
	return mean
}

// DataFrame contain the dataFrames methods and atributes
type DataFrame struct {
	Columns  []string
	Values   book
	Shape    [2]int
	Sep      rune
	Filename string
	Index    []int
	SafeMode bool // To know if should skip with the first error
}

// PrintHelloHigor To get greets from higor library
func PrintHelloHigor() string {
	return "Hello from higor"
}

func stringToType(a string) interface{} {
	// Remove whitespace
	aStrim := strings.TrimSpace(a)

	// try if it's empty
	if aStrim == "" {
		return math.NaN()
	}

	// Try intent convert to Int type
	v, err := strconv.Atoi(aStrim)
	if err != nil {

		// If there is an error, try to convert to float64
		v, err := strconv.ParseFloat(aStrim, 64)

		// If there is an error, try bool
		if err != nil {
			v, err := strconv.ParseBool(aStrim)

			// If there is an error, return as string
			if err != nil {
				return aStrim
			}

			// Return if it's bool
			return v
		}

		// Return if it's float
		return v
	}

	// Return it it's int
	return v
}

/*
func printDataFrame(columns []string, values [][]string) {
	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent)

	// Print Header
	header := strings.Join(columns, "\t")
	fmt.Fprintf(w, "%s\t%v\n", " ", header)

	// Print values
	var value string
	for _, v := range values {
		value = strings.Join(v[:], "\t")
		fmt.Fprintf(w, "%v\n", value)
	}

	w.Flush()
}

// Head to get the first 5 rows
func (df DataFrame) Head() {
	printDataFrame(df.Columns, df.Values[:5])
}

// Tail to get the last 5 rows
func (df DataFrame) Tail() {
	tail := df.Values[len(df.Values)-5 : len(df.Values)]
	printDataFrame(df.Columns, tail)
}
*/

// Drop columns
func (df *DataFrame) Drop(columns ...string) {

	for _, column := range columns {
		// Check if that column exists
		_, ok := df.Values[column]

		// If exists, delete the column
		if ok == true {
			for i, col := range df.Columns {
				if column == col {
					// Delete from the values
					delete(df.Values, column)

					// Exclude the column
					newColumns := append(df.Columns[:i], df.Columns[i+1:]...)
					df.Columns = newColumns
				}
			}

		} else {
			messageError := fmt.Sprintf("The column '%s' don't exists on the DataFrame\n", column)
			if df.SafeMode {
				// If SafeMode is active
				log.Fatal(messageError)
			}
			fmt.Println(messageError)
			continue
		}
	}

}

// String Return string to print it
func (df DataFrame) String() string {
	printDataFrame(df.Columns, df.Values, df.Index)

	return ""
}

func printDataFrame(columns []string, values book, index []int) {

	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)

	// Print Header
	header := strings.Join(columns, "\t")
	fmt.Fprintf(w, "index\t%v\n", header)

	for i := range index {
		var line []string
		for _, col := range columns {
			stringValue := fmt.Sprintf("%v", values[col][i])
			line = append(line, stringValue)
		}
		// Print on the table
		value := strings.Join(line[:], "\t")
		fmt.Fprintf(w, "%d\t%v\n", i, value)
	}

	defer w.Flush()
}

// NewDataFrame set default variables to read dataframe
func NewDataFrame(filename string) *DataFrame {
	return &DataFrame{
		Sep:      ',',
		Filename: filename,
		SafeMode: false,
	}
}

// ReadCSV to read CSV files
func (df *DataFrame) ReadCSV() {

	// Open the iris dataset file
	f, err := os.Open(df.Filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	// Set custom parameters
	reader.Comma = df.Sep
	var values [][]string
	var indexList []int
	index := 0
	indexList = append(indexList, index)
	for {
		var lines []string
		// Read in a row. Check if we are at the end of the line
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		for _, value := range record {
			lines = append(lines, value)

		}

		values = append(values, lines)

		index++
		indexList = append(indexList, index)
	}

	// Set dataframe columns
	df.Columns = values[0]
	df.Index = indexList[:len(indexList)-2]

	// Set values
	valuesPerColumn := make(book)
	for i, v := range df.Columns {
		for _, r := range values[1:] {
			valuesPerColumn[v] = append(valuesPerColumn[v], stringToType(r[i]))
		}
	}
	df.Values = valuesPerColumn
	df.Shape = [2]int{len(df.Columns), len(df.Values)}

}
