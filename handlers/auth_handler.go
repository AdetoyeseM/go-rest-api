package handlers

import (
	"encoding/json"
	"first-rest-api/db"
	"first-rest-api/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register(writer http.ResponseWriter, response *http.Request) {

	setJSONHeader(writer)

	var request models.RegisterRequest
	json.NewDecoder(response.Body).Decode(&request)

	
	//hash password
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	//save to DB

	query := `INSERT INTO users(email, password) VALUES ($1, $2)`

	_, error := db.DB.Exec(query, request.Email, string(hashPassword))

	if error != nil {
		http.Error(writer, "User already exists", http.StatusBadRequest)
		return

	}

	writer.WriteHeader(http.StatusCreated)

	json.NewEncoder(writer).Encode(map[string]string{"message": "User Created Successfully"})

}

func LoginUser(writer http.ResponseWriter, response *http.Request) {

	setJSONHeader(writer)

	var request models.LoginRequest
	var userModel models.UserModel
	json.NewDecoder(response.Body).Decode(&request)

	query := `SELECT id, email, password FROM users WHERE email=$1`

	error := db.DB.QueryRow(query, request.Email).Scan(&userModel.ID, &userModel.Email, &userModel.Password)

	if error != nil {
		http.Error(writer, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	//check password

	error = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(request.Password))
	if error != nil {
		http.Error(writer, "Invalid Credentials", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(writer).Encode(map[string]string{"message": "Login Successful"})

}
