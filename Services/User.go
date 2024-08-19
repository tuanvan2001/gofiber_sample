package Services

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

// Verify kiểm tra JWT và trả về thông tin userID nếu hợp lệ
func Verify(tokenString string) (string, error) {
	// Phân tích cú pháp token và xác minh chữ ký
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Kiểm tra xem phương thức ký có đúng không
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("phương thức ký không hợp lệ: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return "", err
	}

	// Lấy các claims nếu token hợp lệ
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := claims["userID"].(string)
		return userID, nil
	}

	return "", fmt.Errorf("token không hợp lệ")
}
