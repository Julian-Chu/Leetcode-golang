package leetcode98

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := true
	if root.Left != nil {
		if root.Val <= root.Left.Val {
			return false
		}
		left = isLeftValidSubBST(root.Left, root.Val)
	}

	right := true
	if root.Right != nil {
		if root.Val >= root.Right.Val {
			return false
		}
		right = isRightValidSubBST(root.Right, root.Val)
	}

	return left && right

}

func isLeftValidSubBST(node *TreeNode, max int) bool {
	if node == nil {
		return true
	}
	left := true
	if node.Left != nil {
		if node.Val <= node.Left.Val || node.Left.Val >= max {
			return false
		}
		left = isLeftValidSubBST(node.Left, max)
	}

	right := true
	if node.Right != nil {
		if node.Val >= node.Right.Val || node.Right.Val >= max {
			return false
		}
		right = isRightValidSubBST(node.Right, max)
	}

	return left && right
}
func isRightValidSubBST(node *TreeNode, min int) bool {
	if node == nil {
		return true
	}
	left := true
	if node.Left != nil {
		if node.Val <= node.Left.Val || node.Left.Val <= min {
			return false
		}
		left = isLeftValidSubBST(node.Left, min)
	}

	right := true
	if node.Right != nil {
		if node.Val >= node.Right.Val || node.Right.Val <= min {
			return false
		}
		right = isRightValidSubBST(node.Right, min)
	}

	return left && right
}
