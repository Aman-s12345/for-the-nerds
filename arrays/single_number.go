package arrays

func (a *array) SingleNumberBruteForce(nums []int) int {
    lookup := make(map[int]int , len(nums))
    for _,val := range nums {
        count , ok := lookup[val]
        if !ok {
            lookup[val] = 1
        }else{
            lookup[val] = count+1
        }
    }

    for _,val := range nums {
        count , _ := lookup[val]
        if count == 1 {
            return val
        }
    }
    return -1

}

func (a *array) SingleNumberoptimized(nums []int) int {
    result := 0
    for _, val := range nums{
        result = result ^ val
    }
    return result

}