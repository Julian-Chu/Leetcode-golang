package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}

	traverse(root)
	return res
}

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		traverse(node.Right)
		res = append(res, node.Val)
	}

	traverse(root)
	return res
}
