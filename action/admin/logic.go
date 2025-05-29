package admin

import (
	"database/sql"
	"ecommerce-backend-go/db"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// ValidateAdminCredentials checks credentials against DB and returns error if invalid.
func ValidateAdminCredentials(usernameOrEmail, password string) error {
	var hashedPwd string
	err := db.DB.QueryRow(`
		SELECT password 
		FROM admin_users 
		WHERE (username = ? OR email = ?) AND is_active = 1
	`, usernameOrEmail, usernameOrEmail).Scan(&hashedPwd)

	if err == sql.ErrNoRows {
		return errors.New("invalid username/email or inactive account")
	} else if err != nil {
		return errors.New("internal server error")
	}

	if bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password)) != nil {
		return errors.New("incorrect password")
	}

	return nil
}
