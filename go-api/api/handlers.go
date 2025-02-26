package api

import (
	"fmt"
	"go-api/pkg/auth"
	"go-api/pkg/matrix"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

// Maneja el proceso de inicio de sesión y generación de JWT
func Login(c *fiber.Ctx) error {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Parsear el cuerpo de la solicitud JSON
	if err := c.BodyParser(&creds); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Credenciales inválidas"})
	}

	// Validar credenciales (reemplazar con la lógica real de autenticación)
	if creds.Username != "admin" || creds.Password != "password" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Usuario o contraseña incorrectos"})
	}

	// Generar un token JWT
	token, err := auth.GenerateJWT(creds.Username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error al generar el token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

// Maneja la solicitud de factorización QR y la comunicación con la API de Node.js
func QR(c *fiber.Ctx) error {
	// Registrar el cuerpo de la solicitud para depuración
	fmt.Println("Cuerpo de la solicitud:", string(c.Body()))

	// Verificar si el encabezado de autorización está presente
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Encabezado de autorización faltante"})
	}

	var request struct {
		Matrix [][]float64 `json:"matrix"`
	}

	// Parsear la solicitud JSON
	if err := c.BodyParser(&request); err != nil {
		fmt.Println("Error al parsear el cuerpo de la solicitud:", err) // Registrar el error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Matriz inválida"})
	}

	// Registrar la matriz parseada para depuración
	fmt.Println("Matriz parseada:", request.Matrix)

	q, r := matrix.QRFactorization(request.Matrix)

	nodeAPIURL := os.Getenv("NODE_API_URL")
	if nodeAPIURL == "" {
		nodeAPIURL = "http://localhost:3000"
	}

	// Enviar Q y R a la API de Node.js
	client := fiber.Post(nodeAPIURL + "/statistics")
	client.Set("Authorization", authHeader)
	client.JSON(fiber.Map{"Q": q, "R": r})
	status, body, errs := client.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error al comunicarse con la API de Node.js"})
	}

	return c.Status(status).Send(body)
}
