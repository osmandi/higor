package higor

import (
	"encoding/csv"
	"log"
	"os"

	c "github.com/osmandi/higor/csv"
	"github.com/osmandi/higor/dataframe"
)

const Version string = "v0.6.0"

// ReadCSV Read a file with CSV format
func ReadCSV(filename string, schema dataframe.Schema, csvOptions ...c.CSVOptions) dataframe.DataFrame {

	csvInternal := &c.CSV{}
	// Default values

	for _, csvOption := range csvOptions {
		csvOption(csvInternal)
	}

	csvFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer csvFile.Close()

	records := csv.NewReader(csvFile)

	csvLines, err := records.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.NewDataFrame(csvLines[1:], csvLines[0], schema, "")

	return df
}
