package arrays

func (a *array) MergingSortedArrays(nums1 []int, m int, nums2 []int, n int) []int {
	helper := make([]int, m+n)
	i, j, k := 0, 0, 0

	for i < m && j < n {
		if nums1[i] >= nums2[j] {
			helper[k] = nums2[j]
			j++
		} else {
			helper[k] = nums1[i]
			i++
		}
		k++
	}

	for i < m {
		helper[k] = nums1[i]
		i++
	}

	for j < n {
		helper[k] = nums2[j]
		j++
	}

	return helper
}
