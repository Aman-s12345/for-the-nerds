package linkedlist

type NodeRandom struct {
	Val    int
	Next   *NodeRandom
	Random *NodeRandom
}

func (l *linkedList) CopyRandomList(head *NodeRandom) *NodeRandom {
	if head == nil {
		return nil
	}
	lookup := make(map[*NodeRandom]*NodeRandom)

	curr := head

	for curr != nil {
		lookup[curr] = &NodeRandom{Val: curr.Val}
		curr = curr.Next
	}

	curr = head

	for curr != nil {
		clone := lookup[curr]
		clone.Next = lookup[curr.Next]
		clone.Random = lookup[curr.Random]
		curr = curr.Next
	}

	return lookup[head]
}
