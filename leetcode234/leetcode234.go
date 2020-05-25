package leetcode234

import "Leetcode-golang/utils"

type ListNode = utils.ListNode

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	s := make([]int, 0)

	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}

	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
