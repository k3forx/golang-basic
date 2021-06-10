package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Kanata",
		LastName:    "Miyahana",
		PhoneNumber: "1 555-444-666",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate: ", user.BirthDate)
}
