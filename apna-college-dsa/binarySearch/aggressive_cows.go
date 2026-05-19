package binarysearch

import "sort"

func (a *bs) AggressiveCows(stalls []int, k int) int{
	sort.Ints(stalls)

	min := 1
	max := stalls[len(stalls)-1]-stalls[0]

	ans := 0 

	for min < max {
		mid := min + (max-min)/2

		if (canPlace(stalls, k , mid)){
			ans = mid
			min = mid+1
		}else {
			max = mid-1
		}
	}
	return ans

}

func canPlace(stalls []int, k int, dist int) bool {
	count := 1
	last := stalls[0]

	for i := 1; i < len(stalls); i++ {
		if stalls[i]-last >= dist {
			count++
			last = stalls[i]
		}

		if count >= k {
			return true
		}
	}

	return false
}