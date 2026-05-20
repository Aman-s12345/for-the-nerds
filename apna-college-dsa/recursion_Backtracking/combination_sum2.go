package recursionbacktracking

import "sort"

func (a *rb) CombinationSum2(candidates []int, target int) [][]int {
	  ans := [][]int{}
    sort.Ints(candidates)
	findSumII(candidates, target, 0, 0, &ans, []int{})

	return ans
}

func findSumII(arr []int, tar int, sum int, idx int, ans *[][]int, store []int) {
	if tar == sum {
		temp := make([]int, len(store))
		copy(temp, store)
		*ans = append(*ans, temp)
		return
	}

	if sum > tar || idx >= len(arr) {
		return
	}

	store = append(store, arr[idx])
	findSumII(arr, tar, sum+arr[idx], idx+1, ans, store)
	store = store[:len(store)-1]
	for idx+1 < len(arr) && arr[idx] == arr[idx+1] {
		idx++
	}
	findSumII(arr, tar, sum, idx+1, ans, store)
}
