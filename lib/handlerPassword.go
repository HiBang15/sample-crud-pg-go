package lib

import "golang.org/x/crypto/bcrypt"

const PassCOST = 12

func GeneratePassword(password string) ([]byte, error)  {
	return bcrypt.GenerateFromPassword([]byte(password), PassCOST)
}