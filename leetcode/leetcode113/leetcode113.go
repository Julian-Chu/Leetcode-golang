package leetcode113

import "github.com/Julian-Chu/Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	res = recur(root, sum, res, []int{})
	return res
}

func recur(node *TreeNode, sum int, res [][]int, path []int) [][]int {
	if node == nil {
		return res
	}
	n := len(path)
	path = append(path[:n:n], node.Val)
	if node.Left == nil && node.Right == nil {
		if node.Val == sum {
			res = append(res, path)
		}
		return res
	}
	sum -= node.Val
	if node.Left != nil {
		res = recur(node.Left, sum, res, path)
	}
	if node.Right != nil {
		res = recur(node.Right, sum, res, path)
	}
	return res
}

func pathSum_DFS(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var dfs func(*TreeNode, int, []int)
	dfs = func(node *TreeNode, sum int, path []int) {
		if node == nil {
			return
		}

		sum += node.Val
		path = append(path[:], node.Val)
		if node.Left == nil && node.Right == nil {
			if sum == targetSum {
				// be careful to create new slice here, to avoid data change by append
				path = append([]int{}, path...)
				res = append(res, path)
			}
			return
		}

		dfs(node.Left, sum, path)
		dfs(node.Right, sum, path)
	}
	dfs(root, 0, []int{})
	return res
}
