package leetcode530

import "math"

func getMinimumDifferenceStack(root *TreeNode) int {
	stack := make([]*TreeNode, 0)
	prev := (*TreeNode)(nil)
	ans := math.MaxInt32

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[: len(stack)-1 : len(stack)-1]
		if prev != nil {
			ans = min(root.Val-prev.Val, ans)
		}
		prev = root
		root = root.Right
	}

	return ans
}
