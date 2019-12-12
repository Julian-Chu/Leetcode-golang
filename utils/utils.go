package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Ints2Tree(arr []int) (root *TreeNode) {
	return root
}

func compareTrees(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	switch {
	case node1 == nil && node2 != nil:
	case node1 != nil && node2 == nil:
		return false
	}
	if node1.Val != node2.Val {
		return false
	}

	return compareTrees(node1.Left, node2.Left) && compareTrees(node1.Right, node2.Right)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
