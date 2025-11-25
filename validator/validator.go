package validator

import (
	"errors"
	"regexp"
	"strconv"
	"unicode"
)

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
)

var phoneNumberRegex = regexp.MustCompile(`^[0-9]{10}$`)
var stringRegex = regexp.MustCompile(`^[a-zA-Z]+$`)
var intRegex = regexp.MustCompile(`^[0-9]+$`)

func ValidateEmail(email string) error {
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func ValidatePhoneNumber(phoneNumber int) error {
	if phoneNumber < 10 || !phoneNumberRegex.MatchString(strconv.Itoa(phoneNumber)) {
		return errors.New("phone number must be at least 10 digits long")
	}
	return nil
}

func ValidateString(str string) error {
	if str == "" || !stringRegex.MatchString(str) {
		return errors.New("string is required and must be a valid string")
	}
	return nil
}

func ValidateInt(num int) error {
	if num == 0 || !intRegex.MatchString(strconv.Itoa(num)) {
		return errors.New("number is required and must be a valid number")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	var hasUpper bool
	var hasLower bool
	var hasDigit bool
	var hasSpecial bool

	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasDigit = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !hasDigit {
		return errors.New("password must contain at least one digit")
	}
	if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	return nil
}
