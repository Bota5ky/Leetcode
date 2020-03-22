package leetcode

//https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{root} //从根节点开始
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1] //每次迭代弹出当前栈顶元素
		res = append([]int{node.Val}, res...)
		if node.Left != nil {
			stack = append(stack, node.Left) //并将其孩子节点压入栈中
		}
		if node.Right != nil {
			stack = append(stack, node.Right) //先压左孩子再压右孩子
		}
	}
	return res
}
