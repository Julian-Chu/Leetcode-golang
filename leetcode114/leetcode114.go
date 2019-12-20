package leetcode114

import "Leetcode-golang/utils"

type TreeNode = utils.TreeNode

func flatten(root *TreeNode) {
	recur(root)
}

func recur(node *TreeNode) *TreeNode {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return node
	}
	if node.Left == nil {
		return recur(node.Right)
	}
	if node.Right == nil {
		node.Right = node.Left
		node.Left = nil
		return recur(node.Right)
	}

	end := recur(node.Right)
	recur(node.Left).Right = node.Right
	node.Right = node.Left
	node.Left = nil
	return end
	//var right *TreeNode
	//if node.Left != nil {
	//	if node.Right != nil {
	//		right = node.Right
	//		node.Right = node.Left
	//		recur(node.Left).Right = right
	//		//recur(node.Left).Right = node.Right
	//		//node.Right = node.Left
	//		node.Left = nil
	//		return recur(right)
	//	} else {
	//		node.Right = node.Left
	//		node.Left = nil
	//		return recur(node.Right)
	//	}
	//}
	//if node.Right != nil {
	//	return recur(node.Right)
	//}
	//return node
}
