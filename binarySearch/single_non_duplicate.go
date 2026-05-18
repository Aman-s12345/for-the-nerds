package binarysearch

func (a *bs) SingleNonDuplicate(nums []int) int {
	left := 0 
    right := len(nums)-1

    for left < right {
        mid := left + (right-left)/2
        if mid%2 == 1 {
            mid--
        }

        if nums[mid] == nums[mid+1] {
            left = left + 2
        } else {
            right = mid - 1
        }
    }
    return nums[left]
}
