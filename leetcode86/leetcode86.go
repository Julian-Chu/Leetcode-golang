package leetcode86

import (
	"Leetcode-golang/utils"
)

type ListNode = utils.ListNode

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	lessHead := &ListNode{}
	moreEqHead := &ListNode{}

	lessEnd := lessHead
	moreEqEnd := moreEqHead

	for head != nil {
		if head.Val < x {
			lessEnd.Next = head
			lessEnd = lessEnd.Next
		} else {
			moreEqEnd.Next = head
			moreEqEnd = moreEqEnd.Next
		}
		head = head.Next
	}

	lessEnd.Next = moreEqHead.Next
	moreEqEnd.Next = nil
	return lessHead.Next
}
