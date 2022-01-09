package util

import "golang.org/x/crypto/bcrypt"

//HashPassword  is to encrypt the password of the user
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

//CheckPasswordHash is to compare and bcrypt the password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
