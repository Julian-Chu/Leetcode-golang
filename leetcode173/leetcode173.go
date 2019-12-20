package leetcode173

import "Leetcode-golang/utils"

type TreeNode = utils.TreeNode
type BSTIterator struct {
	vals   []int
	curIdx int
}

func Constructor(root *TreeNode) BSTIterator {
	var (
		res   []int
		recur func(*TreeNode)
	)
	recur = func(node *TreeNode) {
		if node == nil {
			return
		}
		recur(node.Left)
		res = append(res, node.Val)
		recur(node.Right)
	}
	recur(root)
	return BSTIterator{vals: res, curIdx: -1}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	this.curIdx++
	return this.vals[this.curIdx]
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.curIdx < len(this.vals)-1
}
