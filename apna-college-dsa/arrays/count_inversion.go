package arrays

func (a *array) CountInversions(arr []int32) int64 {
	var ans int64
	inversionCount(arr, 0, len(arr)-1, &ans)
	return ans

}

func inversionCount(arr []int32, si int, ei int, ans *int64) {
	if si >= ei {
		return
	}
	mid := si + (ei-si)/2
	inversionCount(arr, si, mid, ans)
	inversionCount(arr, mid+1, ei, ans)
	count(arr, si, mid, ei, ans)

}
func count(arr []int32, si int, mid int, ei int, ans *int64) {
	p1 := si
	p2 := mid + 1
	temp := make([]int32, ei-si+1)
	j := 0
	for p1 <= mid && p2 <= ei {
		if arr[p1] <= arr[p2] {
			temp[j] = arr[p1]
			p1++
			j++
		} else {
			temp[j] = arr[p2]
			j++
			*ans += int64(mid - p1 + 1)
			p2++
		}
	}
	for p1 <= mid {
		temp[j] = arr[p1]
		p1++
		j++
	}
	for p2 <= ei {
		temp[j] = arr[p2]
		j++
		p2++

	}
	j = 0
	for i := si; i <= ei; i++ {
		arr[i] = temp[j]
		j++
	}

}
