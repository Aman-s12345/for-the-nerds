package recursionbacktracking

type RB interface {

	// Combination Sum I
	CombinationSumINaive(candidates []int, target int) [][]int 
	CombinationSumIOptimized(candidates []int, target int) [][]int

	// Partition Palindrome
	Partition(s string) [][]string 

	// Solve N Queens
	SolveNQueens(n int) [][]string 

	// knight move
	CheckValidGrid(grid [][]int) bool

	// Subset with duplicates
	SubsetsWithDup(nums []int) [][]int

	// Merge Sort
	SortArray(nums []int) []int

}