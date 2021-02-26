package lintcode69

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * @param root: A Tree
 * @return: Level order a list of lists of integer
 */
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		next_queue := make([]*TreeNode, 0)
		levelOrderTraversal := []int{}
		for _, node := range queue {
			levelOrderTraversal = append(levelOrderTraversal, node.Val)
			if node.Left != nil {
				next_queue = append(next_queue, node.Left)
			}
			if node.Right != nil {
				next_queue = append(next_queue, node.Right)
			}
		}
		res = append(res, levelOrderTraversal)
		queue = next_queue
	}

	return res
}
