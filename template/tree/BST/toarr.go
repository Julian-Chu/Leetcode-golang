package BST

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func toarr(root *TreeNode) []int {
	var nums []int

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		nums = append(nums, node.Val)
		traverse(node.Right)
	}
	traverse(root)

	return nums
}
