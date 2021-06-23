package dataframe

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/olekukonko/tablewriter"
)

// PageAny To save any data type, when if you don't know what data type is and with support for NaN values
type PageAny []interface{}

// PageString Data type for string values with support for NaN values
type PageString []string

// PageBool Data type for boolean values. Not support for NaN values
type PageBool []bool

// PageFloat64 Data type for numbers and float values with support for NaN values
type PageFloat64 []float64

// PageDatetime To date dates with support for NaN values
type PageDatetime []time.Time

// Book Interface to save a DataFrame
type Book []interface{}

// DataFrame Structure for DataFrame
type DataFrame struct {
	Columns []string
	Values  Book
	Shape   [2]int
}

// IsEqual to kown if two DataFrame are equal
func IsEqual(dataFrame1, dataFrame2 DataFrame) (bool, string) {
	columnComparation := reflect.DeepEqual(dataFrame1.Columns, dataFrame2.Columns)

	// Columns comparation
	if !columnComparation {
		return false, "Columns are different"
	}

	// Values comparation
	valuesDataFrame1 := dataFrame1.Values
	valuesDataFrame2 := dataFrame2.Values

	if len(valuesDataFrame1) != len(valuesDataFrame2) {
		return false, ""
	}

	for i, v := range valuesDataFrame1 {
		switch v.(type) {
		case PageString:
			for i2, v2 := range v.(PageString) {
				v3 := valuesDataFrame2[i].(PageString)[i2]
				if v2 != v3 {
					return false, fmt.Sprintf("PageStringError: v2: %v | v3: %v", v2, v3)
				}
			}
		case PageBool:
			for i2, v2 := range v.(PageBool) {
				v3 := valuesDataFrame2[i].(PageBool)[i2]
				if v2 != v3 {
					return false, fmt.Sprintf("PageBoolError: v2: %v | v3: %v", v2, v3)
				}
			}
		case PageFloat64:
			for i2, v2 := range v.(PageFloat64) {
				v3 := valuesDataFrame2[i].(PageFloat64)[i2]
				if math.IsNaN(v2) != math.IsNaN(v3) {
					return false, fmt.Sprintf("PageFloat64ErrorIsNaN: v2: %v | v3: %v | v2 == v3: %v", math.IsNaN(v2), math.IsNaN(v3), math.IsNaN(v2) == math.IsNaN(v3))
				} else if v2 != v3 && !math.IsNaN(v2) {
					return false, fmt.Sprintf("PageFloat64ErrorComparation: v2: %v | v3: %v | v2 == v3: %v", v2, v3, v2 == v3)
				}
			}
		case PageAny:
			for i2, v2 := range v.(PageAny) {
				if fmt.Sprintf("%T", v2) == "float64" {
					if math.IsNaN(v2.(float64)) != math.IsNaN(valuesDataFrame2[i].(PageAny)[i2].(float64)) {
						return false, ""
					}

				} else if v2 != valuesDataFrame2[i].(PageAny)[i2] {
					fmt.Println("PageAny")
					fmt.Printf("v2: %v, v3: %v\n", v2, valuesDataFrame2[i].(PageAny)[i2])
					fmt.Printf("v2 == v3: %v\n", v2 == valuesDataFrame2[i].(PageAny)[i2])
					return false, ""
				}
			}

		}
	}

	return true, ""

}

// IsNaN To know if a interface{} is NaN
func IsNaN(a interface{}) bool {
	switch a.(type) {
	case time.Time:
		return a.(time.Time).IsZero()
	case string:
		if a.(string) == "" {
			return true
		}
	}

	return false
}

/////////
// CSV /
///////

// CSVCreatorMock sample csv to tests
func CSVCreatorMock(data [][]string, separator rune) *os.File {
	// Temp file
	tmpFile, err := ioutil.TempFile(os.TempDir(), "higorCSVTest-*.csv")
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(tmpFile)
	writer.Comma = separator
	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			log.Fatal(err)
		}
	}

	writer.Flush()

	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}
	return tmpFile

}

