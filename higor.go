package higor

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

// DataFrame contain the dataFrames methods and atributes
type DataFrame struct {
	Columns  []string
	Values   [][]string
	Shape    [2]int
	Sep      rune
	Filename string
}

// PrintHelloHigor To get greets from higor library
func PrintHelloHigor() string {
	return "Hello from higor"
}

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

// String Return string to print it
func (df DataFrame) String() string {
	printDataFrame(df.Columns, df.Values)

	return ""
}

// NewDataFrame set default variables to read dataframe
func NewDataFrame(filename string) *DataFrame {
	return &DataFrame{
		Sep:      ',',
		Filename: filename,
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
	index := 0
	for {
		var lines []string
		// Read in a row. Check if we are at the end of the line
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Addd index value
		if index >= 1 {
			lines = append(lines, strconv.Itoa(index))
		}

		for _, value := range record {
			lines = append(lines, value)

		}

		values = append(values, lines)

		index++

	}

	// Set DataFrame parameters
	df.Columns = values[0]
	df.Values = values[1:]
	df.Shape = [2]int{len(df.Columns), len(df.Values)}

}
