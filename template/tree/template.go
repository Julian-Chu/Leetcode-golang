package tree

func inorderTraversal_template(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}

			stack = append(stack, node)
			stack = append(stack, nil)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
		}
	}
	return res
}

func preorderTraversal_template(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			if node.Right != nil {
				stack = append(stack, node.Right)
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			stack = append(stack, node)
			stack = append(stack, nil)
		} else {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
		}
	}
	return res
}

func postorderTraversal_template(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			stack = append(stack, node)
			stack = append(stack, nil)

			if node.Right != nil {
				stack = append(stack, node.Right)
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
		}
	}
	return res
}
