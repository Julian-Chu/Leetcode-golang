package leetcode563

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	tilt, _ := calcTiltAndSum(root)

	return tilt
}

func calcTiltAndSum(root *TreeNode) (tilt, sum int) {
	var lt, rt, ls, rs int
	if root == nil {
		return 0, 0
	}
	if root.Left == nil && root.Right == nil {
		return 0, root.Val
	}

	if root.Left != nil {
		lt, ls = calcTiltAndSum(root.Left)
	}

	if root.Right != nil {
		rt, rs = calcTiltAndSum(root.Right)
	}

	return (lt + rt + int(math.Abs(float64(ls-rs)))), ls + rs + root.Val

}
