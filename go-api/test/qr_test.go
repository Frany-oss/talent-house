package test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"go-api/api"
	"go-api/internal/middleware"
	"go-api/pkg/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestQR prueba el endpoint /qr, verificando la autenticación y la factorización QR.
func TestQR(t *testing.T) {
	// Crear una nueva aplicación Fiber
	app := fiber.New()

	// Definir la ruta protegida con middleware JWT
	app.Post("/qr", middleware.JWT(), api.QR)

	// Generar un token JWT válido
	token, err := auth.GenerateJWT("admin")
	assert.NoError(t, err)

	// Casos de prueba
	tests := []struct {
		name           string
		payload        map[string]interface{}
		token          string
		expectedStatus int
	}{
		{
			name: "Solicitud válida con token",
			payload: map[string]interface{}{
				"matrix": [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			token:          token,
			expectedStatus: fiber.StatusOK,
		},
		{
			name: "Falta de token de autorización",
			payload: map[string]interface{}{
				"matrix": [][]float64{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			token:          "",
			expectedStatus: fiber.StatusBadRequest,
		},
	}

	// Ejecutar cada caso de prueba
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convertir la carga útil a formato JSON
			payload, _ := json.Marshal(tt.payload)

			// Crear una nueva solicitud HTTP POST
			req := httptest.NewRequest("POST", "/qr", bytes.NewBuffer(payload))
			req.Header.Set("Content-Type", "application/json")
			if tt.token != "" {
				req.Header.Set("Authorization", "Bearer "+tt.token)
			}

			// Ejecutar la solicitud
			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
		})
	}
}