// CSVChecker To Check csv result
func CSVChecker(dataExpected, dataResult [][]string, t *testing.T) {
	if !reflect.DeepEqual(dataExpected, dataResult) {
		t.Errorf("Header with errors. \nExpected: \n%s. \nReceived: \n%s", dataExpected, dataResult)
	}
}

// DataFrameChecker To check if two DataFrame are equal
func DataFrameChecker(dfExpected, dfResult DataFrame, t *testing.T) {
	isEqual, message := IsEqual(dfExpected, dfResult)
	if !isEqual {
		t.Errorf("dfExpected and dfResult are distinct: %s.\ndfExpected: \n%v \ndfResult: \n%v", message, dfExpected, dfResult)
		t.Errorf("Values:\n - dfExpected.Values: %v\n - dfResult.Values: %s\n", dfExpected.Values, dfResult.Values)
		t.Errorf("Columns:\n - dfExpected.Columns: %v\n - dfResult.Columns: %s\n", dfExpected.Columns, dfResult.Columns)
	}

}

//////

// CSV type
type CSV struct {
	Sep        rune
	LineString string
	LazyQuotes bool
	Schema     Book
	Dateformat string
	None       string
}

// ErrorChecker to kown if there are error
func ErrorChecker(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ErrorSchema to kown if there is error when read an schema
func ErrorSchema(colName string, pageName string, dataValue interface{}, err error) {
	if err != nil {
		log.Fatalf("The schema \"%s\" is incorrect for column \"%s\". Error parsing: \"%v\".\n\t\t You can use PageAny to save any data value.", pageName, colName, dataValue)
	}
}

// CSVOption alternative parameters
type CSVOption func(c *CSV)

// Sep CSV separator in rune type: ',', ';', '|', etc...
func Sep(separator rune) CSVOption {
	return func(c *CSV) {
		c.Sep = separator
	}
}

// Schema set the schema
func Schema(schema Book) CSVOption {
	return func(c *CSV) {
		c.Schema = schema
	}
}

// None to set custom value for None
func None(none string) CSVOption {
	return func(c *CSV) {
		c.None = none
	}
}

// Dateformat Set date format in all columns
func Dateformat(dateformat string) CSVOption {
	return func(c *CSV) {
		c.Dateformat = dateformat
	}
}

// GetValues get all values
func (df DataFrame) GetValues() [][]string {
	return trasposeRows(df)
}

func trasposeRows(df DataFrame) [][]string {

	data := make([][]string, df.Shape[0])

	for i := range df.Columns {
		colValues := df.Values[i]
		switch colValues.(type) {
		case PageString:
			for i2, v2 := range colValues.(PageString) {
				if IsNaN(v2) {
					data[i2] = append(data[i2], "NaN")
				} else {
					data[i2] = append(data[i2], v2)

				}
			}

		case PageFloat64:
			for i2, v2 := range colValues.(PageFloat64) {
				data[i2] = append(data[i2], fmt.Sprintf("%v", v2))
			}

		case PageBool:
			for i2, v2 := range colValues.(PageBool) {
				data[i2] = append(data[i2], fmt.Sprintf("%v", v2))
			}

		case PageAny:
			for i2, v2 := range colValues.(PageAny) {
				data[i2] = append(data[i2], fmt.Sprintf("%v", v2))
			}

		case PageDatetime:
			for i2, v2 := range colValues.(PageDatetime) {
				if IsNaN(v2) {
					data[i2] = append(data[i2], "NaN")
				} else {
					data[i2] = append(data[i2], fmt.Sprintf("%v", v2))
				}
			}
		}
	}
	return data
}

// exportCSV To export a dataframe to CSV file
func exportCSV(filename string, data [][]string, opts ...CSVOption) {
	csvInternal := &CSV{}
	csvInternal.Sep = ','

	for _, opt := range opts {
		opt(csvInternal)
	}

	csvFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	ErrorChecker(err)
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Comma = csvInternal.Sep
	err = csvWriter.WriteAll(data)
	ErrorChecker(err)

}

func (df DataFrame) String() string {
	tableString := &strings.Builder{}
	data := trasposeRows(df)
	table := tablewriter.NewWriter(tableString)
	table.SetHeader(df.Columns)
	table.AppendBulk(data)
	table.SetBorder(true)
	table.SetCenterSeparator("+")

	table.Render()

	return tableString.String()
}

// ToCSV Export DataFrame to CSV
func (df DataFrame) ToCSV(filename string) {
	data := [][]string{}
	data = append(data, df.Columns)

	dfInternal := DataFrame{
		Columns: df.Columns,
		Values:  df.Values,
		Shape:   df.Shape,
	}

	data = append(data, dfInternal.GetValues()...)
	exportCSV(filename, data)
}

// Head get first 5 rows
func (df DataFrame) Head() DataFrame {
	valuesInternal := Book{}
	numberToHead := 5
	if df.Shape[1] > numberToHead {
		for _, v := range df.Values {
			switch v.(type) {
			case PageString:
				valuesInternal = append(valuesInternal, v.(PageString)[:numberToHead])
			case PageFloat64:
				valuesInternal = append(valuesInternal, v.(PageFloat64)[:numberToHead])
			case PageBool:
				valuesInternal = append(valuesInternal, v.(PageBool)[:numberToHead])
			case PageAny:
				valuesInternal = append(valuesInternal, v.(PageAny)[:numberToHead])
			case PageDatetime:
				valuesInternal = append(valuesInternal, v.(PageDatetime)[:numberToHead])
			}
		}

	} else {
		for _, v := range df.Values {
			switch v.(type) {
			case PageString:
				valuesInternal = append(valuesInternal, v.(PageString))
			case PageFloat64:
				valuesInternal = append(valuesInternal, v.(PageFloat64))
			case PageBool:
				valuesInternal = append(valuesInternal, v.(PageBool))
			case PageAny:
				valuesInternal = append(valuesInternal, v.(PageAny))
			case PageDatetime:
				valuesInternal = append(valuesInternal, v.(PageDatetime))
			}
		}

	}

	dfInternal := DataFrame{
		Columns: df.Columns,
		Values:  valuesInternal,
		Shape:   [2]int{df.Shape[0], numberToHead},
	}

	return dfInternal

}

// Tail to get last five rows
func (df DataFrame) Tail() DataFrame {
	valuesInternal := Book{}
	numberToTail := 5
	totalRows := df.Shape[1]
	if totalRows > numberToTail {
		for _, v := range df.Values {
			switch v.(type) {
			case PageString:
				valuesInternal = append(valuesInternal, v.(PageString)[totalRows-numberToTail:])
			case PageFloat64:
				valuesInternal = append(valuesInternal, v.(PageFloat64)[totalRows-numberToTail:])
			case PageBool:
				valuesInternal = append(valuesInternal, v.(PageBool)[totalRows-numberToTail])
			case PageAny:
				valuesInternal = append(valuesInternal, v.(PageAny)[totalRows-numberToTail:])
			case PageDatetime:
				valuesInternal = append(valuesInternal, v.(PageDatetime)[totalRows-numberToTail:])
			}
		}

	} else {
		for _, v := range df.Values {
			switch v.(type) {
			case PageString:
				valuesInternal = append(valuesInternal, v.(PageString))
			case PageFloat64:
				valuesInternal = append(valuesInternal, v.(PageFloat64))
			case PageBool:
				valuesInternal = append(valuesInternal, v.(PageBool))
			case PageAny:
				valuesInternal = append(valuesInternal, v.(PageAny))
			case PageDatetime:
				valuesInternal = append(valuesInternal, v.(PageDatetime))
			}
		}

	}

	dfInternal := DataFrame{
		Columns: df.Columns,
		Values:  valuesInternal,
		Shape:   [2]int{df.Shape[0], numberToTail},
	}

	return dfInternal

}
