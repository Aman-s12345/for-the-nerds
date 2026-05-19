package arrays


func (a *array) SearchMatrixBetter(matrix [][]int, target int) bool {

    for _,val := range matrix {
        if(search(val,target,0,len(val))){
            return true
        }
    }

    return false
}

func search (slice []int, t int, si int, ei int) bool {
    if si>=ei { return false}
    mid := si + (ei-si)/2
    if slice[mid] == t {
        return true
    }else if slice[mid] > t {
      return  search(slice, t, si, mid)
    }else {
        return search(slice, t, mid+1, ei)
    }
}


func (a *array) SearchMatrixOptimal(matrix [][]int, target int) bool {

	rows := len(matrix)
	cols := len(matrix[0])

	row := 0
	col := cols - 1

	for row < rows && col >= 0 {

		current := matrix[row][col]

		if current == target {
			return true
		}

		if current > target {

			// move left
			col--

		} else {

			// move down
			row++
		}
	}

	return false
}