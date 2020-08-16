package leetcode25

import "Leetcode-golang/utils"

type ListNode = utils.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 || head == nil || head.Next == nil {
		return head
	}

	tail, needReverse := getTail(head, k)

	if needReverse {
		tailNext := tail.Next
		tail.Next = nil
		head, tail = reverse(head, tail)
		tail.Next = reverseKGroup(tailNext, k)

	}

	return head
}

func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	var preCur *ListNode
	cur := head

	for cur != nil {
		preCur, cur, cur.Next = cur, cur.Next, preCur
	}

	return tail, head
}

func getTail(head *ListNode, k int) (*ListNode, bool) {
	for k > 1 && head != nil {
		head = head.Next
		k--
	}

	return head, k == 1 && head != nil
}
