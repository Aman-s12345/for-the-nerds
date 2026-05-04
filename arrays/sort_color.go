package arrays

import "sort"

func (a *array) SortColorsNaive(nums []int) {
	sort.Slice(nums , func(i, j int) bool {
		return  nums[i] < nums[j]
	})
}


func (a *array) SortColorsOptimized(nums []int) {
	low, mid, high := 0, 0, len(nums)-1

    for mid <= high {
        if nums[mid] == 0 {
            nums[low], nums[mid] = nums[mid], nums[low]
            low++
            mid++
        } else if nums[mid] == 1 {
            mid++
        } else { 
            nums[mid], nums[high] = nums[high], nums[mid]
            high--
        }
    }
}
