package leetcode105

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	mid := 0
	for i := range inorder {
		if inorder[i] == preorder[0] {
			mid = i
			break
		}
	}

	build(root, preorder[1:], inorder[:mid], inorder[mid+1:])
	return root
}

func build(root *TreeNode, preorder, inorderLeft, inorderRight []int) {
	if len(inorderLeft) == 0 && len(inorderRight) == 0 {
		return
	}
	for i := range inorderLeft {
		if preorder[0] == inorderLeft[i] {
			root.Left = &TreeNode{
				Val: preorder[0],
			}
			build(root.Left, preorder[1:], inorderLeft[:i], inorderRight[i+1:])
		}
	}
	for i := range inorderRight {
		if preorder[1] == inorderRight[i] {
			root.Right = &TreeNode{
				Val: preorder[1],
			}
			build(root.Left, preorder[1:], inorderLeft[:i], inorderRight[i+1:])
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
