package main

import (
	signin "go-db/signinhandler"
	signup "go-db/signuphandler"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/signup", signup.SignupUser)
	http.HandleFunc("/signin", signin.SigninUser)
	http.HandleFunc("/fbconcern", signup.Fbconcern)
	http.ListenAndServe(":5000", nil)
}
