package Leetcode_543_DiameterOfBinaryTree

import (
	"fmt"
	. "leetcode-golang/helper"
)

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxPath := 0
	getLongestPath(root, &maxPath)
	return maxPath
}

func getLongestPath(root *TreeNode, maxPath *int) int {
	if root == nil {
		return 0
	}

	lh := getLongestPath(root.Left, maxPath)
	rh := getLongestPath(root.Right, maxPath)

	if root.Val == 2 {
		fmt.Println(lh, rh)
	}

	if lh+rh > *maxPath {
		*maxPath = lh + rh
		fmt.Println(*maxPath)

	}

	if lh > rh {
		return lh + 1
	}
	return rh + 1

}
