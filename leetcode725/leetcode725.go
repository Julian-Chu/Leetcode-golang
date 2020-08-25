package leetcode725

import "Leetcode-golang/utils"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = utils.ListNode

func splitListToParts(root *ListNode, k int) []*ListNode {
	if root == nil {
		return make([]*ListNode, k)
	}
	size := 0

	head := root
	for head != nil {
		size++
		head = head.Next
	}

	remainder := size % k
	length := size / k
	lenArr := make([]int, 0, k)

	for i := 0; i < k; i++ {
		if remainder != 0 {
			lenArr = append(lenArr, length+1)
			remainder--
			continue
		}
		lenArr = append(lenArr, length)
	}

	res := make([]*ListNode, 0, k)

	head = root
	for _, l := range lenArr {
		var pre *ListNode
		res = append(res, head)

		for head != nil && l > 0 {
			pre = head
			head = head.Next
			l--
		}
		if pre != nil {
			pre.Next = nil
		}
	}

	return res
}
