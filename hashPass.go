package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hashMain() {
	pass := "12345678"
	pwHash, err := hashPassword(pass)
	if err != nil {
		log.Fatal(err)
	}
	err = comparePassword(pass, pwHash)
	if err != nil {
		log.Fatalln("Login failed: ", err)
	}
	log.Println("Logged in!")
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("Hello world!")))
}

func hashPassword(password string) ([]byte, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while generating bcrypt hash from password: %w", err)
	}
	return b, nil
}

func comparePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	return err
}
