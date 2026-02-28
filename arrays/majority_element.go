package arrays

import "sort"

type array struct{}

func NewArray() Array {
	return &array{}
}

func (a *array) MajorityElementBruteForce(num []int) int {
	return 0

}

func (a *array) MajorityElementHashMap(num []int) int {
	hm := make(map[int]int)
	for _, v := range num {
		hm[v]++
		if hm[v] > len(num)/2 {
			return v
		}
	}
	return -1

}

func (a *array) MajorityElementSorting(num []int) int {
	sort.Ints(num)
	return num[len(num)/2]
}

func (a *array) MajorityElementBoyerMoore(num []int) int {
	count := 0
	candidate := num[0]
	for _, v := range num {
		if v == candidate {
			count++
		} else {
			count--
		}

		if count == 0 {
			count = 1
			candidate = v
		}
	}

	return candidate
}
