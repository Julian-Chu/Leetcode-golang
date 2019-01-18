package leetcode563

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	tilt := 0
	var left *TreeNode
	var right *TreeNode
	if root == nil {
		return 0
	}
	if root.Left != nil {
		tilt += findTilt(root.Left)
		left = root.Left
	}
	if root.Right != nil {
		tilt += findTilt(root.Right)
		right = root.Right
	}

	tilt += int(math.Abs(float64(sumOfNode(left) - sumOfNode(right))))
	return tilt
}

func sumOfNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := root.Val

	if root.Left != nil {
		sum += sumOfNode(root.Left)
	}
	if root.Right != nil {
		sum += sumOfNode(root.Right)
	}
	return sum
}
