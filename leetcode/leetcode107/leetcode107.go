package leetcode107

import "github.com/Julian-Chu/Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{{root.Val}}
	arr := []*TreeNode{root}

	for len(arr) > 0 {
		size := len(arr)
		var temp []int
		for i := 0; i < size; i++ {
			node := arr[i]
			if node.Left != nil {
				temp = append(temp, node.Left.Val)
				arr = append(arr, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right.Val)
				arr = append(arr, node.Right)
			}
		}
		arr = arr[size:]
		if len(temp) > 0 {
			res = append(res, temp)
		}
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}

func levelOrderBottom_1(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, 0, size)
		for i := 0; i < size; i++ {
			cur := queue[i]
			tmp = append(tmp, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		queue = queue[size:]
		res = append(res, tmp)
	}

	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}
