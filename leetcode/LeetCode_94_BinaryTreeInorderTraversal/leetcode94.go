package LeetCode_94_BinaryTreeInorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal_stack(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		if len(stack) > 0 {
			node := stack[len(stack)-1]
			res = append(res, node.Val)
			stack = stack[:len(stack)-1]
			cur = node.Right
		}
	}

	return res
}

func inorderTraversal_stack2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)

	cur := root

	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			cur = node.Right
		}
	}
	return res
}
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)
		res = append(res, node.Val)
		traverse(node.Right)
	}

	traverse(root)
	return res
}
