package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSalt(text []byte) string {
	hash, err := bcrypt.GenerateFromPassword(text, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func ComparePassword(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
