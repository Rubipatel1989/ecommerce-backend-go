package validation

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"regexp"
	"strings"
)

type ProductInput struct {
	Name  string
	Price float64
	Stock int
}

func (p ProductInput) Validate() error {
	if strings.TrimSpace(p.Name) == "" {
		return errors.New("product name is required")
	}
	if p.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	if p.Stock < 0 {
		return errors.New("stock cannot be negative")
	}
	return nil
}

// IsValidEmail checks if the email is in a valid format
func IsValidEmail(email string) bool {
	email = strings.TrimSpace(email)
	// simple regex for email validation
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
func ParseJSONStringOrRaw(ns sql.NullString) interface{} {
	if !ns.Valid || ns.String == "" {
		return nil
	}
	var parsed interface{}
	if err := json.Unmarshal([]byte(ns.String), &parsed); err != nil {
		log.Printf("Warning: Failed to parse JSON in response_options: %v", err)
		return ns.String
	}
	return parsed
}
func IsMobileNumberValid(mobile string) bool {
	mobile = strings.TrimSpace(mobile)
	// simple regex for mobile number validation
	re := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	return re.MatchString(mobile)
}

// IsValidURL checks if the URL is in a valid format
func IsValidURL(url string) bool {
	url = strings.TrimSpace(url)
	// simple regex for URL validation
	re := regexp.MustCompile(`^(http|https)://[a-zA-Z0-9.-]+(?:\.[a-zA-Z]{2,})+.*$`)
	return re.MatchString(url)
}

// IsValidUsername checks if the username is in a valid format
func IsValidUsername(username string) bool {
	username = strings.TrimSpace(username)
	// simple regex for username validation
	re := regexp.MustCompile(`^[a-zA-Z0-9._-]{3,20}$`)
	return re.MatchString(username)
}

// IsValidPassword checks if the password is in a valid format
func IsValidPassword(password string) bool {
	password = strings.TrimSpace(password)
	// simple regex for password validation
	re := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$`)
	return re.MatchString(password)
}

// IsValidName checks if the name is in a valid format
func IsValidName(name string) bool {
	name = strings.TrimSpace(name)
	// simple regex for name validation
	re := regexp.MustCompile(`^[a-zA-Z\s]{2,50}$`)
	return re.MatchString(name)
}

// IsValidPhoneNumber checks if the phone number is in a valid format
func IsValidPhoneNumber(phone string) bool {
	phone = strings.TrimSpace(phone)
	// simple regex for phone number validation
	re := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	return re.MatchString(phone)
}

// IsValidAddress checks if the address is in a valid format
func IsValidAddress(address string) bool {
	address = strings.TrimSpace(address)
	// simple regex for address validation
	re := regexp.MustCompile(`^[a-zA-Z0-9\s,.'-]{5,100}$`)
	return re.MatchString(address)
}

// IsValidPostalCode checks if the postal code is in a valid format
func IsValidPostalCode(postalCode string) bool {
	postalCode = strings.TrimSpace(postalCode)
	// simple regex for postal code validation
	re := regexp.MustCompile(`^[0-9]{5}(?:-[0-9]{4})?$`)
	return re.MatchString(postalCode)
}
