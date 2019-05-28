package main

import "fmt"

func main() {
	var name string = "Arnab"

	age := 27

	var size float32 = 233.3

	fmt.Println(name, age, size)
	fmt.Printf("%T\n", name)

	deptName, empID, emailID := "IT", 75125, "arnab@golang.com"

	fmt.Println(deptName, empID, emailID)
}
