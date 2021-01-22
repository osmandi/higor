package higor

import (
	"math"
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
	valuesExpected["id"] = page{1, 2, 3, 4, 5}
	valuesExpected["name"] = page{"Hamish", "Anson", "Willie", "Eimile", "Rawley"}
	valuesExpected["work_remotely"] = page{false, math.NaN(), true, true, true}
	valuesExpected["salary"] = page{"$4528.90", "$1418.86", "$1311.34", "$3895.20", "$2350.92"}
	valuesExpected["age"] = page{96, math.NaN(), math.NaN(), 80, math.NaN()}
	valuesExpected["country_code"] = page{"PE", math.NaN(), "PH", "ID", "ZA"}

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
