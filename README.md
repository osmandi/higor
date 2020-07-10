# higor
Dataframe for Golang, simple but powerfull


# Install

```Go
go get -v -u github.com/osmandi/higor
```

# Shape

```Go
package main

import (
    "github.com/osmandi/higor/higor"
)

func main(){
    dfOption := higor.ReadCSVOptions{
        Filename: "datasets/Measurement_summary.csv",
    }
    df := higor.ReadCSV(dfOption)
    fmt.Println(df.Shape())
}
```

# Performance tests

dataset from: https://www.kaggle.com/bappekim/air-pollution-in-seoul
Dataset name: Measurement_summary.csv


Execution with Pandas in Bash
```Bash
time python main.py
```

Execution with Higor in Bash
```Bash
go build main.go
time ./main
```

The time execution from `real XXX` and transform to milliseconds

| Command | Describe | Pandas (milliseconds) | Higor (milliseconds) |
|---------|----------|--------|-------|
| shape        | To know how many rows and columns | 1438 | 572 |
