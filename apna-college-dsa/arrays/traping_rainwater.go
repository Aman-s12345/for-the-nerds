package arrays

func (a *array) TrapBetter(height []int) int {
    n := len(height)
    left_max := make([]int, n)
    left_max[0] = height[0]
    for i := 1 ; i < n ; i++ {
        left_max[i] = max(left_max[i-1], height[i])
    }

    right_max := make([]int , n)
    right_max[n-1] = height[n-1]
     for i := n-2 ; i >=0 ; i-- {
        right_max[i] = max(right_max[i+1], height[i])
    }
    ans := 0
    for i :=1 ; i < n-1 ; i++ {
        ans = ans + (min(left_max[i] , right_max[i]) - height[i])
    }
    return ans

}

func (a *array) TrapOptimal(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	lmax := height[0]
	rmax := height[n-1]
	low := 1
	high := n-2
	ans := 0

	for low <= high {
		lmax = max(lmax, height[low])
		rmax = max(rmax, height[high])

		if lmax < rmax {
			ans += lmax - height[low]
			low++
		}else {
			ans += rmax - height[high]
			high--
		}
	}
	return ans

}
