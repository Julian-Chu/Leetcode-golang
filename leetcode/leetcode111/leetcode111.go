package leetcode111

import "github.com/Julian-Chu/Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if root.Left == nil || root.Right == nil {
		if l > r {
			return l + 1
		}
		return r + 1
	}

	if l > r {
		return r + 1
	}
	return l + 1
}
