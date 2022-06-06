package LeetCode_106_ConstructBinaryTreeFromInorderAndPostorderTraversal

/*
  inorder = [9,3,15,20,7]
postorder = [9,15,7,20,3]
    3
   / \
  9  20
    /  \
   15   7

Definition for a binary tree node.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	mid := 0
	for i := range inorder {
		if inorder[i] == postorder[len(postorder)-1] {
			mid = i
		}
	}

	root.Left = buildTree(inorder[:mid], postorder[:mid])
	root.Right = buildTree(inorder[mid+1:], postorder[mid:len(postorder)-1])

	return root
}
