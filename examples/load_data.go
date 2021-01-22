package main

import (
	"fmt"

	"github.com/osmandi/higor"
)

func main() {
	dfHigor := higor.NewDataFrame("examples/data/example1.csv")
	dfHigor.ReadCSV()

	fmt.Println(dfHigor.Head())

}
