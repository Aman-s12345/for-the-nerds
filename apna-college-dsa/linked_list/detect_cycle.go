package linkedlist


func (l *linkedList) DetectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head

	// Detect cycle
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			break
		}
	}

	// No cycle
	if fast == nil || fast.Next == nil {
		return nil
	}

	// Find starting point of cycle
	slow = head

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}