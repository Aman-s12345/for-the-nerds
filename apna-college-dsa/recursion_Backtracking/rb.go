package recursionbacktracking

type RB interface {

	// Combination Sum I
	CombinationSumINaive(candidates []int, target int) [][]int 
	CombinationSumIOptimized(candidates []int, target int) [][]int

	// Partition Palindrome
	Partition(s string) [][]string 

}