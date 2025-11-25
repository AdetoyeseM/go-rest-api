package handlers

import (
	"encoding/json"
	"first-rest-api/db"
	"first-rest-api/jwt"
	"first-rest-api/models"
	"first-rest-api/response"
	"first-rest-api/security"
	"first-rest-api/validator"

	"log"
	"net/http"
	// "strconv"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var u models.UserModel
	json.NewDecoder(r.Body).Decode(&u)
	missing := []string{}

	if u.Email == "" {
		missing = append(missing, "email")
	}
	if u.Password == "" {
		missing = append(missing, "password")
	}
	if u.FirstName == "" {
		missing = append(missing, "first_name")
	}
	if u.LastName == "" {
		missing = append(missing, "last_name")
	}
	if u.PhoneNumber == 0 {
		missing = append(missing, "phone_number")
	}

	// If missing fields exist, log them + return JSON
	if len(missing) > 0 {
		log.Printf("❌ Missing fields in Register request: %v\n", missing)

		response.JSON(w, http.StatusBadRequest, false, "Missing required fields", map[string]interface{}{
			"missing_fields": missing,
		})
		return
	}

	// Validation
	if err := validator.ValidateString(u.FirstName); err != nil {
		response.JSON(w, 400, false, err.Error(), nil)
		return
	}

	if err := validator.ValidateString(u.LastName); err != nil {
		response.JSON(w, 400, false, err.Error(), nil)
		return
	}

	if err := validator.ValidatePhoneNumber(u.PhoneNumber); err != nil {
		response.JSON(w, 400, false, err.Error(), nil)
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

	// Check email already exists
	var emailExists, phoneExists bool
	err := db.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE email=$1)", u.Email).Scan(&emailExists)
	if err != nil {
		response.JSON(w, 500, false, "Error checking existing user", nil)
		return
	}
	phoneError := db.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE phone_number=$1)", u.PhoneNumber).Scan(&phoneExists)

	if phoneError != nil {
		response.JSON(w, 500, false, "Error checking existing user", nil)
		return
	}
	if emailExists {
		response.JSON(w, 400, false, "User with this email already exists", nil)
		return
	}
	if phoneExists {
		response.JSON(w, 400, false, "User with this phone number already exists", nil)
		return
	}

	// Encrypt password
	hashed, err := security.HashPassword(u.Password)
	if err != nil {
		response.JSON(w, 500, false, "Could not encrypt password", nil)
		return
	}

	_, err = db.DB.Exec(
		"INSERT INTO users (email, first_name, last_name, phone_number,password_hash) VALUES ($1, $2, $3, $4, $5)",
		u.Email, u.FirstName, u.LastName, u.PhoneNumber, hashed,
	)

	if err != nil {
		response.JSON(w, 500, false, "Error saving user", nil)
		return
	}

	response.JSON(w, 200, true, "Registered successfully", nil)
}

func LoginUser(writer http.ResponseWriter, r *http.Request) {

	setJSONHeader(writer)

	var request models.LoginRequest
	var userModel models.UserModel
	json.NewDecoder(r.Body).Decode(&request)

	query := `SELECT id, email, password_hash FROM users WHERE email=$1`

	error := db.DB.QueryRow(query, request.Email).Scan(&userModel.ID, &userModel.Email, &userModel.Password)

	if error != nil {
		response.JSON(writer, 401, false, "Invalid Credential", nil)
		return
	}
	error = security.CheckPassword(userModel.Password, request.Password)

	if error != nil {
		response.JSON(writer, 401, false, "Invalid Credential", nil)
		return
	}

	token, err := jwt.GenerateToken(userModel.ID, userModel.Email)

	if err != nil {
		response.JSON(writer, 401, false, "Error generating token", nil)
		return

	}

	response.JSON(writer, 200, true, "Login Successful", map[string]interface{}{
		"token": token,
		"email": userModel.Email,
	})

}
