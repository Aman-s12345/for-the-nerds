package arrays

import (
	"container/heap"
)

// MAX HEAP approach

type Pair struct {
	idx int
	val int
}

type PairHeap []Pair

// Len func implement
func (h PairHeap) Len() int {
	return len(h)
}

// Swap func implement
func (h PairHeap) Swap(i, j int) {
	h[i] , h[j] = h[j], h[i]
}

// Less func implement
func (h PairHeap) Less(i, j int) bool {
	return h[i].val > h[j].val
}

// Push func implement
func (h *PairHeap) Push(val interface{}){
	*h = append(*h, val.(Pair))
}

// pop func implement
func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]

	return val


}

func(a *array) MaxSlidingWindowWithHeap(nums []int, k int) []int {
	h := &PairHeap{}
	heap.Init(h)

	ans := []int{}

	for i:=0 ; i<len(nums) ; i++ {
		heap.Push(h, Pair{i ,nums[i]})
		for h.Len() > 0 && (*h)[0].idx <= i-k{
			heap.Pop(h)
		}
		if i >= k-1 {
			ans = append(ans, (*h)[0].val)
		}
	}
	return  ans


}