package medium048

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			idx := len(matrix) - 1 - j
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][idx]
			matrix[i][idx] = tmp
		}
	}
}
