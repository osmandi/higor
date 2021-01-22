# Higor

Dataframe for Golang, simple but powerful

## Install

```Bash
go get -v -u github.com/osmandi/higor
```

## Hello Gorld

```Go
package main

import (
	"fmt"

	"github.com/osmandi/higor"
)

func main() {
	fmt.Println(higor.PrintHelloHigor())
}
```

`Output:`
```Bash
Hello from higor
```

## Read a CSV

```Go
package main

import (
	"github.com/osmandi/higor"
)

func main() {
    dfHigor := higor.NewDataFrame("examples/data/example1.csv")
    dfHigor.ReadCSV()
	fmt.Println(dfHigor.Head())	
	fmt.Println(dfHigor.Tail())
}
```
`Output:`
```Bash
index   |id   |name     |work_remotely   |salary     |age   |country_code
0       |1    |Hamish   |false           |4528.90   |96    |PE
1       |2    |Anson    |NaN             |1418.86   |NaN   |NaN
2       |3    |Willie   |true            |1311.34   |NaN   |PH
3       |4    |Eimile   |true            |3895.20   |80    |ID
4       |5    |Rawley   |true            |2350.92   |NaN   |ZA

index   |id    |name       |work_remotely   |salary     |age   |country_code
95      |96    |NaN        |false           |NaN        |54    |GF
96      |97    |Novelia    |true            |3948.23   |NaN   |JP
97      |98    |Maegan     |false           |2905.48   |48    |UA
98      |99    |Andreana   |true            |3732.29   |73    |CN
99      |100   |Freeman    |false           |2850.99   |39    |TH
```

Credits to [mockaroo](https://www.mockaroo.com/) by CSV generator content on `examples/data` folder