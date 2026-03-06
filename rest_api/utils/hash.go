package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(passwods string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwods), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
