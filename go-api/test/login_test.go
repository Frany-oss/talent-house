package test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"go-api/api"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestLogin verifica la funcionalidad del endpoint de inicio de sesión
func TestLogin(t *testing.T) {
	// Crear una nueva aplicación Fiber
	app := fiber.New()

	// Definir la ruta de login
	app.Post("/login", api.Login)

	// Casos de prueba
	tests := []struct {
		name           string
		payload        map[string]string
		expectedStatus int
		expectedToken  bool
	}{
		{
			name: "Credenciales válidas",
			payload: map[string]string{
				"username": "admin",
				"password": "password",
			},
			expectedStatus: fiber.StatusOK,
			expectedToken:  true,
		},
		{
			name: "Credenciales inválidas",
			payload: map[string]string{
				"username": "wrong",
				"password": "wrong",
			},
			expectedStatus: fiber.StatusUnauthorized,
			expectedToken:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convertir payload a JSON
			payload, _ := json.Marshal(tt.payload)

			// Crear una nueva solicitud HTTP
			req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(payload))
			req.Header.Set("Content-Type", "application/json")

			// Ejecutar la solicitud
			resp, err := app.Test(req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Verificar el cuerpo de la respuesta
			if tt.expectedToken {
				var result map[string]string
				err := json.NewDecoder(resp.Body).Decode(&result)
				assert.NoError(t, err)
				assert.NotEmpty(t, result["token"])
			}
		})
	}
}
