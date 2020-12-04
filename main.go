package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"go-mysql-react/helpers"
	"go-mysql-react/models"
	"github.com/gorilla/mux"
)

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users Users
	var arrUser []Users
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select id, name ,email, password from users")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.Id, &users.name, &users.email, &users.password); err != nil {
			log.Fatal(err.Error())

		} else {
			arrUser = append(arrUser, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrUser

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	r := mux.NewRouter()
	r.HandleFunc("/getUsers", returnAllUsers).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	fmt.Println("Connected to port 1234")
	http.ListenAndServe(":3000", r)
}
