package higor

import (
	"fmt"
)

var Version string = "v0.2.1"

// HelloHigor Print a simple message to check if Higor are installed correctly
// and print the version installed
func HelloHigor() string {

	helloMessage := fmt.Sprintf("Hello from Higor :) %s", Version)
	return helloMessage
}

// Higor interface
// TODO: Add interface to use higor as "hg" alias. Require code reorganization

// Print DataFrame section
// TODO: Print DataFrame with Index
// TODO: Print a large DataFrame
// TODO: Print head DataFrame
// TODO: Print tail dataframe

// Read DataFrame
// TODO: ReadCSV with parsing values
// TODO: ReadCSV with multiples data types
// TODO: ReadCSv with nil datatypes
// TODO: ReadCSV with an specific nan value
// TODO: ReadCSV without header
// TODO: ReadCSV with more rows than columns

// Export CSV
// TODO: Export with nils values
// TODO: Export with multiple DataTypes
// TODO: Export without header
// TODO: Export without index
