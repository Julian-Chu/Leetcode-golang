package leetcode110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, res := getHeightAndIsBalanced(root)

	return res
}

func getHeightAndIsBalanced(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftH, leftB := getHeightAndIsBalanced(root.Left)
	rightH, rightB := getHeightAndIsBalanced(root.Right)

	h := leftH + 1
	diff := leftH - rightH
	if rightH > leftH {
		h = rightH + 1
		diff = rightH - leftH
	}

	return h, diff <= 1 && leftB && rightB

}

func isBalanced_2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := maxdepth(root.Left)
	r := maxdepth(root.Right)

	if l-r > 1 || l-r < -1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)

}

func maxdepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(maxdepth(node.Left), maxdepth(node.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
