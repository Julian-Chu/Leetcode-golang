package lintcode939

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * @param root: the root node
 * @param k: an integer
 * @return: the number of nodes in the k-th layer of the binary tree
 */
func kthfloorNode(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		k--
		if k == 0 {
			return len(queue)
		}

		next_q := []*TreeNode{}
		for _, node := range queue {
			if node.Left != nil {
				next_q = append(next_q, node.Left)
			}

			if node.Right != nil {
				next_q = append(next_q, node.Right)
			}
		}
		queue = next_q
	}
	return 0
}
