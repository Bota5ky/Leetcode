package leetcode

//https://leetcode-cn.com/problems/first-common-ancestor-lcci/
//和236相同
func lowestCommonAncestor3(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    if root==nil {return nil}
    if root==p {return p}
    if root==q {return q}
    left:=lowestCommonAncestor(root.Left,p,q)
    right:=lowestCommonAncestor(root.Right,p,q)
    if left!=nil && right!=nil {return root}
    if left!=nil {
        return left
    }
    return right
}