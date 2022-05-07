package util

import "golang.org/x/crypto/bcrypt"

// EncodePassword encode plain password
func EncodePassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("failed to encode password")
	}
	return string(bytes)
}

// VerifyPassword verify if the encoded string is the hash of the given password
func VerifyPassword(password, encoded string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encoded), []byte(password))
	return err == nil
}
