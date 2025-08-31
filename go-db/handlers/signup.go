package signup

import (
	"encoding/json"
	"fmt"
	"go-db/database"
	"log"
	"net/http"
)

type customerstruct struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
	Fbconcern bool
	Xconcern  bool
}

func checkUserExist(email string) bool {
	dbconn, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("error on connecting DB")
	}
	var exist int
	dbconn.QueryRow("select 1 from customers where email = $1", email).Scan(&exist)
	return exist != 0
}

func AddUser(newuser customerstruct) error {
	sqlStatement := fmt.Sprintf("insert into customers (firstname,lastname,email,password,fbconcern,xconcern) values('%s','%s','%s','%s',%t,%t)",
		newuser.Firstname, newuser.Lastname, newuser.Email, newuser.Password, newuser.Fbconcern, newuser.Xconcern)
	dbconn, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("error on connecting DB")
	}
	rows, err := dbconn.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("error on querying row")
	}
	rowsAffected, err := rows.RowsAffected()
	fmt.Printf("Rows affected: %d\n", rowsAffected)
	return err
}

func SignupUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		defer r.Body.Close()
		var data customerstruct
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		isUserExist := checkUserExist(data.Email)
		fmt.Println("isUserExist", isUserExist)

		if isUserExist {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("email already exist"))
			return
		}
		err = AddUser(data)
		if err != nil {
			http.Error(w, "Failed to Add User", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("user added successfully"))
	} else {
		w.Header().Set("Allow", "POST")
		http.Error(w, "invalid method", http.StatusBadRequest)
	}
}

func Fbconcern(w http.ResponseWriter, r *http.Request) {

}

func SigninUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		defer r.Body.Close()
		var data customerstruct
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		isUserExist := checkUserExist(data.Email)
		if isUserExist {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("user exist"))
			return
		} else {
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("user not successfully"))
		}
	} else {
		w.Header().Set("Allow", "GET")
		http.Error(w, "invalid method", http.StatusBadRequest)
	}
}
