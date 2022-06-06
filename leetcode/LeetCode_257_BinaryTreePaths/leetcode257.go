package LeetCode_257_BinaryTreePaths

import (
	"strconv"

	"github.com/Julian-Chu/Leetcode-golang/utils"
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

func binaryTreePaths_stqck(root *TreeNode) []string {
	stack := []*TreeNode{}
	paths := make([]string, 0)
	res := make([]string, 0)
	if root == nil {
		return res
	}
	stack = append(stack, root)
	paths = append(paths, "")

	for len(stack) > 0 {
		size := len(stack)
		node := stack[size-1]
		path := paths[size-1]
		stack = stack[:size-1]
		paths = paths[:size-1]
		if node.Left == nil && node.Right == nil {
			res = append(res, path+strconv.Itoa(node.Val))
			continue
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			paths = append(paths, path+strconv.Itoa(node.Val)+"->")
		}
	}
	return res
}
