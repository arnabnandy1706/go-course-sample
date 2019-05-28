package main

import "fmt"

func main() {
	// Arrays
	var fruitArray [2]string

	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	fmt.Println(fruitArray)

	carArray := [2]string{"Lamborghini", "Ferrari"}

	fmt.Println(carArray)

	// Slices
	countrySlices := []string{"India", "USA", "China"}

	fmt.Println(countrySlices)
}
