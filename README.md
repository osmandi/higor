# Higor

Dataframe for Golang, simple but powerful

## Install

```Bash
go get -v -u github.com/osmandi/higor
```

# Hellow Gorld

- From CSV files

```Go
package main

func main() {
    dfHigor := higor.NewDataFrame("exampleData/example1.csv")
    dfHigor.ReadCSV()
    fmt.Println(dfHigor)
}
```

Credits to [mockaroo](https://www.mockaroo.com/) by CSV generator content on `examples/data` folder