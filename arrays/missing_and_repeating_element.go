package arrays

func (a *array) MissingAndRepeatedValuesBruteForce(grid [][]int) []int {
	row := len(grid)
	col := len(grid[0])

	size := row * col

	visit := make(map[int]bool, size)
	var missing = -1
	var repeating = -1

	for _, arr := range grid {
		for _, val := range arr {
			if visit[val] {
				repeating = val
			} else {
				visit[val] = true
			}

		}
	}

	for i := 1; i < size+1; i++ {
		if !visit[i] {
			missing = i
			break
		}
	}

	return []int{repeating, missing}
}

