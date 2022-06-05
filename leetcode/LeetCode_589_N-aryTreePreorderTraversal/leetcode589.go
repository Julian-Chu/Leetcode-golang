package LeetCode_589_N_aryTreePreorderTraversal

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var res []int
	var traverse func(node *Node)
	traverse = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)

		if node.Children != nil {
			for i := 0; i < len(node.Children); i++ {
				traverse(node.Children[i])
			}
		}
	}
	traverse(root)
	return res
}
