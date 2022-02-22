package main

import "fmt"

type Animal interface {
	Says() string
	Type() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Cat struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Shiba",
		Breed: "Shiba dog",
	}
	myAnimalInfo(&dog)

	cat := Cat{
		Name:          "Bernie",
		Color:         "Grey",
		NumberOfTeeth: 2,
	}
	myAnimalInfo(&cat)

}

func myAnimalInfo(a Animal) {
	fmt.Println("My", a.Type(), "says", a.Says(), "has a", a.NumberOfLegs())
}

func (d *Dog) Says() string {
	return "Woof-Woof"
}

func (d *Dog) Type() string {
	return "Dog"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (c *Cat) Says() string {
	return "Miaw-Miaw"
}

func (c *Cat) Type() string {
	return "Cat"
}

func (c *Cat) NumberOfLegs() int {
	return 4
}
