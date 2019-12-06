package leetcode107

import "Leetcode-golang/utils"

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
