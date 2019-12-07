package leetcode109

import "Leetcode-golang/utils"

type (
	TreeNode = utils.TreeNode
	ListNode = utils.ListNode
)

func sortedListToBST(head *ListNode) *TreeNode {
	var nums []int
	node := head
	for {
		nums = append(nums, node.Val)
		if node.Next == nil {
			break
		}
		node = node.Next
	}

}
