package arrays

import (
	"sort"
)

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

func (a *array) MergingSortedArrayWithoutSpace(nums1 []int, m int, nums2 []int, n int) []int{
	k := 0 
	for i := m ; i<m+n ; i++ {
		nums1[i] = nums2[k]
		k++
	}
	sort.Slice(nums1 , func(i, j int) bool {
		return nums1[i]< nums1[j]
	})
	return nums1

}


func (a *array) MergingSortedArrayWithoutSpaceOptimized(nums1 []int, m int, nums2 []int, n int) []int {
	i,j,k := m-1, n-1, len(nums1)-1
   for i >= 0 && j>= 0{
    if (nums1[i]>= nums2[j]){
        nums1[k] = nums1[i]
        i--
        k--
    }else{
        nums1[k] = nums2[j]
        j--
        k--
    }
   }
   if (i >= 0 ){
    for k >= 0 {
        nums1[k] = nums1[i]
        i--
        k--
    }
   }

   if (j >= 0 ){
    for k >= 0 {
        nums1[k] = nums2[j]
        j--
        k--
    }
   }
 return nums1
}

