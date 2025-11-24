package handlers

import (
	"encoding/json"
	"first-rest-api/db"
	"first-rest-api/models"
	"first-rest-api/response"
	"first-rest-api/security"
	"first-rest-api/validator"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var u models.UserModel
	json.NewDecoder(r.Body).Decode(&u)

	if u.Email == "" || u.Password == "" {
		response.JSON(w, http.StatusBadRequest, false, "Email and password are required", nil)
		return
	}

	if err := validator.ValidateEmail(u.Email); err != nil {
		response.JSON(w, 400, false, err.Error(), nil)
		return
	}

	if err := validator.ValidatePassword(u.Password); err != nil {
		response.JSON(w, 400, false, err.Error(), nil)
		return
	}

	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email=$1)", u.Email).Scan(&exists)
	if err != nil {
		response.JSON(w, 500, false, "Error checking existing user", nil)
		return
	}

	if exists {
		response.JSON(w, 400, false, "User already exists", nil)
		return
	}

	hashed, err := security.HashPassword(u.Password)
	if err != nil {
		response.JSON(w, 500, false, "Could not hash password", nil)
		return
	}

	_, err = db.DB.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", u.Email, hashed)
	if err != nil {
		response.JSON(w, 500, false, "Error saving user", nil)
		return
	}

	response.JSON(w, 201, true, "Registered successfully", nil)
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
