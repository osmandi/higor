package main

import (
	"higor"
)

func main() {
	dfHigor := higor.NewDataFrame("examples/data/example1.csv")
	dfHigor.ReadCSV()
	dfHigor.Head(2)
}
