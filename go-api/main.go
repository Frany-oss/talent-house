package main

import (
	"fmt"
	"go-api/api"
	"go-api/internal/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Cargar la configuración desde el archivo correspondiente
	cfg := config.Load()

	// Inicializar la aplicación Fiber
	app := fiber.New()

	// Configurar las rutas de la API
	api.SetupRoutes(app)

	// Iniciar el servidor y escuchar en el puerto configurado
	fmt.Printf("El servidor está corriendo en el puerto %s\n", cfg.Port)
	app.Listen(":" + cfg.Port)
}
