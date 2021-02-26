package leetcode160

import "Leetcode-golang/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = utils.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	m := make(map[*ListNode]bool)

	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if m[headB] == true {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
