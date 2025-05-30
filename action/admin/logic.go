package admin

import (
	"database/sql"
	"ecommerce-backend-go/db"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func ValidateAdminCredentials(usernameOrEmail, password string) error {
	var hashedPwd string
	var userType int
	err := db.DB.QueryRow(`
		SELECT password, type FROM users 
		WHERE (username = ? OR email = ?) AND is_active = 1
	`, usernameOrEmail, usernameOrEmail).Scan(&hashedPwd, &userType)

	if err == sql.ErrNoRows {
		return errors.New("invalid username/email or inactive account")
	} else if err != nil {
		return errors.New("internal server error")
	}

	if userType != 1 { // 1 = Admin
		return errors.New("access denied: not an admin user")
	}

	if bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password)) != nil {
		return errors.New("incorrect password")
	}

	return nil
}
