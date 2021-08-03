package higor

import (
	"encoding/csv"
	"os"

	c "github.com/osmandi/higor/csv"
	"github.com/osmandi/higor/dataframe"
)

const Version string = "v0.5.0"

// ReadCSV Read a file with CSV format
func ReadCSV(filename string, csvOptions ...c.CSVOptions) dataframe.DataFrame {
	df := dataframe.NewDataFrame()

	csvInternal := &c.CSV{}
	// Default values
	csvInternal.Sep = ','
	csvInternal.NaNLayout = ""

	for _, csvOption := range csvOptions {
		csvOption(csvInternal)
	}

	csvFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer csvFile.Close()

	records := csv.NewReader(csvFile)

	// Set options
	records.Comma = csvInternal.Sep
	df.NaNLayout = csvInternal.NaNLayout

	csvLines, err := records.ReadAll()
	if err != nil {
		panic(err)
	}

	df.Columns = csvLines[0]
	df.Shape = [2]int{len(csvLines[1:]), len(csvLines[0])}

	for _, v := range csvLines[1:] {
		df.AddLine(v)
	}

	return df
}
