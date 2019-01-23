package leetcode110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	
}

var levels = map[int]bool
var moveCount = 0
func traverse(root *TreeNode ){
	if root.Left ==nil && root.Right == nil{
		if levels[moveCount] == false{
			levels[moveCount] == true
		}
	}
	if root.Left !=nil{
		traverse(root.Left)
	}

	if root.Right !=nil{
		traverse(root.Right)
	}
}


