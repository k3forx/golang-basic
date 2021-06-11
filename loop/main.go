package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	mySlice := []string{"dog", "cat", "horse", "fish", "banana"}

	for _, x := range mySlice {
		log.Println(x)
	}

	myMap := make(map[string]string)
	myMap["dog"] = "dog"
	myMap["fish"] = "fish"
	myMap["hat"] = "hat"

	for i, x := range myMap {
		log.Println(i, x)
	}

	var mySlice2 []User
	u1 := User{
		FirstName: "Kanata",
		LastName:  "Miyahana",
	}
	u2 := User{
		FirstName: "Tanaka",
		LastName:  "Taro",
	}
	mySlice2 = append(mySlice2, u1)
	mySlice2 = append(mySlice2, u2)

	for i, x := range mySlice2 {
		log.Println(i, x)
	}
}
