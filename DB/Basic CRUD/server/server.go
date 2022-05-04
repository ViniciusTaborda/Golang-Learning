package server

import (
	"crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		return
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
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Failed to get user ID."))
		return
	}

	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Failed to connect with database."))
		return
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM USERS WHERE ID= ?", ID)
	if err != nil {
		w.Write([]byte("Failed to fetch user."))
		return
	}

	var user user
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Failed to scan user."))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Failed to json enconde user."))
		return
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Failed to get user ID."))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed to get user ID."))
		return
	}

	var user user
	if err := json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Failed to create user struct."))
		return
	}

	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Failed to connect with database."))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("UPDATE USER SET NAME = ?, EMAIL = ? WHER ID = ?")
	if err != nil {
		w.Write([]byte("Failed to connect with database."))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte("Failed to update user."))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Failed to get user ID."))
		return
	}

	db, err := database.Connect()

	if err != nil {
		w.Write([]byte("Failed to connect with database."))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("DELETE FROM USERS WHERE ID = ?")
	if err != nil {
		w.Write([]byte("Failed to connect with database."))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Failed to delete user"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
