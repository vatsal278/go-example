package main

import (
	"fmt"
	"gobasics/controller"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Please Validate your credentials")
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware", r.URL)

		h.ServeHTTP(w, r)
	})
}

func main() {
	controller.Books = []controller.Book{
		{Title: "Book 1", Author: "Author 1"},
		{Title: "Book 2", Author: "Author 2"},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/validate", controller.ValidCredentials)

	mux.Handle("/allbooks", controller.Middleware(http.HandlerFunc(controller.AllBooks)))

	err := http.ListenAndServe(":8080", mux)
	log.Print("started server")

	if err != nil {
		fmt.Print(err.Error())
	}
}
