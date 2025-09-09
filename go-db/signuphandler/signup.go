package signup

import (
	"encoding/json"
	"fmt"
	"go-db/database"
	"log"
	"net/http"
)

type Customerstruct struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
	Fbconcern bool
	Xconcern  bool
}

func CheckUserExist(email string) bool {
	var exist int
	database.DB.QueryRow("select 1 from customers where email = $1", email).Scan(&exist)
	return exist != 0
}

func AddUser(newuser Customerstruct) error {
	sqlStatement := fmt.Sprintf("insert into customers (firstname,lastname,email,password,fbconcern,xconcern) values('%s','%s','%s','%s',%t,%t)",
		newuser.Firstname, newuser.Lastname, newuser.Email, newuser.Password, newuser.Fbconcern, newuser.Xconcern)
	rows, err := database.DB.Exec(sqlStatement)
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
		var data Customerstruct
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		isUserExist := CheckUserExist(data.Email)
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
