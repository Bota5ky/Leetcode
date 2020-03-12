package temp

func deleteNode(head *ListNode, val int) *ListNode {
	mark := &ListNode{0, head}
	res := mark
	for head != nil {
		if head.Val == val {
			mark.Next = mark.Next.Next
			break
		}
		mark = head
		head = head.Next
	}
	return res.Next
}