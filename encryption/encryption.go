package encryption

import (
	"encoding/base64"
	"fmt"
)

func main() {
	msg := encode("This is very fun")
	fmt.Println(msg)
	decoded, err := decode(msg)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Printf("%s", decoded)
	}
}

func encode(msg string) (string) {
	encoded := base64.URLEncoding.EncodeToString([]byte(msg))
	return encoded
}

func decode (encodedMsg string) (string, error) {
	msg, err := base64.URLEncoding.DecodeString(encodedMsg)
	if err != nil {
		return "", fmt.Errorf("could not decode string %w", err)
	}
	return string(msg), nil
}