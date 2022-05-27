package leetcode1367

import "github.com/Julian-Chu/Leetcode-golang/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type ListNode = utils.ListNode
type TreeNode = utils.TreeNode

func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}

	return dfs(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	return head.Val == root.Val && (dfs(head.Next, root.Left) || dfs(head.Next, root.Right))
}
