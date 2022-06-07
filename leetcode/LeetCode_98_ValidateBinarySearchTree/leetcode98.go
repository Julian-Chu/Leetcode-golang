package LeetCode_98_ValidateBinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(*TreeNode, int64, int64) bool
	dfs = func(node *TreeNode, upper, lower int64) bool {
		if node == nil {
			return true
		}
		val := int64(node.Val)
		if val >= upper || val <= lower {
			return false
		}

		return dfs(node.Left, val, lower) && dfs(node.Right, upper, val)
	}

	return dfs(root, 1<<31, -1<<31-1)
}

// inorder traverse -> slice -> compare
func isValidBST_traverse(root *TreeNode) bool {
	arr := make([]int, 0)

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		arr = append(arr, node.Val)
		traverse(node.Right)
	}
	traverse(root)

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			continue
		}

		if arr[i] <= arr[i-1] {
			return false
		}
	}
	return true
}

func isValidBST_inorder_traverse(root *TreeNode) bool {
	var prev *TreeNode

	var traverse func(node *TreeNode) bool
	traverse = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		left := traverse(node.Left)
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node
		right := traverse(node.Right)
		return left && right
	}
	return traverse(root)
}
