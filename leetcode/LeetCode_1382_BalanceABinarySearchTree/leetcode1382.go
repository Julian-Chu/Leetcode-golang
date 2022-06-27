package LeetCode_1382_BalanceABinarySearchTree

import "github.com/Julian-Chu/Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func balanceBST(root *TreeNode) *TreeNode {
	nums := make([]int, 0)

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		nums = append(nums, node.Val)
		dfs(node.Right)
	}

	dfs(root)

	var buildBST func(nums []int) *TreeNode
	buildBST = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		l, r := 0, len(nums)-1

		mid := l + (r-l)>>1
		node := &TreeNode{
			Val: nums[mid],
		}

		node.Left = buildBST(nums[:mid])
		node.Right = buildBST(nums[mid+1:])
		return node
	}
	return buildBST(nums)
}

func isSameTree_iter(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	stack := []*TreeNode{p, q}

	for len(stack) > 0 {
		pNode := stack[len(stack)-2]
		qNode := stack[len(stack)-1]
		stack = stack[:len(stack)-2]

		if pNode.Val != qNode.Val {
			return false
		}

		if !(pNode.Left == nil && qNode.Left == nil) && (pNode.Left == nil || qNode.Left == nil) {
			return false
		}

		if !(pNode.Right == nil && qNode.Right == nil) && (pNode.Right == nil || qNode.Right == nil) {
			return false
		}

		if pNode.Left != nil && qNode.Left != nil {
			stack = append(stack, pNode.Left, qNode.Left)
		}
		if pNode.Right != nil && qNode.Right != nil {
			stack = append(stack, pNode.Right, qNode.Right)
		}
	}
	return true

}
