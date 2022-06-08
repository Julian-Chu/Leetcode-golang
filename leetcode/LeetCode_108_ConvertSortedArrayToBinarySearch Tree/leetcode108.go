package LeetCode_108_ConvertSortedArrayToBinarySearch_Tree

import "github.com/Julian-Chu/Leetcode-golang/utils"

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

func sortedArrayToBST_1(nums []int) *TreeNode {
	size := len(nums)
	if size == 0 {
		return nil
	}

	mid := size / 2

	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
