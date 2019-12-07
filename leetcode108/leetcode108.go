package leetcode108

import "Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	return dfs(nums, 0, len(nums)-1)
}

func dfs(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}

	mid := (l + r) / 2
	root := &TreeNode{
		Val: nums[mid],
	}

	root.Left = dfs(nums, l, mid-1)
	root.Right = dfs(nums, mid+1, r)
	return root
}
