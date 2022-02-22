package main

import "log"

func main() {
	// if else case

	cat := "Cat"
	if cat == "Cat" {
		log.Println("Cat is Cat")
	} else {
		log.Println("Cat isn't Cat")
	}

	myNum := 100
	isTrue := false

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	} else if myNum < 100 && isTrue {
		log.Println("1")
	} else if myNum == 100 || isTrue {
		log.Println("2")
	} else if myNum > 1000 && isTrue == false {
		log.Println("3")
	}

	// switch case
	myAnimal := "Cat"
	myAnimal = "Rabbit"
	switch myAnimal {
	case "Cat":
		log.Println("myAnimal is set to Cat")
		break
	case "Dog":
		log.Println("myAnimal is set to Dog")
		break
	case "Fish":
		log.Println("myAnimal is set to Fish")
		break
	default:
		log.Println("myAnimal is something else")
	}

}
