// Función para calcular estadísticas a partir de dos matrices Q y R
const computeStatistics = (Q, R) => {
    // Aplanar ambas matrices y combinarlas en un solo array
    const allValues = [...Q.flat(), ...R.flat()];

    // Calcular el valor máximo en las matrices
    const max = Math.max(...allValues);
    // Calcular el valor mínimo en las matrices
    const min = Math.min(...allValues);
    // Calcular la suma de todos los valores en las matrices
    const sum = allValues.reduce((acc, val) => acc + val, 0);
    // Calcular el promedio de los valores
    const avg = sum / allValues.length;

    // Función auxiliar para verificar si una matriz es diagonal
    const isDiagonal = (matrix) => {
        for (let i = 0; i < matrix.length; i++) {
            for (let j = 0; j < matrix[i].length; j++) {
                // Si el elemento no está en la diagonal principal y no es 0, la matriz no es diagonal
                if (i !== j && matrix[i][j] !== 0) return false;
            }
        }
        return true; // Retorna verdadero si la matriz es diagonal
    };

    // Aplicar la función de verificación a ambas matrices Q y R
    const diagonalCheck = [Q, R].map(isDiagonal);

    // Retornar las estadísticas calculadas
    return { max, min, avg, sum, diagonalCheck };
};

// Exportar la función para su uso en otros módulos
module.exports = {
    computeStatistics,
};