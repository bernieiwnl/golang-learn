package main

import "fmt"

func main() {

	//looping an i
	fmt.Println("Looping an i")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("--------------------")

	//looping using an slice
	fruits := []string{
		"Orange",
		"Mango",
		"Banana",
	}
	fmt.Println("Looping using data from an slice fruits")
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	fmt.Println("--------------------")

	//looping using an map
	animals := make(map[int]string)
	animals[1] = "dog"
	animals[2] = "cat"
	animals[3] = "duck"

	fmt.Println("Looping using data from an map animals")
	for _, animal := range animals {
		fmt.Println(animal)
	}

	fmt.Println("--------------------")

	//looping using an type
	type User struct {
		FullName string
		Age      int
	}

	//make slice with using custom type
	users := []User{
		{"Agustian", 21},
		{"Farhan", 21},
		{"Galih", 21},
	}

	fmt.Println("Looping using data from an map animals")
	for _, dataUser := range users {
		fmt.Println(dataUser.FullName, "Berumur", dataUser.Age)
	}

}
