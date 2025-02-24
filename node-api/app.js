const express = require("express");
const bodyParser = require("body-parser");
const matrixRoutes = require("./src/routes/matrixRoutes");
const config = require("./src/config/config");

const app = express();

// Middleware para procesar JSON en las solicitudes
app.use(bodyParser.json());

// Rutas de la API
app.use("/", matrixRoutes);

// Exportar la aplicación para pruebas unitarias
module.exports = app;

// Iniciar el servidor solo si no está en modo de prueba
if (process.env.NODE_ENV !== "test") {
    const PORT = config.PORT || 3000;
    app.listen(PORT, () => {
        console.log(`Server is running on port ${PORT}`);
    });
}