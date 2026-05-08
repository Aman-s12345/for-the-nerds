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
	MergingSortedArrayWithoutSpace(nums1 []int, m int, nums2 []int, n int) []int
	MergingSortedArrayWithoutSpaceOptimized(nums1 []int, m int, nums2 []int, n int) []int

	//  Single Number
	SingleNumberBruteForce(nums []int) int 
	SingleNumberoptimized(nums []int) int 

	// Stock Buy and sell
	MaxProfitOptimized(prices []int) int 

	// Math Pow
	MyPowNaive(x float64, n int) float64
	MyPowOptmized(x float64, n int) float64 

	// Kadane Algorithm
	MaxSubArray(nums []int) int 

	// Max Area
	MaxArea(height []int) int 

	// SortColor
	SortColorsNaive(nums []int) 
	SortColorsOptimized(nums []int) 

	// three sum
	ThreeSum(nums []int) [][]int

	// four sum
	FourSumOptmized(nums []int, target int) [][]int
	
	// search Matric
	SearchMatrix(matrix [][]int, target int) bool 

	// set matrix zero
	SetZeroesBruteForce(matrix [][]int)
	SetZeroesBetter(matrix [][]int)
	SetMatrixOptimal(matrix [][]int) 

	// Merge overlapping interval
	MergeInterval(intervals [][]int) [][]int

	// Length Of Longest Substring
	LengthOfLongestSubstring(s string) int 

	// Word Search
	WordExist(board [][]byte, word string) bool



    
}

