package main

import "log"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

func (d Dog) Says() string {
	return "woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func (g Gorilla) Says() string {
	return "grunt"
}

func (g Gorilla) NumberOfLegs() int {
	return 4
}

func PrintInfo(a Animal) {
	log.Println("This animal says", a.Says(), "has", a.NumberOfLegs(), "legs")
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepherd",
	}
	PrintInfo(dog)

	gorilla := Gorilla{
		Name:          "King Kong",
		Color:         "Black",
		NumberOfTeeth: 32,
	}
	PrintInfo(gorilla)
}
