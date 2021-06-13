package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Persion struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
	{
		"first_name": "Clark",
		"last_name": "Kent",
		"hair_color": "black",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair_color": "black",
		"has_dog": false
	}
]`

	var unmarshalled []Persion
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json: ", err)
	}
	log.Printf("Unmarshalled: %v", unmarshalled)

	// write JSON from a struct
	var mySlice []Persion

	var marshall1 Persion
	marshall1.FirstName = "Wally"
	marshall1.LastName = "West"
	marshall1.HairColor = "red"
	marshall1.HasDog = false

	mySlice = append(mySlice, marshall1)

	var marshall2 Persion
	marshall2.FirstName = "Diana"
	marshall2.LastName = "Prince"
	marshall2.HairColor = "black"
	marshall2.HasDog = true

	mySlice = append(mySlice, marshall2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Println("Error marshalling", err)
	}
	fmt.Println(string(newJson))
}
