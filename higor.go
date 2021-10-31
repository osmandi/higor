package higor

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"

	c "github.com/osmandi/higor/csv"
	"github.com/osmandi/higor/dataframe"
)

const Version string = "v0.6.0"

func loadDataFrame(filename string, schema dataframe.Schema, csvOptions ...c.CSVOptions) dataframe.DataFrame {

	csvInternal := &c.CSV{}
	// Default values
	csvInternal.Sep = ','

	for _, csvOption := range csvOptions {
		csvOption(csvInternal)
	}

	csvFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer csvFile.Close()

	records := csv.NewReader(csvFile)
	records.Comma = csvInternal.Sep

	csvLines, err := records.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.NewDataFrame(csvLines[1:], csvLines[0], schema, "")

	return df
}

// ReadCSV Read a file with CSV format
func ReadCSV(filename string, schema dataframe.Schema, csvOptions ...c.CSVOptions) dataframe.DataFrame {
	filesMatch, _ := filepath.Glob(filename)
	df := dataframe.DataFrame{}
	dfs := []dataframe.DataFrame{}

	for _, v := range filesMatch {
		fmt.Printf("Loading... %s\n", v)
		dfIterator := loadDataFrame(v, schema, csvOptions...)
		dfs = append(dfs, dfIterator)
	}

	df = dataframe.Concat(dfs...)

	return df

}
