package linkedlist

type LinkedList interface {

	// Middle of the Node
	MiddleNode(head *ListNode) *ListNode
	// Reverse the List
	ReverseList(head *ListNode) *ListNode
}

type linkedList struct{}

func NewLinkedList() LinkedList {
	return &linkedList{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
