package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Payload struct {
	Email   string
	Isadmin bool
	Issuadmin bool
	jwt.StandardClaims
}

func GenerateJwt(email string, isadmin,issuadmin bool, secret []byte) (string, error) {

	expiresat := time.Now().Add(48 * time.Hour)

	jwtclaims := &Payload{
		Email:   email,
		Isadmin: isadmin,
		Issuadmin: issuadmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresat.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtclaims)

	tokenstring, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return tokenstring, nil

}

func ValidateToken(tokenstring string, secret []byte) (map[string]interface{}, error) {

	token, err := jwt.ParseWithClaims(tokenstring, &Payload{}, func(t *jwt.Token) (interface{}, error) {

		if t.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("invalid token")
		}

		return secret, nil

	})

	if err != nil {
		return nil, err
	}

	if token == nil || !token.Valid {
		return nil, fmt.Errorf("token is not valid or its empty")
	}

	cliams, ok := token.Claims.(*Payload)

	if !ok {
		return nil, fmt.Errorf("cannot parse claims")
	}

	cred := map[string]interface{}{
		"email":   cliams.Email,
		"isadmin": cliams.Isadmin,
		"issuadmin": cliams.Issuadmin,
	}

	if cliams.ExpiresAt < time.Now().Unix() {
		return nil, fmt.Errorf("token expired")
	}

	return cred, nil

}
