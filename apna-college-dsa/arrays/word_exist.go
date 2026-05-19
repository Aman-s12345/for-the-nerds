package arrays


func (a *array) WordExist(board [][]byte, word string) bool {

	for i := 0; i<len(board) ; i++ {
		for j := 0; j< len(board[0]) ; j++ {
			if (board[i][j] == word[0] && finderINBoard(board, word, i , j , 0)){
				return true
			}
		}
	}
	return false
}

func finderINBoard(board [][]byte, word string, i int, j int, index int) bool{
	if(index == len(word)){
		return true
	}

	if(i <0 || i >= len(board) || j < 0 || len(board[0]) <= j){
		return false
	}

	if (board[i][j] != word[index]){
		return false
	}

	temp := board[i][j]
	board[i][j] = '#'

	found := finderINBoard(board, word, index+1, i-1, j) ||
			finderINBoard(board, word, index+1, i+1, j) ||
			finderINBoard(board, word, index+1, i, j-1) ||
			finderINBoard(board, word, index+1, i, j+1)

	board[i][j] = temp

	return found
}