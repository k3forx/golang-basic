package main

import "log"

func main() {
	myString := "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString) // Can get pointer of "myString" with "&"
	log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to ", *s)
	newValue := "Red"
	*s = newValue
}
