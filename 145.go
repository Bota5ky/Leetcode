package leetcode

//https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
func postorderTraversal(root *TreeNode) []int {
	cur := root
	var res []int
	var stack []*TreeNode
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			res = append([]int{cur.Val}, res...)
			stack = append(stack, cur)
			cur = cur.Right
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = node.Left
		}
	}
	return res
}
