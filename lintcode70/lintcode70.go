package lintcode70

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		next_q := make([]*TreeNode, 0, len(queue))
		level := make([]int, 0, len(queue))
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			level = append(level, node.Val)
			if node.Left != nil {
				next_q = append(next_q, node.Left)
			}

			if node.Right != nil {
				next_q = append(next_q, node.Right)
			}
		}
		queue = next_q
		res = append(res, level)
	}

	i, j := 0, len(res)-1
	for i <= j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res

}
