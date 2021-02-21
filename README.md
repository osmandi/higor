[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# Higor

Dataframe for Golang, simple but powerful

<!--

## Why Golang to Data Engineering?

Go has a multiple features that help on this topic, for example:
- It's very easy implement the concurrency to processing a lot of data with nativie way (without external libraries)
- His compilation is fast.
- Has your native library to implement tests.
- His sintax is simple such as Python but powerfull as C language.
- Its mascot is very cute (this is a personal opinion)

## Why Higor?

Actualy Python is used in all steps for a Machine Learning project, from Data Engineering to Machine Learning model train, to adapt this Python doesn't to do native instead use precompiled code to another language, for this way Python change syntax depending of the external libraly. For example: Numpy use C and Spark uses Scala.

Depending how much data do you have in Python you need different libraries (Pandas -> Dask -> PySpark)

The proposal with `Higor` is uses Golang during Data Engineering process no matter how much data do you have with the same library for that keeping the same syntax code.

If you are a Data Engineer or Golang developer this question is for you: If Docker y Kubernetes están hechos en Go, ¿qué hace falta para implementar Go para Big Data?

-->

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
