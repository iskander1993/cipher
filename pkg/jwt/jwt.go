package jwt

import (
	"time"

	golangjwt "github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("supersecretKey") // лучше вынести в ENV

// Claims — структура, которую мы зашифруем в токене
type Claims struct {
	UserID int `json:"user_id"`
	golangjwt.RegisteredClaims
}

// GenerateToken генерирует JWT токен
func GenerateToken(userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: golangjwt.RegisteredClaims{
			ExpiresAt: golangjwt.NewNumericDate(expirationTime),
		},
	}

	token := golangjwt.NewWithClaims(golangjwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ParseToken проверяет токен и возвращает userID
func ParseToken(tokenStr string) (int, error) {
	claims := &Claims{}
	token, err := golangjwt.ParseWithClaims(tokenStr, claims, func(token *golangjwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return 0, err
	}
	return claims.UserID, nil
}

//// Простейший вариант для тестов
//func ParseToken(token string) (int, error) {
//	if token == "validation" {
//		return 1, nil //Возвращаем любой userID и nil ошибку для нужного токена
//	}
//	return 0, errors.New("invalid token")
//}
