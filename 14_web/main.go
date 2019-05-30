package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Arnab !!! Congrats !!! This is your first GO WEB Program !!! ")
}

func indexTwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Arnab !!! Congrats !!! You have reached Page 2 now !!! ")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/page2", indexTwo)
	fmt.Println("Server Starting.....")
	http.ListenAndServe(":3000", nil)
	fmt.Println("Server Started !!!")
}
