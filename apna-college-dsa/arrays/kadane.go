package arrays

func (a *array) MaxSubArray(nums []int) int {
    currentSum, maxSubArraySum := nums[0], nums[0]
    for i := 1 ; i < len(nums) ; i++{ 
        val := nums[i]
         currentSum = max(currentSum+val ,val)
         maxSubArraySum = max(currentSum,maxSubArraySum) 
    }
    return maxSubArraySum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}