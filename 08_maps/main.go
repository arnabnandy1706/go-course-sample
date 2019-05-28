package main

import "fmt"

func main() {
	emails := make(map[string]string)

	emails["Arnab"] = "arnabnandy@email.com"
	emails["Sonia"] = "soniadas@email.com"
	fmt.Println(emails)
	fmt.Println(emails["Sonia"])

	desg := map[string]string{"Arnab": "Consultant", "Sonia": "Senior Analyst"}
	fmt.Print(desg)
}
