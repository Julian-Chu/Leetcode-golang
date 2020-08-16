package leetcode25

import "Leetcode-golang/utils"

type ListNode = utils.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	preHead := &ListNode{Next: head}
	pre := preHead
	cur := head
	tail := cur

	for tail != nil {
		for i := 0; i < k-1; i++ {
			tail = tail.Next
			if tail == nil {
				pre.Next = cur
				break
			}
		}
		if tail != nil {
			tmp := tail.Next
			tail.Next = nil

			newHead, newTail := reverse(cur)

			pre.Next = newHead
			pre = newTail
			cur, tail = tmp, tmp
		}

	}

	return preHead.Next
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil {
		return head, head
	}

	if head.Next.Next == nil {
		tmp := head
		head = head.Next
		tmp.Next = nil
		head.Next = tmp
		return head, head.Next
	}

	var pre *ListNode
	cur := head
	tail := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre, tail
}
