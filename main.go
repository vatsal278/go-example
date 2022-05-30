package main

import (
	"fmt"
	"gobasics/controller"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func main() {
	controller.Books = []controller.Book{
		{Title: "Book 1", Author: "Author 1"},
		{Title: "Book 2", Author: "Author 2"},
	}
	log.Printf("%t", controller.Books)
	log.Print("started server")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/allbooks", controller.AllBooks)

	err := http.ListenAndServe(":8080", nil)
	log.Print("started server")

	if err != nil {
		fmt.Print(err.Error())
	}
}
