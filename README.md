# Higor

Dataframe for Golang, simple but powerful

## Install

```Bash
go get -v -u github.com/osmandi/higor
```

# Load information

- From CSV files

```Go
package main

func main() {
    dfHigor := higor.NewDataFrame("csv_path.csv")
    dfHigor.Sep = ',' // Set only if the comma separator is different to ','
    dfHigor.SafeMode = false // Skip the runtime if there is an error (false is disactive)
    dfHigor.ReadCSV()
    fmt.Println(dfHigor)
}
```

## How to use

- Calculate the mean for a specific column

```Go
column := "col_name"
fmt.Printf("The mean for the column %s is %v\n", column, dfHigor.Values[column].Mean())
```

- Drop one or more columns

```Go
dfHigor.Drop("col1", "col2")
```