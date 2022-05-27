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
