package token

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

var SECRET_KEY = "KagdD1153ASD2" // никогда так не делать!

func New(userId int) (string, error) {
	const op = "token.New()"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
	})

	signedToken, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", fmt.Errorf("op: %s, err: %w", op, err)
	}
	return signedToken, nil
}
