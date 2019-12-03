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
		if root.Val <= root.Left.Val {
			return false
		}
		l = isValidSubBST(root.Left, nil, root)
	}
	r := true
	if root.Right != nil {
		if root.Val >= root.Right.Val {
			return false
		}
		r = isValidSubBST(root.Right, root, nil)
	}

	return l && r
}

func isValidSubBST(node *TreeNode, minNode *TreeNode, maxNode *TreeNode) bool {
	l := true
	if node.Left != nil {
		if node.Val <= node.Left.Val {
			return false
		}
		if minNode != nil && node.Left.Val <= minNode.Val {
			return false
		}
		if maxNode != nil && node.Left.Val >= maxNode.Val {
			return false
		}
		if minNode != nil {
			l = isValidSubBST(node.Left, minNode, node)
		} else {
			l = isValidSubBST(node.Left, nil, node)
		}
	}
	r := true
	if node.Right != nil {
		if node.Val >= node.Right.Val {
			return false
		}

		if minNode != nil && node.Right.Val <= minNode.Val {
			return false
		}
		if maxNode != nil && node.Right.Val >= maxNode.Val {
			return false
		}
		if maxNode != nil {
			r = isValidSubBST(node.Right, node, maxNode)
		} else {
			r = isValidSubBST(node.Right, node, nil)
		}
	}

	return l && r
}
