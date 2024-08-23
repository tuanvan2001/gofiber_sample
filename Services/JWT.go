package Services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"time"
)

type Claims struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateTokenJWT(userID int, username string) (string, error) {
	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	expire, _ := strconv.Atoi(os.Getenv("JWT_EXPIRE_HOUR"))
	expirationTime := time.Now().Add(time.Duration(expire) * time.Hour)

	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Tạo token với phương thức ký HMAC và claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ký token và trả về chuỗi token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateTokenJWT(tokenString string) (*Claims, error) {
	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("token chữ ký không hợp lệ.")
		}
		return nil, fmt.Errorf("không thể phân tích cú pháp token: %v.", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("token không hợp lệ.")
	}

	return claims, nil
}
