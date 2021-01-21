package main

import (
	"github.com/osmandi/higor"
)

func main() {
	dfHigor := higor.NewDataFrame("examples/data/example1.csv")
	dfHigor.ReadCSV()

	// Print the first 2 rows, (you can input another rows numbers)
	dfHigor.Head(2)

}
