package arrays

import "sort"

func (a *array) FourSumOptmized(nums []int, target int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	ans := [][]int{}

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			k, l := j+1, len(nums)-1

			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]

				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--

					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}

				} else if sum < target {
					k++
				} else {
					l--
				}
			}
		}
	}

	return ans

}
