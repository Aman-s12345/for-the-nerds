package recursionbacktracking

func (a *rb) CheckValidGrid(grid [][]int) bool {

    // must start from top-left
    if grid[0][0] != 0 {
        return false
    }

	steps := [][]int{
        {2, 1}, {2, -1}, {-2, -1}, {-2, 1},
        {1, 2}, {-1, 2}, {1, -2}, {-1, -2},
    }

	return riding(grid, 0, 0, steps, 0)
}

func riding(ground [][]int, x int, y int, steps [][]int, idx int) bool {

	if idx >= len(ground)*len(ground[0])-1 {
		return true
	}

	for _, val := range steps {

		newX := x + val[0]
		newY := y + val[1]

		if newX >= 0 && newX < len(ground) &&
			newY >= 0 && newY < len(ground[0]) {

			if ground[newX][newY] == idx+1 {

				if riding(ground, newX, newY, steps, idx+1) {
					return true
				}
			}
		}
	}

	return false
}