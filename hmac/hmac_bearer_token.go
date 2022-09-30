package hmac

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
)


func SignMessage(message []byte) ([]byte, error) {
	var key []byte = []byte("abcdefghijklmnopqrstuvwxyz")
	hmac := hmac.New(sha512.New, key)
	_, err := hmac.Write(message)
	if err != nil {
		return nil, fmt.Errorf("error while signing message: %v", err)
	}
	signature := hmac.Sum(nil)
	return signature, nil
}

func CheckSignature(message, signature []byte) (bool, error) {
	newSig, err := SignMessage(message)
	if err != nil {
		return false, err
	}
	same := hmac.Equal(newSig, signature)
	return same, nil
}
