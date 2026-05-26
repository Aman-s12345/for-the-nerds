package recursionbacktracking

import "sort"

func (a *rb) SubsetsWithDup(nums []int) [][]int {

	ans := [][]int{}
	sort.Ints(nums)

	allSubset(nums, 0, []int{}, &ans)
	return ans
}

func allSubset(arr []int, idx int, store []int, ans *[][]int) {

	temp := make([]int, len(store))
	copy(temp, store)
	*ans = append(*ans, temp)

	for i := idx; i < len(arr); i++ {

		// skip duplicates
		if i > idx && arr[i] == arr[i-1] {
			continue
		}
		// taking the element
		store = append(store, arr[i])
		allSubset(arr, i+1, store, ans)

		// backtracking
		store = store[:len(store)-1]

	}
}
