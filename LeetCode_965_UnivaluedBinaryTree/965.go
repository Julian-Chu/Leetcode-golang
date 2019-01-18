package leetcode965

func isUnivalTree(root *TreeNode) bool {
	return compareWithRootVal(root, root.Val)
}

func compareWithRootVal(node *TreeNode, val int) bool {
	if node == nil {
		return true
	}
	if node.Val != val {
		return false
	}
	return compareWithRootVal(node.Left, val) && compareWithRootVal(node.Right, val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
