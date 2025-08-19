package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	var todolist []todo
	handler1 := func(w http.ResponseWriter, r *http.Request) {
		resposne, err := http.Get("https://jsonplaceholder.typicode.com/todos")
		if err != nil {
			fmt.Print(err)
		}
		responseBody, err := io.ReadAll(resposne.Body)
		if err != nil {
			fmt.Print(err)
		}
		err = json.Unmarshal(responseBody, &todolist)
		if err != nil {
			fmt.Print(err)
		}
		for _, value := range todolist {
			fmt.Println("id : ", value.Id)
			fmt.Println("userId : ", value.UserId)
			fmt.Println("task : ", value.Title)
			fmt.Println("status : ", value.Completed)
		}
		w.Write([]byte("Hi"))
	}
	http.HandleFunc("/", handler1)
	http.ListenAndServe(":5000", nil)
}
