package LeetCode_559_MaximumDepthofN_aryTree

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	if root.Children == nil {
		return 1
	}

	max := 0
	for i := 0; i < len(root.Children); i++ {
		dep := maxDepth(root.Children[i])
		if dep > max {
			max = dep
		}
	}

	return max + 1
}
