package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

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
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	} else {
		cookieValue := "E8hxQS4FGHiB0qV0ShW__zqaScbTdyK18Kda8Lsu39K4mlP6EbvumaYqgFCDLMrepGuSypcf1O01P-o8m7bz1Q"
		addCookie(w, "cookie", cookieValue, 30*time.Minute)
		//json.NewEncoder(w).Encode("Bearer token is : E8hxQS4FGHiB0qV0ShW__zqaScbTdyK18Kda8Lsu39K4mlP6EbvumaYqgFCDLMrepGuSypcf1O01P-o8m7bz1Q")
	}

}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if VerifyToken(r) {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
		}

		//cookie := r.Cookies()

		//ck := fmt.Sprint(cookie)
		//ckk := strings.Split(ck, "Cookie_1=")
		//ckkk := strings.Split(ckk[1], "]")

		//if ckkk[0] == "E8hxQS4FGHiB0qV0ShW__zqaScbTdyK18Kda8Lsu39K4mlP6EbvumaYqgFCDLMrepGuSypcf1O01P-o8m7bz1Q" {
		//	fmt.Print("token matched")
		//	next.ServeHTTP(w, r)
		//} else {
		//	log.Fatal()
		//}
		//ckk := strings.TrimPrefix(ck, "Cookie_1=")
		//fmt.Print(ckk)

	})
}

func addCookie(w http.ResponseWriter, name, value string, ttl time.Duration) {
	expire := time.Now().Add(ttl)
	cookie := http.Cookie{
		Name:    name,
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
}

func VerifyToken(r *http.Request) bool {

	reqToken := r.Header.Get("Authorization")
	//fmt.Print(reqToken)
	token := strings.Split(reqToken, " ")
	if len(token) != 2 {
		return false
	}
	fmt.Print(token[1])

	if token[(len(token)-1)] == "E8hxQS4FGHiB0qV0ShW__zqaScbTdyK18Kda8Lsu39K4mlP6EbvumaYqgFCDLMrepGuSypcf1O01P-o8m7bz1Q" {
		fmt.Println("valid token")
		return true
	} else {
		fmt.Println("invalid token", http.StatusUnauthorized)
		return false
	}
}

func GetbyAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("params were:", r.URL.Query())
	authorName := r.URL.Query().Get("author")
	for _, j := range Books {
		if authorName == j.Author {
			fmt.Printf("%s ", j.Title)
			json.NewEncoder(w).Encode(j.Title)
		}
	}
}

func GetByTitle(w http.ResponseWriter, r *http.Request) {
	title, ok := r.URL.Query()["title"]
	fmt.Print(title)
	if !ok {
		fmt.Print("error fetching params ", http.StatusBadRequest)
	}
	for _, j := range Books {
		if j.Title == title[0] {
			fmt.Printf("%s ", j.Title)
			json.NewEncoder(w).Encode(j.Title)

		}
	}
}
