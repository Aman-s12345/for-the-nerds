package recursionbacktracking



func (a *rb) SolveNQueens(n int) [][]string {
	ans := [][]string{}


	// create empty chess board
	board := make([][]byte, n)

	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)

		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	var backtrack func(row int)

	backtrack = func(row int) {
		if row == n {
			temp := []string{}

			for i := 0; i < n; i++ {
				temp = append(temp, string(board[i]))
			}

			ans = append(ans, temp)
			return
		}

		for col := 0; col < n; col++ {
			if isSafe(board, row, col, n) {

				board[row][col] = 'Q'

				backtrack(row + 1)

				board[row][col] = '.'
			}
		}
	}

	backtrack(0)

	return ans
}

func isSafe(board [][]byte, row int, col int, n int) bool {

	// check upper column
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// upper left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// upper right diagonal
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}