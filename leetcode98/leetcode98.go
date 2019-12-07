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
	l := true
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			return false
		}
		l = isValidLeftSubBST(root.Left, root.Val)
	}

	r := true
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			return false
		}
		r = isValidRightSubBST(root.Right, root.Val)
	}

	return r && l
}

func isValidSubBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	l := true
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			l = false
		} else {
			l = isValidLeftSubBST(root.Left, max)
		}
	}
	r := true
	if root.Right != nil {
		if root.Right.Val <= root.Val || root.Right.Val >= max {
			r = false
		} else {
			r = isValidRightSubBST(root.Right, root.Right.Val)
		}
	}
	return l && r
}
func isValidLeftSubBST(root *TreeNode, max int) bool {
	if root == nil {
		return true
	}

	l := true
	if root.Left != nil {
		if root.Left.Val >= root.Val {
			l = false
		} else {
			l = isValidLeftSubBST(root.Left, max)
		}
	}
	r := true
	if root.Right != nil {
		if root.Right.Val <= root.Val || root.Right.Val >= max {
			r = false
		} else {
			r = isValidRightSubBST(root.Right, root.Right.Val)
		}
	}
	return l && r
}

func isValidRightSubBST(root *TreeNode, min int) bool {
	if root == nil {
		return true
	}

	l := true
	if root.Left != nil {
		if root.Left.Val >= root.Val || root.Left.Val <= min {
			l = false
		} else {
			l = isValidLeftSubBST(root.Left, root.Left.Val)
		}
	}
	r := true
	if root.Right != nil {
		if root.Right.Val <= root.Val {
			r = false
		} else {
			r = isValidRightSubBST(root.Right, min)
		}
	}
	return l && r
}
