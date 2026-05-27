package linkedlist

func (l *linkedList) MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}

	ptr1 := list1
	ptr2 := list2
	ptr3 := dummy

	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val >= ptr2.Val {
			ptr3.Next = &ListNode{Val: ptr2.Val}
			ptr3 = ptr3.Next
			ptr2 = ptr2.Next
		} else {
			ptr3.Next = &ListNode{Val: ptr1.Val}
			ptr3 = ptr3.Next
			ptr1 = ptr1.Next
		}
	}
	for ptr1 != nil {
		ptr3.Next = &ListNode{Val: ptr1.Val}
		ptr3 = ptr3.Next
		ptr1 = ptr1.Next
	}
	for ptr2 != nil {
		ptr3.Next = &ListNode{Val: ptr2.Val}
		ptr3 = ptr3.Next
		ptr2 = ptr2.Next
	}

	return dummy.Next
}

// without any extra space allocated
func (l *linkedList) MergeTwoListsInPlace(list1 *ListNode, list2 *ListNode) *ListNode {
	curr := &ListNode{Val: 0}
	dummy := curr
	ptr1 := list1
	ptr2 := list2

	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val >= ptr2.Val {
			dummy.Next = ptr2
			ptr2 = ptr2.Next
			dummy = dummy.Next
		}else {
			dummy.Next = ptr1
			ptr1 = ptr1.Next
			dummy = dummy.Next
		}
	}
	if ptr1 != nil {
		dummy.Next = ptr1

	}
	if ptr2 != nil {
		dummy.Next = ptr2	
	}

	return curr.Next


}