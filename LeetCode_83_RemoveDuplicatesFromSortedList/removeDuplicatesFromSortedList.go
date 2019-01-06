package LeetCode_83_RemoveDuplicatesFromSortedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	mapper := map[int]bool{}
	var pre, cur *ListNode
	cur = head
	for cur != nil {
		if mapper[cur.Val] == true {
			if cur.Next == nil {
				pre.Next = nil
			}
			cur = cur.Next
		} else {
			mapper[cur.Val] = true
			if pre != nil {
				pre.Next = cur
			}
			pre = cur
			cur = pre.Next
		}
	}

	return head
}
