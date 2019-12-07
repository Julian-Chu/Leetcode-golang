package leetcode109

import "Leetcode-golang/utils"

type (
	TreeNode = utils.TreeNode
	ListNode = utils.ListNode
)

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var nums []int
	node := head
	for {
		nums = append(nums, node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	mid := len(nums) / 2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = buildBST(nums[:mid])
	root.Right = buildBST(nums[mid+1:])

	return root
}

func buildBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{
		Val: nums[mid],
	}
	root.Left = buildBST(nums[:mid])
	root.Right = buildBST(nums[mid+1:])
	return root
}
