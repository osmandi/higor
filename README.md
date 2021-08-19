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

## How to use Higor?

[Download iris.csv](https://gist.githubusercontent.com/netj/8836201/raw/6f9306ad21398ea43cba4f7d537619d0e07d5ae3/iris.csv)

```Go
package main

import (
	"fmt"

	"github.com/osmandi/higor"
)

func main() {
	// Print Head
	df := higor.ReadCSV("iris.csv")
	fmt.Println("Head:")
	fmt.Println(df.Head(5))

	// Print Tail
	fmt.Println("Tail:")
	fmt.Println(df.Tail(5))

	// Select columns
	fmt.Println(`Select "sepal.width","sepal.width","variety" columns:`)
	fmt.Println(df.Select("sepal.width", "sepal.width", "variety").Head(5))

	// Apply filters (WhereEqual, WhereNotEqual, WhereGreater, WhereLess, WhereGreaterOrEqual, WhereLessOrEqual)
	fmt.Println(`Where "variety" == "Setosa":`)
	fmt.Println(df.WhereEqual("variety", "Setosa").Head(5))
	fmt.Println(`Where "sepal.length" > 5:`)
	fmt.Println(df.WhereGreater("sepal.length", float64(5)).Head(5))
	fmt.Println(`Where "variety" != "Setosa":`)
	fmt.Println(df.WhereNotEqual("variety", "Setosa").Head(5))
	fmt.Println(`Where "sepal.width" > 4 && "variety" != "Versicolor":`)
	fmt.Println(df.WhereGreater("sepal.width", float64(4)).WhereNotEqual("variety", "Versicolor").Tail(5))

	// Drop one or more columns
	fmt.Println(`Drop "sepal.width", "sepal.length" columns:`)
	df.Drop("sepal.width", "sepal.length")
	fmt.Println(df.Head(5))

	// Insert and Sum
	fmt.Println(`Insert new column called "petal.length2" where "petal.length" + 100`)
	df.Insert("petal.length2", df.Column("petal.length").Add(float64(100)))
	fmt.Println(df.Head(5))

	// Insert and concat
	fmt.Println(`Insert new column called "variety2" where "variety" column + " Iris"`)
	df.Insert("variety2", df.Column("variety").Add(" Iris"))
	fmt.Println(df.Head(5))
}
```

Result:

```Bash
Head:
+---+--------------+-------------+--------------+-------------+---------+
|   | sepal.length | sepal.width | petal.length | petal.width | variety |
+---+--------------+-------------+--------------+-------------+---------+
| 0 | 5.1          | 3.5         | 1.4          | 0.2         | Setosa  |
| 1 | 4.9          | 3           | 1.4          | 0.2         | Setosa  |
| 2 | 4.7          | 3.2         | 1.3          | 0.2         | Setosa  |
| 3 | 4.6          | 3.1         | 1.5          | 0.2         | Setosa  |
| 4 | 5            | 3.6         | 1.4          | 0.2         | Setosa  |
+---+--------------+-------------+--------------+-------------+---------+

Tail:
+-----+--------------+-------------+--------------+-------------+-----------+
|     | sepal.length | sepal.width | petal.length | petal.width |  variety  |
+-----+--------------+-------------+--------------+-------------+-----------+
| 145 | 6.7          | 3           | 5.2          | 2.3         | Virginica |
| 146 | 6.3          | 2.5         | 5            | 1.9         | Virginica |
| 147 | 6.5          | 3           | 5.2          | 2           | Virginica |
| 148 | 6.2          | 3.4         | 5.4          | 2.3         | Virginica |
| 149 | 5.9          | 3           | 5.1          | 1.8         | Virginica |
+-----+--------------+-------------+--------------+-------------+-----------+

Select "sepal.width","sepal.width","variety" columns:
+---+-------------+-------------+---------+
|   | sepal.width | sepal.width | variety |
+---+-------------+-------------+---------+
| 0 | 3.5         | 3.5         | Setosa  |
| 1 | 3           | 3           | Setosa  |
| 2 | 3.2         | 3.2         | Setosa  |
| 3 | 3.1         | 3.1         | Setosa  |
| 4 | 3.6         | 3.6         | Setosa  |
+---+-------------+-------------+---------+

Where "variety" == "Setosa":
+---+--------------+-------------+--------------+-------------+---------+
|   | sepal.length | sepal.width | petal.length | petal.width | variety |
+---+--------------+-------------+--------------+-------------+---------+
| 0 | 5.1          | 3.5         | 1.4          | 0.2         | Setosa  |
| 1 | 4.9          | 3           | 1.4          | 0.2         | Setosa  |
| 2 | 4.7          | 3.2         | 1.3          | 0.2         | Setosa  |
| 3 | 4.6          | 3.1         | 1.5          | 0.2         | Setosa  |
| 4 | 5            | 3.6         | 1.4          | 0.2         | Setosa  |
+---+--------------+-------------+--------------+-------------+---------+

Where "sepal.length" > 5:
+----+--------------+-------------+--------------+-------------+---------+
|    | sepal.length | sepal.width | petal.length | petal.width | variety |
+----+--------------+-------------+--------------+-------------+---------+
| 0  | 5.1          | 3.5         | 1.4          | 0.2         | Setosa  |
| 5  | 5.4          | 3.9         | 1.7          | 0.4         | Setosa  |
| 10 | 5.4          | 3.7         | 1.5          | 0.2         | Setosa  |
| 14 | 5.8          | 4           | 1.2          | 0.2         | Setosa  |
| 15 | 5.7          | 4.4         | 1.5          | 0.4         | Setosa  |
+----+--------------+-------------+--------------+-------------+---------+

Where "variety" != "Setosa":
+----+--------------+-------------+--------------+-------------+------------+
|    | sepal.length | sepal.width | petal.length | petal.width |  variety   |
+----+--------------+-------------+--------------+-------------+------------+
| 50 | 7            | 3.2         | 4.7          | 1.4         | Versicolor |
| 51 | 6.4          | 3.2         | 4.5          | 1.5         | Versicolor |
| 52 | 6.9          | 3.1         | 4.9          | 1.5         | Versicolor |
| 53 | 5.5          | 2.3         | 4            | 1.3         | Versicolor |
| 54 | 6.5          | 2.8         | 4.6          | 1.5         | Versicolor |
+----+--------------+-------------+--------------+-------------+------------+

Where "sepal.width" > 4 && "variety" != "Versicolor":
+----+--------------+-------------+--------------+-------------+---------+
|    | sepal.length | sepal.width | petal.length | petal.width | variety |
+----+--------------+-------------+--------------+-------------+---------+
| 15 | 5.7          | 4.4         | 1.5          | 0.4         | Setosa  |
| 32 | 5.2          | 4.1         | 1.5          | 0.1         | Setosa  |
| 33 | 5.5          | 4.2         | 1.4          | 0.2         | Setosa  |
+----+--------------+-------------+--------------+-------------+---------+

Drop "sepal.width", "sepal.length" columns:
+---+--------------+-------------+---------+
|   | petal.length | petal.width | variety |
+---+--------------+-------------+---------+
| 0 | 1.4          | 0.2         | Setosa  |
| 1 | 1.4          | 0.2         | Setosa  |
| 2 | 1.3          | 0.2         | Setosa  |
| 3 | 1.5          | 0.2         | Setosa  |
| 4 | 1.4          | 0.2         | Setosa  |
+---+--------------+-------------+---------+

Insert new column called "petal.length2" where "petal.length" + 100
+---+--------------+-------------+---------+---------------+
|   | petal.length | petal.width | variety | petal.length2 |
+---+--------------+-------------+---------+---------------+
| 0 | 1.4          | 0.2         | Setosa  | 101.4         |
| 1 | 1.4          | 0.2         | Setosa  | 101.4         |
| 2 | 1.3          | 0.2         | Setosa  | 101.3         |
| 3 | 1.5          | 0.2         | Setosa  | 101.5         |
| 4 | 1.4          | 0.2         | Setosa  | 101.4         |
+---+--------------+-------------+---------+---------------+

Insert new column called "variety2" where "variety" column + " Iris"
+---+--------------+-------------+---------+---------------+-------------+
|   | petal.length | petal.width | variety | petal.length2 |  variety2   |
+---+--------------+-------------+---------+---------------+-------------+
| 0 | 1.4          | 0.2         | Setosa  | 101.4         | Setosa Iris |
| 1 | 1.4          | 0.2         | Setosa  | 101.4         | Setosa Iris |
| 2 | 1.3          | 0.2         | Setosa  | 101.3         | Setosa Iris |
| 3 | 1.5          | 0.2         | Setosa  | 101.5         | Setosa Iris |
| 4 | 1.4          | 0.2         | Setosa  | 101.4         | Setosa Iris |
+---+--------------+-------------+---------+---------------+-------------+
```

## How to contribute?
- Give this repo a star.
- Create tutorials about Data Engineering with Go.
- Use this library and if you have some issues please put it on issues section with the Data.
- If you need a specific feature, please create a PR to README.md to request it.

## What does higor know how to do today?
- ReadCSV (With custom NaN, custom Separator)
- Print DataFrames (completly, head, tail)
- Column manipulations (Select, Drop)

## Next features:
- Implement GroupBy functions (count, max, mean, median, min, sum)
- Implement concurrency on internal functions
