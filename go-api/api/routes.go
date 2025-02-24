package api

import (
	"go-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

// Configura las rutas de la aplicación
func SetupRoutes(app *fiber.App) {
	// Rutas públicas
	app.Post("/login", Login) // Ruta para autenticación e inicio de sesión

	// Rutas protegidas mediante autenticación JWT
	app.Post("/qr", middleware.JWT(), QR) // Ruta para la factorización QR, requiere autenticación
}
