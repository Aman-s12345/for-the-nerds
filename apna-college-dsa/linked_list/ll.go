package linkedlist

type LinkedList interface {
}

type linkedList struct{}

func NewLinkedList() LinkedList {
	return &linkedList{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
