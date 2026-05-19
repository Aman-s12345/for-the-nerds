package arrays

import "sort"

func (a *array) ThreeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--

				for j < k && nums[j] == nums[j-1] {
					j++
				}

				for j < k && nums[k] == nums[k+1] {
					k--
				}

			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return ans

}
