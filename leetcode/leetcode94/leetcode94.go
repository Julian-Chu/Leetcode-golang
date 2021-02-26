package leetcode94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		if len(stack) > 0 {
			node := stack[len(stack)-1]
			res = append(res, node.Val)
			stack = stack[:len(stack)-1]
			cur = node.Right
		}
	}

	return res
}
