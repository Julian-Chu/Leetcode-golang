package LeetCode_637_AverageOfLevelsInBinaryTree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		sum := 0

		for i := 0; i < size; i++ {
			cur := queue[i]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			sum += cur.Val
		}
		queue = queue[size:]
		res = append(res, float64(sum)/float64(size))
	}

	return res
}
