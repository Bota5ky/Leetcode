package temp

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	var res []string
	temp := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		res = append(res, temp)
		return res
	}
	routine(root.Left, temp, &res)
	routine(root.Right, temp, &res)
	return res
}

//之前积累的字符串
func routine(node *TreeNode, temp string, res *[]string) {
	if node == nil {
		return
	}
	num := strconv.Itoa(node.Val)
	temp += "->" + num
	if node.Left == nil && node.Right == nil {
		*res = append(*res, temp)
		return
	}
	routine(node.Left, temp, res)
	routine(node.Right, temp, res)
}
