package main

import (
	"log"

	"github.com/k3forx/golang-basic/helpers"
)

func main() {
	log.Println("Hello World")

	var myVar helpers.SomeType
	myVar.TypeName = "string"
	myVar.TypeNumber = 123
	log.Println(myVar)
}
