package lintcode103

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			for slow.Next != nil && head.Next != nil {
				slow = slow.Next
				head = head.Next
				if slow == head {
					return head
				}
			}
		}
	}

	return nil
}
