package linkedlist

type LinkedList interface {

	// Middle of the Node
	MiddleNode(head *ListNode) *ListNode
	// Reverse the List
	ReverseList(head *ListNode) *ListNode
	// Detect Cycle
	 DetectCycle(head *ListNode) *ListNode 
	 // Has Cycle 
	 HasCycle(head *ListNode) bool
	 // Merge Two Sorted Lists
	 MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode
	 MergeTwoListsInPlace(list1 *ListNode, list2 *ListNode) *ListNode
	 // Flatten a Multilevel Doubly Linked List
	 Flatten(root *Node) *Node
	 // Check Palindrome
	 IsPalindrome(head *ListNode) bool
	 // Deep copy of a linked list with random pointer
	 CopyRandomList(head *NodeRandom) *NodeRandom
	 

	 
}

type linkedList struct{}

func NewLinkedList() LinkedList {
	return &linkedList{}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
