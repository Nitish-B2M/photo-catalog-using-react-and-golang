package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashData(data []byte) (string, error) {
	hashedData, err := bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing data:", err)
		return "", err
	}
	return string(hashedData), nil
}

func VerifyData(hashedData string, data []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedData), data)
	if err != nil {
		log.Println("Data verification failed:", err)
		return false
	}
	return true
}
