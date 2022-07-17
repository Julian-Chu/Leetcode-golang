package LeetCode_203_RemoveLinkedListElements

import "github.com/Julian-Chu/Leetcode-golang/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = utils.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	p := &head
	if *p == nil {
		return nil
	}

	for *p != nil {
		if (*p).Val == val {
			*p = (*p).Next
		} else {
			p = &(*p).Next
		}
	}

	return head
}

//func removeElements(head *ListNode, val int) *ListNode {
//	if head == nil {
//		return head
//	}
//	preHead := &ListNode{Next: head}
//	cur := preHead
//
//	for cur.Next != nil {
//		if cur.Next.Val == val {
//			cur.Next = cur.Next.Next
//		} else {
//			cur = cur.Next
//		}
//	}
//	return preHead.Next
//}

func removeElements_recur(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val == val {
		return removeElements(head.Next, val)
	}

	head.Next = removeElements(head.Next, val)
	return head

}
