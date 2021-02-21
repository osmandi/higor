[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# Higor

Dataframe for Golang, simple but powerful.

## Why Golang for Data Engineering?

Go has a multiple features that help on this topic, for example:
- It's very easy implement the concurrency to processing a lot of data using native resources (without external libraries)
- Its compilation is fast.
- It has a native library to implement tests.
- Its sintax is simple such as Python but powerfull as C language.
- Its mascot is very cute (this is a personal opinion)

## Why Higor?

Currently, Python is used in all steps for a Machine Learning project, from Data Engineering to Machine Learning model train. However, to make this possible, Python uses precompiled code (i.e. Numpy uses C and Spark uses Scala) instead a native way. Doing this, causes syntaxis changes, depending on the external library used.

The aim of Higor is to provide a library that allows you to work with different amounts of data using Golang with having a simple interface.

> Remember, Kubernetes is made in Golang. Can you imagine that power in Data Engineering and Big Data?

## Install

```Bash
go get -v -u github.com/osmandi/higor
```

