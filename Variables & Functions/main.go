package main

import "fmt"

func main() {
	//print line
	fmt.Println("Hello Word")

	//string variable
	var whatToSay string
	whatToSay = "Hello, Cruel World"
	fmt.Println(whatToSay)

	//int variable
	var i int
	i = 7
	fmt.Println("i set to", i)

	//calling a method with new variable
	//go is possible to return 2 and more types of data
	whatWasSaid, theOtherThingThatWasSaid := saySomethingElse()
	fmt.Println(whatWasSaid, theOtherThingThatWasSaid)
	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

//go is possible 2 return 2 types of data
func saySomethingElse() (string, string) {
	return "Something", "Else"
}
