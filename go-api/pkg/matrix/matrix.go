package matrix

import "gonum.org/v1/gonum/mat"

// QRFactorization realiza la factorización QR de una matriz dada.
// Devuelve las matrices Q y R resultantes.
func QRFactorization(matrix [][]float64) (q, r [][]float64) {
	rows := len(matrix)
	cols := len(matrix[0])
	data := make([]float64, rows*cols)

	// Convertir la matriz de entrada en una estructura compatible con Gonum
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			data[i*cols+j] = matrix[i][j]
		}
	}
	dense := mat.NewDense(rows, cols, data)
	var qr mat.QR
	qr.Factorize(dense)

	// Obtener las matrices Q y R de la factorización QR
	var Q, R mat.Dense
	qr.QTo(&Q)
	qr.RTo(&R)
	q = denseToSlice(&Q)
	r = denseToSlice(&R)
	return
}

// denseToSlice convierte una matriz densa de Gonum en una representación de slices en Go
func denseToSlice(d *mat.Dense) [][]float64 {
	rows, cols := d.Dims()
	slice := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		slice[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			slice[i][j] = d.At(i, j)
		}
	}
	return slice
}
