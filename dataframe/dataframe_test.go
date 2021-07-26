package dataframe

import "testing"

type book1 struct {
	name string
}

type book2 struct {
	name string
}

type book3 struct {
	noName string
}

type bookComplete struct {
	colString   PageString
	colInt      PageInt
	colFloat64  PageInt
	colBool     PageBool
	colDatetime PageDatetime
}

func TestIsEqualBookEqual(t *testing.T) {
	internalName := "Higor"

	book1Example := book1{
		name: internalName,
	}

	book2Example := book2{
		name: internalName,
	}

	book3Example := book3{
		noName: internalName,
	}

	bookComparation := isEqualBook(book1Example, book2Example)

	if !bookComparation {
		t.Errorf("Error, both DataFrame are different but equal expected. %+v vs %+v", book1Example, book2Example)
	}

	bookComparation = isEqualBook(book1Example, book3Example)

	if bookComparation {
		t.Errorf("Error, both DataFrame are equal but different expected. %+v vs %+v", book1Example, book3Example)
	}
}

func TestBookGenerator(t *testing.T) {
	//	columns = []string{"colString", "colInt", "colFloat64", "colBool", "colDatetime"}
}
