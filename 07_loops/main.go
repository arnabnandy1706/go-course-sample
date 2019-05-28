package main

import "fmt"

func main() {
	// i := 1

	// for i <= 10 {
	// 	fmt.Println(i)
	// 	// i = i + 1
	// 	i++
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("Number %d\n", i)
	// }

	for i := 1; i <= 100; i++ {
		rem3 := i % 3
		rem5 := i % 5

		if rem3 == 0 && rem5 == 0 {
			fmt.Printf("%d: FizzBuzz\n", i)
		} else if rem3 == 0 {
			fmt.Printf("%d: Fizz\n", i)
		} else if rem5 == 0 {
			fmt.Printf("%d: Buzz\n", i)
		} else {
			fmt.Println(i)
		}
	}
}
