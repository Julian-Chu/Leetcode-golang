package LeetCode_24_SwapNodesInPairs

import (
	"github.com/Julian-Chu/Leetcode-golang/utils"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = utils.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(head.Next.Next)
	newHead.Next = head
	return newHead
}

//func swapPairs(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil{
//		return head
//	}
//
//	dummy:= &ListNode{
//		Next: head,
//	}
//
//	cur:= dummy
//
//	for cur.Next != nil&& cur.Next.Next != nil{
//		prev:= cur.Next
//		next:= cur.Next.Next
//		tmp:= next.Next
//
//		prev.Next = tmp
//		next.Next = prev
//		cur.Next = next
//		cur = cur.Next.Next
//	}
//
//	return dummy.Next
//}
