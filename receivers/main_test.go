package main

import "testing"

func TestMain(t *testing.T) {
	expectedFirstName := "John"
	myVar := myStruct{
		FirstName: expectedFirstName,
	}

	actualFirstName := myVar.printFirstName()
	if actualFirstName != expectedFirstName {
		t.Errorf("want %v, got %v", expectedFirstName, actualFirstName)
	}
}
