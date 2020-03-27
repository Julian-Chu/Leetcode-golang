package leetcode257

import (
	"Leetcode-golang/utils"
	"strconv"
)

type TreeNode = utils.TreeNode

func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	if root == nil {
		return res
	}

	var findPath func(*TreeNode, string)
	findPath = func(node *TreeNode, path string) {
		path = path + strconv.Itoa(node.Val)
		if node.Right == nil && node.Left == nil {
			res = append(res, path)
		}

		if node.Left != nil {
			findPath(node.Left, path+"->")
		}

		if node.Right != nil {
			findPath(node.Right, path+"->")
		}
	}

	findPath(root, "")

	return res
}
