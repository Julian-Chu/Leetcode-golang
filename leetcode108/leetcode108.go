package leetcode108

import "Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{
		Val: nums[mid],
	}

	if mid > 0 {
		root.Left = sortedArrayToBST(nums[:mid])
	}
	if mid < len(nums)-1 {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}

	return root

}
