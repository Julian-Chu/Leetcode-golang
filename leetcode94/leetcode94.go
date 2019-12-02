package leetcode94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var l []int
	var r []int
	if root.Left != nil {
		l = inorderTraversal(root.Left)
	}
	if root.Right != nil {
		r = inorderTraversal(root.Right)
	}

	return append(append(l, root.Val), r...)
}
