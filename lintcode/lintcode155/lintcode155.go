package lintcode155

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

/**
 * @param root: The root of binary tree
 * @return: An integer
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}

	step := 0
	for len(queue) > 0 {
		next_q := []*TreeNode{}
		step++
		for _, node := range queue {
			if node.Left == nil && node.Right == nil {
				return step
			}
			if node.Left != nil {
				next_q = append(next_q, node.Left)
			}
			if node.Right != nil {
				next_q = append(next_q, node.Right)
			}
		}
		queue = next_q
	}
	return step

}
