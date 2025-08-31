package main

import (
	signup "go-db/handlers"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/signup", signup.SignupUser)
	http.HandleFunc("/signIn", signup.SigninUser)
	http.HandleFunc("/fbconcern", signup.Fbconcern)
	http.ListenAndServe(":5000", nil)
}
