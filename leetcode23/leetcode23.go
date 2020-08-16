package leetcode23

import "Leetcode-golang/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = utils.ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)

	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	mid := n / 2

	return helper(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))
}

func helper(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		l2.Next = helper(l1, l2.Next)
		return l2
	} else {
		l1.Next = helper(l1.Next, l2)
		return l1
	}
}
