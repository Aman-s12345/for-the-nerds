package recursionbacktracking

import "sort"

func CombinationSumINaive(candidates []int, target int) [][]int {
	ans := [][]int{}

	findSum(candidates, target, 0, 0, &ans, []int{})

	return ans

}

func findSum(arr []int, tar int, sum int, idx int, ans *[][]int, store []int) {

	if tar == sum {
		temp := make([]int, len(store))
		copy(temp, store)
		*ans = append(*ans, temp)
		return
	}

	if sum > tar || idx >= len(arr) {
		return
	}

	// add current number in the store
	store = append(store, arr[idx])
	findSum(arr, tar, sum+arr[idx], idx, ans, store)
	// return the last element
	store = store[:len(store)-1]
	// increment the idx
	findSum(arr, tar, sum, idx+1, ans, store)

}



func CombinationSumIOptimized(candidates []int, target int) [][]int {
	ans := [][]int{}

	sort.Ints(candidates)

	findSumOptimized(candidates, target, 0, 0, &ans, []int{})

	return ans

}


func findSumOptimized(arr []int, tar int, sum int, idx int, ans *[][]int, store []int) {

	if tar == sum {
		temp := make([]int, len(store))
		copy(temp, store)
		*ans = append(*ans, temp)
		return
	}

	if sum > tar || idx >= len(arr) {
		return
	}

	if sum+arr[idx] > tar {
		return
	}

	// add current number in the store
	store = append(store, arr[idx])
	findSumOptimized(arr, tar, sum+arr[idx], idx, ans, store)
	// return the last element
	store = store[:len(store)-1]
	// increment the idx
	findSumOptimized(arr, tar, sum, idx+1, ans, store)

}


