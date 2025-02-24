// Importar el módulo 'jsonwebtoken' para manejar JWT
const jwt = require("jsonwebtoken");
const config = require("../config/config");

/**
 * Middleware de autenticación JWT.
 * Verifica la validez del token JWT presente en la cabecera 'Authorization'.
 *
 * @param {Object} req - Objeto de solicitud HTTP.
 * @param {Object} res - Objeto de respuesta HTTP.
 * @param {Function} next - Función para continuar con la siguiente ejecución en la cadena de middleware.
 */
const authenticateJWT = (req, res, next) => {
    // Obtener la cabecera de autorización
    const authHeader = req.headers.authorization;

    // Verificar si la cabecera de autorización está presente
    if (authHeader) {
        // Extraer el token eliminando la palabra "Bearer "
        const token = authHeader.split(" ")[1];

        // Verificar la validez del token con la clave secreta
        jwt.verify(token, config.JWT_SECRET, (err, user) => {
            if (err) {
                // Si el token es inválido o ha expirado, devolver un error 403 (Prohibido)
                return res.status(403).json({ error: "Invalid or expired token" });
            }

            // Adjuntar la información decodificada del usuario a la solicitud
            req.user = user;
            // Continuar con la siguiente función en la cadena de middleware
            next();
        });
    } else {
        // Si no hay cabecera de autorización, devolver un error 401 (No autorizado)
        res.status(401).json({ error: "Authorization header missing" });
    }
};

// Exportar el middleware para su uso en otras partes de la aplicación
module.exports = {
    authenticateJWT,
};