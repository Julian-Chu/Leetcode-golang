package leetcode725

import "github.com/Julian-Chu/Leetcode-golang/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = utils.ListNode

func splitListToParts(root *ListNode, k int) []*ListNode {
	size := 0

	head := root
	for head != nil {
		size++
		head = head.Next
	}

	remainder := size % k
	res := make([]*ListNode, k)
	i := 0
	for {
		res[i] = root
		i++
		if i == k {
			break
		}

		leng := size / k
		if remainder > 0 {
			leng++
			remainder--
		}

		for leng > 1 && root != nil {
			leng--
			root = root.Next
		}

		if root == nil {
			break
		}
		root.Next, root = nil, root.Next
	}

	return res
}
