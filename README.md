![higor_logo](higor_logo.jpg)

------

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) 
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/osmandi/higor)
[![Build Status](https://travis-ci.com/osmandi/higor.svg?branch=master)](https://travis-ci.com/osmandi/higor)
[![Go Report Card](https://goreportcard.com/badge/github.com/osmandi/higor)](https://goreportcard.com/report/github.com/osmandi/higor)
[![codecov](https://codecov.io/gh/osmandi/higor/branch/master/graph/badge.svg)](https://codecov.io/gh/osmandi/higor)
[![gocover](https://gocover.io/_badge/github.com/osmandi/higor?nocache=wapty)](https://gocover.io/github.com/osmandi/higor?nocache=wapty)

# Higor

Dataframe for Golang as column oriented, simple but powerful.

## Why Golang for Data Engineering?

Go has a multiple features that help on this topic, for example:
- It's very easy implement the concurrency to processing a lot of data using native resources (without external libraries).
- Its compilation is fast.
- It has a native library to implement tests.
- Its sintax is simple such as Python but powerfull as C language.
- Its mascot is very cute (this is a personal opinion).

## Why Higor?

Currently, Python is used in all steps for a Machine Learning project, from Data Engineering to Machine Learning model train. However, to make this possible, Python uses precompiled code (i.e. Numpy uses C and Spark uses Scala) instead a native way. Doing this, causes syntaxis changes, depending on the external library used.

The aim of Higor is to provide a library that allows you to work with different amounts of data using Golang with having a simple interface.

> Remember, Kubernetes is made in Golang. Can you imagine that power in Data Engineering and Big Data?

## Install

```Bash
go get -v -u github.com/osmandi/higor
```

## Say hello to Higor

**sample.csv** Content:
```
colFloat64,colString,colBool,colDatetime,colAny
2,no,true,2021-01-30,1.2
5,hello,false,none,true
none,hello,false,none,none
none,none,false,none,helloText
5,hello,false,none,2021-02-03
```

```Go
package main

import (
	"fmt"
	"time"

	hg "github.com/osmandi/higor"
   "github.com/osmandi/higor/csv"
)


func main() {
	df := hg.ReadCSV("sample.csv", csv.NaNLayout("none"))
	fmt.Println(df)
}
```

Result:

```Bash
+------------+-----------+---------+-------------------------------+-------------------------------+
| COLFLOAT64 | COLSTRING | COLBOOL |          COLDATETIME          |            COLANY             |
+------------+-----------+---------+-------------------------------+-------------------------------+
| 2          | no        | true    | 2021-01-30 00:00:00 +0000 UTC | 1.2                           |
| 5          | hello     | false   | NaN                           | true                          |
| NaN        | hello     | false   | NaN                           | NaN                           |
| NaN        | NaN       | false   | NaN                           | helloText                     |
| 5          | hello     | false   | NaN                           | 2021-02-03 00:00:00 +0000 UTC |
+------------+-----------+---------+-------------------------------+-------------------------------+
```

## How to contribute?
- Give this repo a star.
- Create tutorials about Data Engineering with Go.
- Use this library and if you have some issues please put it on issues section with the Data.
- If you need a specific feature, please create a PR to README.md to request it.

## What does higor know how to do today?
- ReadCSV (With custom NaN, custom Separator)
- Print DataFrames (completly, head, tail)

## Next features:
- Column manipulate (select, drop, create, filter)
- Implement GroupBy functions (count, max, mean, median, min, sum)
- Implement concurrency on internal functions
