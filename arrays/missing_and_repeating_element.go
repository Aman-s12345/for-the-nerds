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

func (a *array) MissingAndRepeatedValuesMath(grid [][]int) []int{
	 n := len(grid)*len(grid[0])
	idealNSum := (n*(n+1))/2
	foundSum := 0 
	idealNSquareSum := (n*(n+1)*(2*n+1))/6
	foundNSquareSum := 0
	for i := 0 ; i<len(grid) ; i++ {
		for j := 0; j<len(grid[0]) ; j++ {
			num := grid[i][j]
			foundSum += num
			foundNSquareSum += num*num
		}
	}

	ans := make([]int , 2)
	aMinusB := foundSum - idealNSum
	aSquareMinusBSquare := foundNSquareSum - idealNSquareSum

	aPlusB := aSquareMinusBSquare/aMinusB

	A := (aPlusB + aMinusB)/2
	B := (aPlusB - aMinusB)/2
    
	ans[0] = A
	ans[1] = B

	return ans



}
