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

func TestIsEqualBookEqual(t *testing.T) {
	internalName := "Higor"

	book1Example := book1{
		name: internalName,
	}

	book2Example := book2{
		name: internalName,
	}

	bookComparation := IsEqualBook(book1Example, book2Example)

	if !bookComparation {
		t.Errorf("Error, both DataFrame are different but equal expected. %+v vs %+v", book1Example, book2Example)
	}
}

func TestIsEqualBookDifferent(t *testing.T) {
	internalName := "Higor"

	book1Example := book1{
		name: internalName,
	}

	book2Example := book3{
		noName: internalName,
	}

	bookComparation := IsEqualBook(book1Example, book2Example)

	if bookComparation {
		t.Errorf("Error, both DataFrame are equal but different expected. %+v vs %+v", book1Example, book2Example)
	}
}
