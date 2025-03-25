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

func Compare(tokenString string) (int, error) {
	const op = "tokem.Compare()"

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return 0, fmt.Errorf("op: %s, err: %w", op, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("op: %s, invalid claims", op)
	}

	userIdFloat, exists := claims["user_id"]
	if !exists {
		return 0, fmt.Errorf("op: %s, user_id was not found in claims", op)
	}

	userId, ok := userIdFloat.(float64)
	if !ok {
		return 0, fmt.Errorf("op: %s, user_id claim has invalid type", op)
	}

	return int(userId), nil
}
