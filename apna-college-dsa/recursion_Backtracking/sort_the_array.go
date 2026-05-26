package recursionbacktracking

func (a *rb) SortArray(nums []int) []int {
    
	divide(nums, 0 , len(nums)-1)

	return nums
}

func divide(nums []int, si int, ei int){
	if si >= ei {
		return
	}
	mid := si + (ei-si)/2
	divide(nums , si , mid)
	divide(nums, mid+1, ei)
	conqure(nums, si, mid, ei)
}

func conqure(arr []int, si int, mid int, ei int){
	ans := make([]int, ei-si+1)
	i := si
	j := mid+1
	k := 0

	for i <= mid && j <= ei {
		if arr[i] < arr[j] {
			ans[k] = arr[i]
			i++
		}else{
			ans[k] = arr[j]
			j++
		}
		k++
	}

	for i <= mid {
		ans[k] = arr[i]
		i++
		k++
	}

	for j <= ei {
		ans[k] = arr[j]
		j++
		k++
	}

	for idx := 0; idx < len(ans); idx++ {
		arr[si+idx] = ans[idx]
	}
}