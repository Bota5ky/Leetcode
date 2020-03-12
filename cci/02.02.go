package leetcode

//https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/
//双指针
func kthToLast(head *ListNode, k int) int {
	rear := head
	for rear != nil && k > 0 {
		rear = rear.Next
		k--
	}
	for rear != nil {
		head = head.Next
		rear = rear.Next
	}
	return head.Val
}
