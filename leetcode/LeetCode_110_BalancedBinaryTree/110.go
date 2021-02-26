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
