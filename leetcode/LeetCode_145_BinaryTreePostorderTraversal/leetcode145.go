package LeetCode_145_BinaryTreePostorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

// follow same logic like preorder
// preorder:  middle -> left -> right
// change preorder: middle -> right -> left => reverse, then get postorder
// postorder: left-> right -> middle
func postorderTraversal_stack(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)

	if root == nil {
		return res
	}

	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res

}
