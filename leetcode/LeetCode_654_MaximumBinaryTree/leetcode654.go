package LeetCode_654_MaximumBinaryTree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	rootIdx := findMaxAtIdx(nums)
	root := &TreeNode{Val: nums[rootIdx]}
	root.Left = constructMaximumBinaryTree(nums[:rootIdx])
	root.Right = constructMaximumBinaryTree(nums[rootIdx+1:])
	return root
}

func findMaxAtIdx(nums []int) int {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[idx] {
			idx = i
		}
	}

	return idx
}
