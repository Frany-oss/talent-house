package middleware

import (
	"go-api/pkg/auth"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// JWT devuelve un middleware de Fiber para la autenticación mediante JWT
func JWT() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: auth.JWTSecret, // Utiliza la clave secreta definida en el paquete auth
	})
}
