package BST

func inOrderTraverse(root *TreeNode) {
	var prev *TreeNode
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		// start from Left
		traverse(node.Left)

		// handle logic with mid
		if prev != nil {
			//.....
		}
		prev = node

		// go to right
		traverse(node.Right)
	}
	traverse(root)
}
