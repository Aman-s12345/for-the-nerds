package linkedlist

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func (l *linkedList) Flatten(root *Node) *Node {

	if root != nil {
		return root
	}
	dummy := &Node{}
	doFlatten(root, dummy)
	if dummy.Next != nil {
		dummy.Next.Prev = nil
	}
	return dummy.Next

}

func doFlatten(curr *Node, tail *Node) *Node {
	for curr != nil {

		// create a node
		newMode := &Node{Val: curr.Val}
		tail.Next = newMode
		
		tail.Next.Prev = tail

		tail = tail.Next

		// next store
		next := curr.Next

		if curr.Child != nil {
			doFlatten(curr.Child, tail)
		}
		curr = next


	}
	return tail

}
