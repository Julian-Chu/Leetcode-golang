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

	return FindMidAndBuildBST(head, nil)

}

func FindMidAndBuildBST(start, end *ListNode) *TreeNode {

	if start == end {
		return nil
	}
	if start.Next == end {
		return &TreeNode{Val: start.Val}
	}

	fast, slow := start, start
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}

	mid := slow

	return &TreeNode{
		Val:   mid.Val,
		Left:  FindMidAndBuildBST(start, mid),
		Right: FindMidAndBuildBST(mid.Next, end),
	}
}
