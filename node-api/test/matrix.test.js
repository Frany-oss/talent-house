const request = require("supertest");
const app = require("../app"); // Importa la aplicación Express
const jwt = require("jsonwebtoken");
const config = require("../src/config/config");

describe("POST /statistics", () => {
    let token;

    beforeAll(() => {
        // Genera un token JWT válido para las pruebas
        token = jwt.sign({ username: "admin" }, config.JWT_SECRET, { expiresIn: "1h" });
    });

    // Prueba: solicitud válida con token
    it("should return statistics for a valid request with token", async () => {
        const response = await request(app)
            .post("/statistics") // Realiza una solicitud POST a /statistics
            .set("Authorization", `Bearer ${token}`) // Configura el encabezado con el token JWT
            .send({
                Q: [[1, 0], [0, 1]], // Matriz Q de entrada
                R: [[2, 0], [0, 2]], // Matriz R de entrada
            });

        // Verifica que la respuesta sea exitosa (200 OK)
        expect(response.status).toBe(200);
        // Verifica que el cuerpo de la respuesta contenga las propiedades esperadas
        expect(response.body).toHaveProperty("max");
        expect(response.body).toHaveProperty("min");
        expect(response.body).toHaveProperty("avg");
        expect(response.body).toHaveProperty("sum");
        expect(response.body).toHaveProperty("diagonalCheck");
    });

    // Prueba: solicitud sin token de autenticación
    it("should return 401 for a request without a token", async () => {
        const response = await request(app)
            .post("/statistics")
            .send({
                Q: [[1, 0], [0, 1]],
                R: [[2, 0], [0, 2]],
            });

        // Verifica que la respuesta sea 401 (No autorizado)
        expect(response.status).toBe(401);
        expect(response.body).toHaveProperty("error");
    });

    // Prueba: solicitud con un payload inválido
    it("should return 400 for an invalid request payload", async () => {
        const response = await request(app)
            .post("/statistics")
            .set("Authorization", `Bearer ${token}`)
            .send({}); // Enviando un payload vacío (inválido)

        // Verifica que la respuesta sea 400 (Solicitud incorrecta)
        expect(response.status).toBe(400);
        expect(response.body).toHaveProperty("error");
    });
});