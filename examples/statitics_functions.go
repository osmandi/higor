package main

import (
	"fmt"

	"github.com/osmandi/higor"
)

func main() {
	dfHigor := higor.NewDataFrame("examples/data/example1.csv")
	dfHigor.ReadCSV()

	// Mean
	fmt.Println(dfHigor.Values["id"].Mean())     // Int
	fmt.Println(dfHigor.Values["salary"].Mean()) // Float64
	fmt.Println(dfHigor.Values["name"].Mean())   // String

	// Min
	fmt.Println(dfHigor.Values["id"].Min())     // Int
	fmt.Println(dfHigor.Values["salary"].Min()) // Float64
	fmt.Println(dfHigor.Values["name"].Min())   // string

	// Max
	fmt.Println(dfHigor.Values["id"].Max())     // Int
	fmt.Println(dfHigor.Values["salary"].Max()) // Float64
	fmt.Println(dfHigor.Values["name"].Max())   // String
}
