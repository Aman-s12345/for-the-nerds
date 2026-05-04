package arrays

func (a *array) MaxArea(height []int) int {
    maxCapacity := 0
    left, right := 0 , len(height)-1
    for left < right {
        maxCapacity = max(maxCapacity, calWater(height[left], left , height[right], right))
        if (height[left] > height[right]){
            right--
        }else{
            left++
        }
    }
    return maxCapacity
   
}

func calWater(x1 , y1 , x2, y2 int) int{
    return (y2-y1)*min(x1,x2)
}

func min(a, b int) int{
    if a>b {
        return b
    }
    return a
}