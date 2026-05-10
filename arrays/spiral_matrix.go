package arrays


func (a *array) SpiralOrder(matrix [][]int) []int {

	m := len(matrix)
	n := len(matrix[0])

	ans := make([]int, 0, m*n)

	top := 0
	bottom := m - 1
	left := 0
	right := n - 1

	for top <= bottom && left <= right {

		// left -> right
		for j := left; j <= right; j++ {
			ans = append(ans, matrix[top][j])
		}
		top++

		// top -> bottom
		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--

		// right -> left
		if top <= bottom {
			for j := right; j >= left; j-- {
				ans = append(ans, matrix[bottom][j])
			}
			bottom--
		}

		// bottom -> top
		if left <= right {
			for i := bottom; i >= top; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
		}
	}

	return ans
}