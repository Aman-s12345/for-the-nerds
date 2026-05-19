package binarysearch

type BS interface {
	// Peak Index In Mountain Array
	PeakIndexInMountainArrayNaive(arr []int) int

	// Search in Rotated Sorted Array
	Search(nums []int, target int) int 

	// Single Non Duplicate
	SingleNonDuplicate(nums []int) int 

	// Aggresive Cows
	AggressiveCows(stalls []int, k int) int
}