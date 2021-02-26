package leetcode872

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1 := getLeafs(root1, []int{})
	leafs2 := getLeafs(root2, []int{})

	for i, v := range leafs1 {
		if v != leafs2[i] {
			return false
		}
	}

	return true
}

func getLeafs(root *TreeNode, leafs []int) []int {
	if root.Left == nil && root.Right == nil {
		leafs = append(leafs, root.Val)
	}

	if root.Left != nil {
		leafs = getLeafs(root.Left, leafs)
	}

	if root.Right != nil {
		leafs = getLeafs(root.Right, leafs)
	}

	return leafs

}
