package jwt

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
)

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

type signKey struct {
	key     []byte
	created time.Time
}

var currentKey = uuid.NewV4().String()

var keys = map[string]signKey{
	currentKey: {
		[]byte("abcdefghijklmnopqrstuvwxyz"),
		time.Now(),
	},
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token has expired")
	}
	if u.SessionID == 0 {
		return fmt.Errorf("invalid sesion id")
	}
	return nil
}
func CreateToken(c *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	signedToken, err := t.SignedString(keys[currentKey])
	if err != nil {
		return "", fmt.Errorf("error in createToken when signing toke %w", err)
	}
	return signedToken, nil
}

func ParseToken(token string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(token, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("invalid signing algorithm")
		}
		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid key ID")
		}
		k, ok := keys[kid]
		fmt.Println(k)
		return keys[currentKey], nil
	})
	if err != nil {
		return nil, fmt.Errorf("error in parseToken while parsing token %w", err)
	}
	if !t.Valid {
		return nil, fmt.Errorf("error parsing token; token not valid")
	}
	claims, _ := t.Claims.(*UserClaims)

	return claims, nil
}

func GenerateNewKey() error {
	newKey := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, newKey)
	if err != nil {
		return fmt.Errorf("error generating new key: %w", err)
	}
	id := uuid.NewV4().String()
	keys[id] = signKey{
		key:     newKey,
		created: time.Now(),
	}
	currentKey = id
	return nil
}
