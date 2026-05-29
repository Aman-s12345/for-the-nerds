package linkedlist


func (l *linkedList) IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// find middle
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse second half
	var prev *ListNode
	curr := slow

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// compare
	left := head
	right := prev

	for right != nil {
		if left.Val != right.Val {
			return false
		}

		left = left.Next
		right = right.Next
	}

	return true
}