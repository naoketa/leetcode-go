package medium062

func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for i := range make([]int, m) {
		matrix[i] = make([]int, n)
		matrix[i][0] = 1
		for j := range make([]int, n) {
			matrix[0][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}
	return matrix[m-1][n-1]
}
