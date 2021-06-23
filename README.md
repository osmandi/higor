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

```Go
package main

import (
	"fmt"
	"time"

	hg "github.com/osmandi/higor"
	"github.com/osmandi/higor/dataframe"
)

type myTime time.Time

func main() {
	fmt.Println(hg.HelloHigor())
	fmt.Println("")

	// sample.csv content:
	/*
		colFloat64,colString,colBool,colDatetime,colAny
		2,no,true,2021-01-30,1.2
		5,hello,false,none,true
		none,hello,false,none,none
		none,none,false,none,helloText
		5,hello,false,none,2021-02-03
	*/

	schema := dataframe.Book{
		dataframe.PageFloat64{},  // colFloat64
		dataframe.PageString{},   // colString
		dataframe.PageBool{},     // colBool
		dataframe.PageDatetime{}, // colDatetime
		dataframe.PageAny{},      // colAny
	}
	dateformat := "YYYY-MM-DD"
	none := "none" // Support for custom NaN values, default is ""
	df := hg.ReadCSV("sample.csv", dataframe.Schema(schema), dataframe.Dateformat(dateformat), dataframe.None(none))
	fmt.Println(df)
}
```

Result:

```Bash
Hello from Higor :) v0.3.1

+------------+-----------+---------+-------------------------------+------------+
| COLFLOAT64 | COLSTRING | COLBOOL |          COLDATETIME          |   COLANY   |
+------------+-----------+---------+-------------------------------+------------+
|          2 | no        | true    | 2021-01-30 00:00:00 +0000 UTC |        1.2 |
|          5 | hello     | false   | NaN                           | true       |
| NaN        | hello     | false   | NaN                           | NaN        |
| NaN        | NaN       | false   | NaN                           | helloText  |
|          5 | hello     | false   | NaN                           | 2021-02-03 |
+------------+-----------+---------+-------------------------------+------------+
```

## Column types with support for NaN values
- PageString: To save any string value. For example: "Hello Higor", "value".
- PageDatetime: To save datetime values. For example: "2021-01-30 00:00:00 +0000 UTC".
- PageFloat64: To save numbers integers and floats. For example: 1, 1.2.
- PageAny: To save any data type. For example: 1, "dos", true, etc.

## Column types without support for NaN values
- PageBool: To save values type bool. For example: true, false.

## How to contribute?
- Give this repo a star.
- Create tutorials about Data Engineering with Go.
- Use this library and if you have some issues please put it on issues section with the Data.
- If you need a specific feature, please create a PR to README.md to request it.

# Releases version

> Actual version: v0.3.1

## v0.3.0: DataType by column
- [x] Delete footer
- [x] Set schema to read a CSV with parsing values
- [x] Save values with a specific DataType slice instead a interface
- [x] DataType for datetime values
- [x] DataType for datetime values with custom format

## v0.3.1: DataTypes reading unexpeding
- [x] ReadCSV with none values - PageFloat64
- [x] ReadCSV with none values - PageAny
- [x] ReadCSV with none values - PageString
- [x] ReadCSV with none values - PageDatetime
- [x] ReadCSV custom none values setting
- [x] ReadCSV correct print NaN values
- [x] Message for incorrect schema
- [x] Set what PageTypes has None values support
- [x] Goal: Read [iris.csv](https://gist.github.com/netj/8836201) file successfully

## v0.3.2: Improve DataFrame print it
- [x] Print tail dataframe
- [x] Print head DataFrame
- [ ] Print DataFrame with Index
- [ ] Goal: Read a large csv file

## v0.3.3: STOP! Improve codecoverage percent to >= 90%

## v0.3.4: Improve importing and exporting CSV
- [ ] Export with nils values
- [ ] Export without header
- [ ] Export without index
- [ ] Read CSV with automatic DataTypes setting

## v0.3.5: Improve features
- [ ] Custom datetime format by each column
- [ ] Set a way to ignore rows with NaN values
- [ ] Set a way to change to PageAny if found None values whe that is not supported on the schema

## v0.4.x: Column manipulate (select, update, drop, create)

## v0.5.x: Implement statistics functions (Mean, Median, etc)

## v0.6.x: Implement concurrency by default
