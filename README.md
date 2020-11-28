# Higor

Dataframe for Golang, simple but powerful

# Install

```Bash
go get -v -u github.com/osmandi/higor
```

# How to use

```Go
package main

func main() {

	dfHigor := higor.NewDataFrame("csv_path.csv")
	dfHigor.Sep = ',' // Set only if the comma separaor is different to ','
    dfHigor.ReadCSV()
    fmt.Println(dfHigor)
}
```