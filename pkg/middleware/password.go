package middleware

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreatePasswordHash(password string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}
