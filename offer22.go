package temp

func getKthFromEnd(head *ListNode, k int) *ListNode {
	rear := head
	for ; rear != nil && k > 0; k-- {
		rear = rear.Next
	}
	for rear != nil {
		rear = rear.Next
		head = head.Next
	}
	return head
}
