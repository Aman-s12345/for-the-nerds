package arrays


func (a *array) ProductExceptSelf(nums []int) []int {
    ans := make([]int , len(nums))
    ans[0] = 1
    for i:= 1 ; i < len(nums) ; i++ {
        ans[i] = nums[i-1]* ans[i-1]
    }
    sufix := nums[len(nums)-1]
    for i:= len(nums)-2; i>=0 ; i--{
        ans[i] = ans[i]*sufix
        sufix = sufix*nums[i]
    }
    return ans
}