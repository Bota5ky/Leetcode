package temp

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//双指针
func kthToLast2(head *ListNode, k int) int {
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
