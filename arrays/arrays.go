package arrays


type Array interface {
	// define the interface MajorityElement
	MajorityElementBruteForce (num []int) int
	MajorityElementHashMap (num []int) int
	MajorityElementSorting (num []int) int
	MajorityElementBoyerMoore (num []int) int

	// define the interface MissingAndRepeatedValues
	MissingAndRepeatedValuesBruteForce(grid [][]int) []int
	MissingAndRepeatedValuesMath(grid [][]int) []int

	// Merging two sorted arrays
    MergingSortedArrays(nums1 []int, m int, nums2 []int, n int) []int



	// Set Matrix Zero 
	// SetZeroes(matrix [][]int)
	
    
}

