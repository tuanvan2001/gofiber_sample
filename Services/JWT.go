package Services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type JWT struct {
}

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func (j *JWT) Sign(userID any, secret ...string) (string, error) {
	var key []byte
	if len(secret) > 0 {
		key = []byte(secret[0])
	} else {
		key = secretKey
	}

	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *JWT) Verify(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("phương thức ký không hợp lệ: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return map[string]interface{}{
			"userID": claims["userID"],
		}, nil
	}
	return nil, fmt.Errorf("token không hợp lệ")
}
