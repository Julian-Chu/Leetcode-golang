package lintcode102

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * @param head: The first node of linked list.
 * @return: True if it has a cycle, or false
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
