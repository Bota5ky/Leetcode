package temp

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := depthOfTreeNode(root.Left)
	right := depthOfTreeNode(root.Right)
	if left-right > 1 || left-right < -1 {
		return false
	}
	return isBalanced(root.Right) && isBalanced(root.Left)
}

func depthOfTreeNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := depthOfTreeNode(root.Left)
	right := depthOfTreeNode(root.Right)
	if left >= right {
		return 1 + left
	}
	return 1 + right
}
