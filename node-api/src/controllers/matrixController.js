// Importar el servicio que contiene la lógica para calcular estadísticas de matrices
const matrixService = require("../services/matrixService");

/**
 * Controlador para calcular estadísticas a partir de las matrices Q y R.
 * 
 * @param {Object} req - Objeto de solicitud HTTP con las matrices en el cuerpo.
 * @param {Object} res - Objeto de respuesta HTTP para enviar los resultados.
 * 
 * @returns {Object} JSON con las estadísticas calculadas o un mensaje de error.
 */
const computeStatistics = (req, res) => {
    // Extraer matrices Q y R del cuerpo de la solicitud
    const { Q, R } = req.body;

    // Validar que ambas matrices estén presentes en la solicitud
    if (!Q || !R) {
        return res.status(400).json({ error: "Q and R matrices are required" });
    }

    // Llamar al servicio para calcular las estadísticas de las matrices
    const statistics = matrixService.computeStatistics(Q, R);

    // Enviar la respuesta con las estadísticas calculadas
    res.json(statistics);
};

// Exportar el controlador para ser usado en las rutas de la API
module.exports = {
    computeStatistics,
};