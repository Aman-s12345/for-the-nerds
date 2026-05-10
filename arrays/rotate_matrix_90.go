package arrays

func (a *array) RotateNeive(matrix [][]int)  {
    ans := make([][]int , len(matrix))

     for i := 0; i < len(matrix); i++ {
        ans[i] = make([]int , len(matrix) )
    } 

    for i := 0; i < len(matrix); i++ {
        
        for j, val := range matrix[i] {
            ans[j][len(matrix)-1-i] = val
        }
    }

    for i := 0; i < len(matrix); i++ {
        
        for j,_ := range matrix[i] {
            matrix[i][j] = ans[i][j]
        }
    }

}

func (a *array) RotateOptimal(matrix [][]int)  {
    // take a transpose
    for i := 0 ; i < len(matrix) ; i++ {
        for j :=0; j < len(matrix[0]) ; j++ {
            if i == j || i > j {
                continue
            } else {
              matrix[i][j] , matrix[j][i] = matrix[j][i] , matrix[i][j]
            }

        }
    }
	// reverse row
    for _, val := range matrix {
        reverseSlice(val)
    }


}

func reverseSlice(slice []int){
    i , j := 0 , len(slice)-1
    for i < j {
        slice[i] , slice[j] = slice[j], slice[i]
        i++ 
        j--
    }
}