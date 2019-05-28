package main

import (
	"fmt"
	"math"

	"github.com/arnabnandy1706/go_test_project_1/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Floor(2.789))
	fmt.Println(math.Ceil(2.1971))
	fmt.Println(math.Sqrt(144))
	fmt.Println(strutil.Reverse("olleh"))
}
