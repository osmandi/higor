# Higor

Dataframe for Golang, simple but powerful

# Install

```Bash
go get -v -u github.com/osmandi/higor
```

# How to use

### Read CSV file

```Go
package main

func main() {
    dfHigor := higor.NewDataFrame("csv_path.csv")
    dfHigor.Sep = ',' // Set only if the comma separator is different to ','
    dfHigor.ReadCSV()
    fmt.Println(dfHigor)
}
```

### To know the mean for a column

```Go
column := "col_name"
fmt.Printf("The mean for the column %s is %v\n", column, dfHigor.Values[column].Mean())
```