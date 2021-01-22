package higor

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"

	"gonum.org/v1/gonum/stat"
)

// Page It's one column
type Page []interface{}

// Book It's to have all columns with its values
type Book map[string]Page

// DataFrame contain the dataFrames methods and atributes
type DataFrame struct {
	Columns  []string
	Values   Book
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

// String Return string to print it
func (df DataFrame) String() string {
	printDataFrame(df.Columns, df.Values, df.Index)

	return ""
}

func printDataFrame(columns []string, values Book, index []int) {

	const padding = 3
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.TabIndent|tabwriter.Debug)

	// Print Header
	header := strings.Join(columns, "\t")
	fmt.Fprintf(w, "index\t%v\n", header)

	for i, v := range index {
		var line []string
		for _, col := range columns {
			stringValue := fmt.Sprintf("%v", values[col][i])
			line = append(line, stringValue)
		}
		// Print on the table
		value := strings.Join(line[:], "\t")
		fmt.Fprintf(w, "%d\t%v\n", v, value)
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

// Head Get DataFrame with the first 5 rows
func (df DataFrame) Head() DataFrame {
	interbalBook := Book{}

	for k, v := range df.Values {
		interbalBook[k] = v[:5]
	}

	internalDataFrame := DataFrame{}
	internalDataFrame.Values = interbalBook
	internalDataFrame.Columns = df.Columns
	internalDataFrame.Index = df.Index[:5]

	return internalDataFrame
}

// Tail Get DataFrame with the last 5 rows
func (df DataFrame) Tail() DataFrame {
	interbalBook := Book{}
	totalRows := len(df.Index)

	for k, v := range df.Values {
		interbalBook[k] = v[totalRows-5:]
	}

	internalDataFrame := DataFrame{}
	internalDataFrame.Values = interbalBook
	internalDataFrame.Columns = df.Columns
	internalDataFrame.Index = df.Index[totalRows-5:]

	return internalDataFrame
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
	valuesPerColumn := make(Book)
	for i, v := range df.Columns {
		for _, r := range values[1:] {
			valuesPerColumn[v] = append(valuesPerColumn[v], stringToType(r[i]))
		}
	}
	df.Values = valuesPerColumn
	df.Shape = [2]int{len(df.Columns), len(df.Values)}

}

////////////////////////////
// DataFrame manipulation //
////////////////////////////

// Drop Eliminate 1 or more columns permanently
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

// ExportToCSV Function to save a DataFrame on CSV format
func (df DataFrame) ExportToCSV(filenamePath string) {

	var records [][]string
	records = append(records, df.Columns)

	for i := range df.Index {
		var record []string
		for _, columnName := range df.Columns {
			value := df.Values[columnName][i]
			valueString := fmt.Sprintf("%v", value)
			record = append(record, valueString)
		}
		records = append(records, record)
	}

	// Create CSV
	f, err := os.Create(filenamePath)

	defer f.Close()

	if err != nil {

		log.Fatalln("Failed to open file", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("Error writing record to file", err)
		}
	}

}

// AddColumn Add a specific column
func (df DataFrame) AddColumn(index uint, columnName string, values Page) DataFrame {

	// Insert on columns
	columns := df.Columns

	var newColumns []string
	if index == 0 {
		newColumns = append(newColumns, columnName)
		newColumns = append(newColumns, columns...)
	} else {
		newColumns = append(newColumns, columns[:index]...)
		newColumns = append(newColumns, columnName)
		newColumns = append(newColumns, columns[index:]...)
	}

	// Add values
	df.Values[columnName] = values

	// Add to DataFrame
	df.Columns = newColumns

	return df
}

// EqualDataFrame To know if two DataFrame are equals
func EqualDataFrame(df1, df2 *DataFrame) bool {
	if reflect.TypeOf(*df1) == reflect.TypeOf(*df2) {
		return true
	}

	return false
}

/////////////////////////
// STATITICS FUNCTIONS //
/////////////////////////

// Mean It's a function to calculate the Mean for one spefic column
func (b Page) Mean() float64 {
	var valuesFloat []float64
	for _, v := range b {
		switch v.(type) {
		case float64:
			if !math.IsNaN(v.(float64)) {
				valuesFloat = append(valuesFloat, v.(float64))
			}
		case int:
			valuesFloat = append(valuesFloat, float64(v.(int)))
		default:
			break
		}

	}
	// Calculate the mean
	mean := stat.Mean(valuesFloat, nil)

	return mean
}

// Max Calculate the max value in a specific column
func (b Page) Max() float64 {
	var valuesFloat []float64
	for _, v := range b {
		switch v.(type) {
		case float64:
			if !math.IsNaN(v.(float64)) {
				valuesFloat = append(valuesFloat, v.(float64))
			}
		case int:
			valuesFloat = append(valuesFloat, float64(v.(int)))
		default:
			break
		}

	}

	sort.Float64s(valuesFloat)
	var maxValue float64
	if len(valuesFloat) > 1 {
		maxValue = valuesFloat[len(valuesFloat)-1]
	} else {
		maxValue = math.NaN()
	}

	return maxValue
}

// Min Calculate the min value in a specific column
func (b Page) Min() float64 {
	var valuesFloat []float64
	for _, v := range b {
		switch v.(type) {
		case float64:
			if !math.IsNaN(v.(float64)) {
				valuesFloat = append(valuesFloat, v.(float64))
			}
		case int:
			valuesFloat = append(valuesFloat, float64(v.(int)))
		default:
			break
		}

	}

	sort.Float64s(valuesFloat)
	var minValue float64
	if len(valuesFloat) > 1 {
		minValue = valuesFloat[0]
	} else {
		minValue = math.NaN()
	}

	return minValue
}

// Describe Get a DataFrame with a sumary
// about the original DataFrame
func (df DataFrame) Describe() DataFrame {

	book := Book{}
	index := []int{0, 1, 2}

	for name := range df.Values {
		mean := df.Values[name].Mean()
		max := df.Values[name].Max()
		min := df.Values[name].Min()

		book[name] = Page{mean, max, min}

	}

	dfDescribe := DataFrame{
		Columns: df.Columns,
		Values:  book,
		Index:   index,
	}

	// Add Stats column
	rows := Page{"Mean", "Max", "Min"}
	dfDescribe = dfDescribe.AddColumn(0, "stats", rows)

	return dfDescribe
}
