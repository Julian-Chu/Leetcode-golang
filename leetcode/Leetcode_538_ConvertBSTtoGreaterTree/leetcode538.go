package Leetcode_538_ConvertBSTtoGreaterTree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	traverse(root, &sum)
	return root
}

func traverse(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	traverse(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	traverse(root.Left, sum)

}

func convertBST_recur(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum := 0
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Right)
		sum += node.Val
		node.Val = sum
		traverse(node.Left)
	}
	traverse(root)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST_iter(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	sum := 0
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node != nil {
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		} else {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			sum += cur.Val
			cur.Val = sum
		}
	}
	return root
}
