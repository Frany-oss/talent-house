package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTSecret es la clave secreta utilizada para firmar y validar los tokens JWT
var JWTSecret = []byte("your-secret-key")

// GenerateJWT genera un token JWT con un tiempo de expiración de 24 horas
func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas
	})
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		fmt.Println("Error al firmar el token:", err) // Registrar el error
		return "", err
	}
	return tokenString, nil
}

// ValidateJWT valida un token JWT y devuelve el objeto token si es válido
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
}
