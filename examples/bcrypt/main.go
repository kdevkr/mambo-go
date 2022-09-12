package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

const BcryptCost = 12

func main() {
	plainPass := []byte("mambo")
	log.Println("Plain:", string(plainPass))

	hashedPass, err := bcrypt.GenerateFromPassword(plainPass, BcryptCost)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Hashed:", string(hashedPass))

	err = bcrypt.CompareHashAndPassword(hashedPass, plainPass)
	if err == nil {
		log.Println("Password Equals")
	} else {
		log.Println("Password Not Equals")
	}
}
