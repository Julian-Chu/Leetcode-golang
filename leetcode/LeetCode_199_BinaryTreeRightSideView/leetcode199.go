package LeetCode_199_BinaryTreeRightSideView

import "github.com/Julian-Chu/Leetcode-golang/utils"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = utils.TreeNode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	ls := rightSideView(root.Left)
	rs := rightSideView(root.Right)

	if len(ls) > len(rs) {
		rs = append(rs, ls[len(rs):]...)
	}

	res := make([]int, len(rs)+1)
	res[0] = root.Val
	copy(res[1:], rs)
	return res
}

// better solution by BFS
func rightSideView_1(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, queue[size-1].Val)
		queue = queue[size:]
	}

	return res
}
