package leetcode687

func longestUnivaluePath_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	substr := max(longestUnivaluePath(root.Left), longestUnivaluePath(root.Right))
	return max(substr, findPathLen(root.Left, root.Val)+findPathLen(root.Right, root.Val))
}

func findPathLen(root *TreeNode, val int) int {
	if root == nil || root.Val != val {
		return 0
	}

	return 1 + max(findPathLen(root.Left, val), findPathLen(root.Right, val))
}
