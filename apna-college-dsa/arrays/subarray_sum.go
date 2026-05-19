package arrays

func (a *array) SubarraySumNieve(nums []int, k int) int {
	noOfSum := 0
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum == k {
			noOfSum++
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				noOfSum++
			}
		}
	}
	return noOfSum
}

func (a *array) SubarraySumOptimal(nums []int, k int) int {
	ans := 0
	prefixSum := make(map[int]int, len(nums))

	prefixSum[0] = 1
	currentSum := 0

	for i := 0; i < len(nums); i++ {
		currentSum = currentSum + nums[i]
		if prefixSum[currentSum-k] > 0 {
			ans += prefixSum[currentSum-k]
		}
		prefixSum[currentSum]++

	}
	return ans
}
