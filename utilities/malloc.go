package utilities

func make2D[T any](n, m int) [][]T {
	matrix := make([][]T, n)
	rows := make([]T, n*m)
	for i := 0; i < n; i++ {
		matrix[i] = rows[i*m : (i+1)*m]
	}
	return matrix
}
