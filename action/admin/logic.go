package admin

import (
	"database/sql"
	"ecommerce-backend-go/db"
	"ecommerce-backend-go/queries/admin"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func ValidateAdminCredentials(usernameOrEmail, password string) (string, string, error) {
	var hashedPwd, name, mobile string
	var userType int

	err := db.DB.QueryRow(admin.ValidateAdminLoginQuery, usernameOrEmail, usernameOrEmail).Scan(&hashedPwd, &userType, &name, &mobile)

	if err == sql.ErrNoRows {
		return "", "", errors.New("invalid username/email or inactive account")
	} else if err != nil {
		return "", "", errors.New("internal server error")
	}
	if userType != 1 {
		return "", "", errors.New("access denied: not an admin")
	}
	if bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(password)) != nil {
		return "", "", errors.New("incorrect password")
	}
	return name, mobile, nil
}
