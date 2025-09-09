package signin

import (
	"encoding/json"
	"fmt"
	"go-db/database"
	signup "go-db/signuphandler"
	"net/http"
)

func userCheck(email string, password string) signup.Customerstruct {
	var user signup.Customerstruct
	fmt.Print(email)
	fmt.Print(password)
	database.DB.QueryRow("select id from customers where email = $1", email).Scan(&user.Email, &user.Firstname, &user.Lastname)
	fmt.Print(user.Email)
	return user

}
func SigninUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		defer r.Body.Close()
		var data signup.Customerstruct
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		isUserExist := userCheck(data.Email, data.Password)
		fmt.Print(isUserExist.Email)
	} else {
		w.Header().Set("Allow", "GET")
		http.Error(w, "invalid method", http.StatusBadRequest)
	}
}
