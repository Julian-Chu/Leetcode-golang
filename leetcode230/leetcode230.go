package leetcode230

import "Leetcode-golang/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = utils.TreeNode

func kthSmallest(root *TreeNode, k int) int {
	inoder := traverse(root)

	return inoder[k-1].Val
}

func traverse(root *TreeNode) []*TreeNode {
	if root == nil {
		return []*TreeNode{}
	}
	var left []*TreeNode
	if root.Left != nil {
		left = traverse(root.Left)
	}
	var right []*TreeNode
	if root.Right != nil {
		right = traverse(root.Right)
	}

	return append(append(left, root), right...)

}
