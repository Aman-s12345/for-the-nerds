package binarysearch

import "sort"

type Pair struct {
	val int
	idx int
}

func (a *bs) PeakIndexInMountainArrayNaive(arr []int) int {

	mountains := []Pair{}

	for i, val := range arr {
		mountains = append(mountains, Pair{
			val: val,
			idx: i,
		})
	}
	sort.Slice(mountains, func(i, j int) bool {
		return mountains[i].val < mountains[j].val
	})
	return mountains[len(arr)-1].idx
}



func (a *bs) PeakIndexInMountainArrayOptmized(arr []int) int {

	if len(arr) == 1 {
		return 0
	}

	return findPivot(arr, 0 , len(arr)-1)
}
func findPivot(arr []int, si int, ei int) int {
    if si == ei {
        return si
    }

    mid := si + (ei-si)/2

    if arr[mid] < arr[mid+1] {
        return findPivot(arr, mid+1, ei)
    }

    return findPivot(arr, si, mid)
}