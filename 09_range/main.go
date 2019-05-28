package main

import "fmt"

func main() {
	ids := []int{12, 23, 45, 765, 234, 65, 22}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0

	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum =", sum)

	desg := map[string]string{"Arnab": "Consultant", "Sonia": "Senior Analyst"}

	for k, v := range desg {
		fmt.Printf("%s: %s\n", k, v)
	}
}
