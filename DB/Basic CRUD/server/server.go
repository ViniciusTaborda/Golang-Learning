package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed to read request body."))
		return
	}
	var user user

	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Failed to create user struct."))
	}

	fmt.Println(user)

	db, err := database.Connect()
	defer db.Close()

	if err != nil {
		w.Write([]byte("Failed to connect with database."))
		return
	}

	statement, err := db.Prepare("INSERT INTO USERS (name, email) VALUES (?, ?)")

	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Failed to execute statement."))
		return
	}

	idInserted, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("Failed to get inserted ID."))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("User inserted successfully. ID: %d", idInserted)))

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Failed to connect with database."))
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM USERS")
	if err != nil {
		w.Write([]byte("Failed to get users."))
		return
	}
	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Failed to scan users."))
			return
		}
		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Failed to encode users to json."))
		return
	}

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}
