package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	SecretKey string
	Claims    jwt.MapClaims
}

func (to *Token) Generate(s *string) (err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, to.Claims)

	*s, err = t.SignedString([]byte(to.SecretKey))

	return
}

func (to *Token) Verify(tokenString string, claims *jwt.MapClaims) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(to.SecretKey), nil
	})

	if err != nil {
		return err
	}

	*claims = token.Claims.(jwt.MapClaims)

	return nil
}
