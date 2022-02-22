package main

import (
	"log"
	"sort"
)

//its like class

type User struct {
	FirstName string
	LastName  string
}

func main() {

	//map using string types of data
	myStringMap := make(map[string]string)

	myStringMap["FirstName"] = "Farhan"
	myStringMap["LastName"] = "Pramana"
	log.Println(myStringMap["FirstName"]) // Farhan
	log.Println(myStringMap["FirstName"]) // Pramana

	log.Println("------------------------------")

	//map using int types of data
	myIntMap := make(map[string]int)
	myIntMap["First"] = 1
	myIntMap["Second"] = 2
	myIntMap["Third"] = 3
	log.Println(myIntMap["First"])
	log.Println(myIntMap["Second"])

	log.Println("------------------------------")

	//map using struct
	myMap := make(map[string]User)
	me := User{
		FirstName: "Bernie",
		LastName:  "Limaun",
	}
	myMap["me"] = me
	log.Println(myMap["me"].FirstName)
	log.Println(myMap["me"].LastName)
	log.Println("------------------------------")

	//slice
	var mySlice []int
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)
	sort.Ints(mySlice) //to sort the slice
	log.Println(mySlice)

	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(number)    // print 1 2 3 4 5 6 7 8 9 10
	log.Print(number[0:2]) //print 1 and 2

	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}
