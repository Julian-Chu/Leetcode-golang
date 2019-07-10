package leetcode617

import (
	. "Leetcode-golang/helper"
	"fmt"
)

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	fmt.Println(t1)
	return traverse(t1, t2)
}

func traverse(t1 *TreeNode, t2 *TreeNode) *TreeNode {
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
