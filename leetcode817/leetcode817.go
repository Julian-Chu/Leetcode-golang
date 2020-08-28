package leetcode817

import "Leetcode-golang/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = utils.ListNode

func numComponents(head *ListNode, G []int) int {
	m := make(map[int]bool)
	for i := range G {
		m[G[i]] = true
	}

	connected := m[head.Val]

	cnt := 0
	for head != nil {
		if connected != m[head.Val] {
			if connected == false {
				connected = true
			} else {
				connected = false
				cnt++
			}
		}
		head = head.Next
	}

	if connected {
		return cnt + 1
	}
	return cnt
}
