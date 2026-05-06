package arrays

func (a *array) SetZeroesBruteForce(matrix [][]int) {
	rows := map[int]bool{}
	cols := map[int]bool{}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	ans := make([][]int, len(matrix))

	for i := 0; i < len(matrix); i++ {
		ans[i] = make([]int, len(matrix[0]))
		if rows[i] == true {
			continue
		}
		for j := 0; j < len(matrix[0]); j++ {
			if cols[j] == true {
				ans[i][j] = 0
			} else {
				ans[i][j] = matrix[i][j]
			}
		}

	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] = ans[i][j]
		}
	}

}

func (a *array) SetZeroesBetter(matrix [][]int) {
	rows := map[int]bool{}
	cols := map[int]bool{}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}

}
func (a *array) SetMatrixOptimal(matrix [][]int) {
	isFirstRowZero := false
	isFirstColZero := false

	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			isFirstRowZero = true
		}
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			isFirstColZero = true
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0   
				matrix[i][0] = 0  
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if isFirstRowZero {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	if isFirstColZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}