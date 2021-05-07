![higor_logo](higor_logo.jpg)

------

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) 
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/osmandi/higor)
[![Build Status](https://travis-ci.com/osmandi/higor.svg?branch=master)](https://travis-ci.com/osmandi/higor)
[![Go Report Card](https://goreportcard.com/badge/github.com/osmandi/higor)](https://goreportcard.com/report/github.com/osmandi/higor)
[![codecov](https://codecov.io/gh/osmandi/higor/branch/master/graph/badge.svg)](https://codecov.io/gh/osmandi/higor)

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

	hg "github.com/osmandi/higor"
	"github.com/osmandi/higor/dataframe"
)

func main() {
	fmt.Println(hg.HelloHigor())
	fmt.Println("")

	// sample.csv content:
	/*
		col1,col2,col3,col4
		1,2,no,true
		3,5,hello,false
	*/
	schema := dataframe.Book{
		dataframe.PageFloat64{},
		dataframe.PageFloat64{},
		dataframe.PageString{},
		dataframe.PageBool{},
	}
	df := hg.ReadCSV("sample.csv", dataframe.Schema(schema))
	fmt.Println(df)
}
```

Result:

```Bash
Hello from Higor :) v0.3.0

+------+------+-------+-------+
| COL1 | COL2 | COL3  | COL4  |
+------+------+-------+-------+
|    1 |    2 | no    | true  |
|    3 |    5 | hello | false |
+------+------+-------+-------+
```

## How to contribute?
- Give this repo a star.
- Create tutorials about Data Engineering with Go.
- Use this library and if you have some issues please put it on issues section with the Data.
- If you need a specific feature, please create a PR to README.md to request it.

# Releases version

## v0.3.0: DataType by column
- Delete footer
- Set schema to read a CSV with parsing values
- Save values with a specific DataType slice instead a interface
- DataType for datetime values

## v0.3.1: DataTypes reading unexpeding
- ReadCSV with nil values
- ReadCSV custom nil values setting

## v0.3.2: Improve DataFrame print it
- Print tail dataframe
- Print head DataFrame
- Print DataFrame with Index

## v0.3.3: Automatic reading DataType
- Read CSV with automatic DataTypes setting

## v0.3.4: Improve importing CSV
- Export with nils values
- Export without header
- Export without index
