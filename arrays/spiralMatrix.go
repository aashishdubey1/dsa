package arrays

func SpiralMatrix(matrix [][]int) []int {

	rows := len(matrix)
	cols := len(matrix[0])

	top := 0
	bottom := rows - 1
	left := 0
	right := cols - 1

	result := make([]int, 0, rows*cols)

	for top <= bottom && left <= right {

		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++

		for row := top; row <= bottom; row++ {
			result = append(result, matrix[row][right])
		}
		right--

		if top <= bottom {
			for col := right; col >= left; col-- {
				result = append(result, matrix[bottom][col])
			}
			bottom--
		}

		if left <= right {
			for row := bottom; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++
		}
	}

	return result
}
