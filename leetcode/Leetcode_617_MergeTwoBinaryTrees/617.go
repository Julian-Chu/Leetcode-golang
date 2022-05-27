package leetcode617

import (
	"fmt"

	"github.com/Julian-Chu/Leetcode-golang/utils"
)

func mergeTrees(t1 *utils.TreeNode, t2 *utils.TreeNode) *utils.TreeNode {
	fmt.Println(t1)
	return traverse(t1, t2)
}

func traverse(t1 *utils.TreeNode, t2 *utils.TreeNode) *utils.TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil && t2 != nil {
		return t2
	}

	if t1 != nil && t2 == nil {
		return t1
	}

	if t1 != nil && t2 != nil {
		t1.Val += t2.Val
		t1.Left = traverse(t1.Left, t2.Left)
		t1.Right = traverse(t1.Right, t2.Right)
	}
	return t1

}
