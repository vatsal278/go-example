package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
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

type userdetail struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,alphanum"`
}

func ValidCredentials(w http.ResponseWriter, r *http.Request) {
	var userDetails userdetail
	err := json.NewDecoder(r.Body).Decode(&userDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//var validate *validator.Validate
	validate := validator.New()
	errs := validate.Struct(userDetails)
	if errs != nil {
		fmt.Println(errs)
		return
	} else {
		json.NewEncoder(w).Encode("Bearer token is : E8hxQS4FGHiB0qV0ShW__zqaScbTdyK18Kda8Lsu39K4mlP6EbvumaYqgFCDLMrepGuSypcf1O01P-o8m7bz1Q")
	}

}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie := r.Cookies()

		ck := fmt.Sprint(cookie)
		ckk := strings.Split(ck, "Cookie_1=")
		ckkk := strings.Split(ckk[1], "]")

		if ckkk[0] == "E8hxQS4FGHiB0qV0ShW__zqaScbTdyK18Kda8Lsu39K4mlP6EbvumaYqgFCDLMrepGuSypcf1O01P-o8m7bz1Q" {
			fmt.Print("token matched")
			next.ServeHTTP(w, r)
		} else {
			log.Fatal()
		}
		//ckk := strings.TrimPrefix(ck, "Cookie_1=")
		//fmt.Print(ckk)

	})
}
