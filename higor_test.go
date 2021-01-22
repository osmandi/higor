package higor

import (
	"math"
	"reflect"
	"testing"
)

func TestPrintHelloHigor(t *testing.T) {
	value := PrintHelloHigor()
	if value != "Hello from higor" {
		t.Errorf("HellowHigor failed")
	}
}

func TestHead(t *testing.T) {
	valuesExpected := Book{}
	valuesExpected["id"] = Page{1, 2, 3, 4, 5}
	valuesExpected["name"] = Page{"Hamish", "Anson", "Willie", "Eimile", "Rawley"}
	valuesExpected["work_remotely"] = Page{false, math.NaN(), true, true, true}
	valuesExpected["salary"] = Page{4528.90, 1418.86, 1311.34, 3895.20, 2350.92}
	valuesExpected["age"] = Page{96, math.NaN(), math.NaN(), 80, math.NaN()}
	valuesExpected["country_code"] = Page{"PE", math.NaN(), "PH", "ID", "ZA"}

	dfHigor := NewDataFrame("examples/data/example1.csv")
	dfHigor.ReadCSV()
	dfHigorHead := dfHigor.Head()
	valuesResult := dfHigorHead.Values

	for k, v := range valuesResult {
		for i, element := range v {
			switch element.(type) {
			case float64:

				if math.IsNaN(element.(float64)) {
					if !math.IsNaN(valuesExpected[k][i].(float64)) {
						t.Errorf("Column: \"%s\", expected: %v recived: %v", k, valuesExpected[k], v)
					}
				} else if element != valuesExpected[k][i] {

					t.Errorf("Column: \"%s\", expected: %v recived: %v", k, valuesExpected[k], v)
				}
			default:

				if element != valuesExpected[k][i] {

					t.Errorf("Column: \"%s\", expected: %v recived: %v", k, valuesExpected[k], v)
				}
			}

		}
	}
}

func TestTail(t *testing.T) {

	// Values expected
	valuesExpected := Book{}
	valuesExpected["id"] = Page{96, 97, 98, 99, 100}
	valuesExpected["name"] = Page{math.NaN(), "Novelia", "Maegan", "Andreana", "Freeman"}
	valuesExpected["work_remotely"] = Page{false, true, false, true, false}
	valuesExpected["salary"] = Page{math.NaN(), 3948.23, 2905.48, 3732.29, 2850.99}
	valuesExpected["age"] = Page{54, math.NaN(), 48, 73, 39}
	valuesExpected["country_code"] = Page{"GF", "JP", "UA", "CN", "TH"}

	// Index expected
	indexExpected := []int{95, 96, 97, 98, 99}

	// Get Result
	dfHigor := NewDataFrame("examples/data/example1.csv")
	dfHigor.ReadCSV()
	dfHigorTail := dfHigor.Tail()
	valuesResult := dfHigorTail.Values
	indexResult := dfHigorTail.Index

	// Values test
	for k, v := range valuesResult {
		for i, element := range v {
			switch element.(type) {
			case float64:

				if math.IsNaN(element.(float64)) {
					if !math.IsNaN(valuesExpected[k][i].(float64)) {
						t.Errorf("Column: \"%s\", expected: %v recived: %v", k, valuesExpected[k], v)
					}
				} else if element != valuesExpected[k][i] {

					t.Errorf("Column: \"%s\", expected: %v recived: %v", k, valuesExpected[k], v)
				}
			default:

				if element != valuesExpected[k][i] {

					t.Errorf("Column: \"%s\", expected: %v recived: %v", k, valuesExpected[k], v)
				}
			}

		}
	}

	// Index test
	if !reflect.DeepEqual(indexExpected, indexResult) {
		t.Errorf("Index error, expected: %v recived: %v", indexExpected, indexResult)
	}

}

func TestDrop(t *testing.T) {

	// Result expected
	expectedColumns := []string{"id", "work_remotely", "salary", "country_code"}

	// Get result
	dfHigor := NewDataFrame("examples/data/example1.csv")
	dfHigor.ReadCSV()
	dfHigor.Drop("name", "age")

	if !reflect.DeepEqual(expectedColumns, dfHigor.Columns) {
		t.Errorf("Columns error, expected: %v recived: %v", expectedColumns, dfHigor.Columns)
	}

}

func TestMean(t *testing.T) {
	dfHigor := NewDataFrame("examples/data/example1.csv")
	dfHigor.ReadCSV()

	// Test float64 number
	valueResult := dfHigor.Values["salary"].Mean()
	valueExpected := 2963.707882352941

	if valueResult != valueExpected {
		t.Errorf("Median error, expected: %v, result: %v", valueExpected, valueResult)
	}

	// Test int
	valueIntResult := math.Round(dfHigor.Values["id"].Mean()*100) / 100
	valueIntExpected := 49.54

	if valueIntResult != valueIntExpected {
		t.Errorf("Median error, expected: %v, result: %v", valueIntExpected, valueIntResult)
	}
}
