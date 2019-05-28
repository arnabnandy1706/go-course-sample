package main

import "fmt"

func main() {
	x := 10
	y := 10

	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is lesser than %d\n", y, x)
	}

	colour := "red"

	if colour == "red" {
		fmt.Println("Colour is RED !!!")
	} else if colour == "blue" {
		fmt.Println("Colour is BLUE !!!")
	} else {
		fmt.Println("Colour is neither RED nor BLUE !!!")
	}

	animal := "cat"
	// animal := fmt.Scan("Enter the animal")
	switch animal {
	case "cat":
		fmt.Println("The animal is a CAT")
	case "dog":
		fmt.Println("The animal is a DOG")
	default:
		fmt.Println("The animal is neither CAT nor DOG")
	}

}
