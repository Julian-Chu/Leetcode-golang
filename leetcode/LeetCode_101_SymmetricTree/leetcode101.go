package LeetCode_101_SymmetricTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return compareSymmetric(root.Left, root.Right)
}

func compareSymmetric(Left *TreeNode, Right *TreeNode) bool {
	if Left == nil && Right == nil {
		return true
	}
	if Left == nil || Right == nil {
		return false
	}

	return Left.Val == Right.Val && compareSymmetric(Left.Left, Right.Right) && compareSymmetric(Left.Right, Right.Left)

}

func isSymmetric_DFS(root *TreeNode) bool {
	var dfs func(l, r *TreeNode) bool
	dfs = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}

		if l == nil || r == nil {
			return false
		}

		if l.Val != r.Val {
			return false
		}

		return dfs(l.Left, r.Right) && dfs(l.Right, r.Left)
	}

	return dfs(root.Left, root.Right)
}

func isSymmetric_Stack(root *TreeNode) bool {
	if root == nil {
		return false
	}
	stack := []*TreeNode{root.Right, root.Left}

	for len(stack) > 0 {
		size := len(stack)
		l := stack[size-1]
		r := stack[size-2]
		stack = stack[:size-2]

		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}

		stack = append(stack, l.Left, r.Right)
		stack = append(stack, l.Right, r.Left)
	}
	return true
}

func isSymmetric_queue(root *TreeNode) bool {
	if root == nil {
		return false
	}
	queue := []*TreeNode{root.Right, root.Left}

	for len(queue) > 0 {
		l := queue[0]
		r := queue[1]
		queue = queue[2:]

		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}

		queue = append(queue, l.Left, r.Right)
		queue = append(queue, l.Right, r.Left)
	}
	return true
}
