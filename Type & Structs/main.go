package main

import (
	"log"
	"time"
)

// Type & Structs

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}
//using Capital writing variable or function it makes a public 

func main() {
	user := User{
		FirstName: "Agustian",
		LastName:  "Suhartono",
		PhoneNumber: "+62 85643 480 436",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}


