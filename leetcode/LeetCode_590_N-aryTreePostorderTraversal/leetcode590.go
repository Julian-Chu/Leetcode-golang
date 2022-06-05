package LeetCode_590_N_aryTreePostorderTraversal

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var res []int
	var traverse func(*Node)
	traverse = func(node *Node) {
		if node == nil {
			return
		}

		if node.Children != nil {
			for i := 0; i < len(node.Children); i++ {
				traverse(node.Children[i])
			}
		}

		res = append(res, node.Val)
	}

	traverse(root)
	return res
}
