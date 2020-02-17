package temp
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	var ret []int
	for head != nil {
		ret = append([]int{head.Val}, ret...)
		head = head.Next
	}
	return ret
}