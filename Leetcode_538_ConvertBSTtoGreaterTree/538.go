package Leetcode_538_ConvertBSTtoGreaterTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	traverse(root, &sum)
	return root
}

func traverse(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	traverse(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	traverse(root.Left, sum)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
