package leetcode234

import "github.com/Julian-Chu/Leetcode-golang/utils"

type ListNode = utils.ListNode

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	prev := &ListNode{}

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	first, second := prev, slow
	if fast != nil {
		second = slow.Next
	}

	for second != nil {
		if first.Val != second.Val {
			return false
		}

		first = first.Next
		second = second.Next
	}

	return true
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome_1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast, prev := head, head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}

	prev.Next = nil
	cur1 := head
	cur2 := reverse(slow)

	for cur1 != nil {
		if cur1.Val != cur2.Val {
			return false
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummyHead := &ListNode{}

	for head != nil {
		next := head.Next
		head.Next = dummyHead.Next
		dummyHead.Next = head
		head = next
	}
	return dummyHead.Next
}
