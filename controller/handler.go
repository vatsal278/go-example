package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
}

var Books []Book

func AllBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnAllArticles")
	fmt.Printf("%+v", Books)
	json.NewEncoder(w).Encode(Books)
}

func AllAuthors(w http.ResponseWriter, r *http.Request) {
	Books["Book1"]
}
