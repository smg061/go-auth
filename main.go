package main

import (
	"go-auth/hmac"
	"log"
)

func main() {
	b := []byte("hello world")

	signed, err := hmac.SignMessage(b)

	if err != nil {
		log.Fatal(err)
	}
	equals, err := hmac.CheckSignature([]byte("hello world"), signed)

	if err != nil {
		log.Fatal(err)
	}
	if !equals {
		log.Fatal("message not equal", err)
	}
	log.Printf("message equal")
}
