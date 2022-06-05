package leetcode102

import "github.com/Julian-Chu/Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func levelOrder(root *TreeNode) [][]int {
	level := 0
	res := make([][]int, 0)
	BuildLevelOrderTraversal(root, level, &res)
	return res
}

func BuildLevelOrderTraversal(node *TreeNode, level int, res *[][]int) {
	if node == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, []int{})
	}

	(*res)[level] = append((*res)[level], node.Val)

	BuildLevelOrderTraversal(node.Left, level+1, res)
	BuildLevelOrderTraversal(node.Right, level+1, res)
}

// BFS, better solution
func levelOrder_1(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		size := len(queue)
		nodes := make([]int, 0, size)
		for i := 0; i < size; i++ {
			cur := queue[i]
			nodes = append(nodes, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, nodes)
		queue = queue[size:]
	}

	return res
}
