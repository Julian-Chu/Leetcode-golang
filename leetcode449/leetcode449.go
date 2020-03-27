package leetcode449

import (
	"Leetcode-golang/utils"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = utils.TreeNode
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	res := strconv.Itoa(root.Val)
	res += ","
	if root.Left != nil {
		res += this.serialize(root.Left)
	}

	if root.Right != nil {
		res += this.serialize(root.Right)
	}
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	if data[len(data)-1] == ',' {
		data = data[:len(data)-1]
	}
	nodes := strings.Split(data, ",")
	leftIdx := 0
	rightIdx := 0

	rootVal, _ := strconv.Atoi(nodes[0])
	for i := 1; i < len(nodes); i++ {
		num, _ := strconv.Atoi(nodes[i])
		if leftIdx == 0 && num < rootVal {
			leftIdx = i
		}
		if rightIdx == 0 && num > rootVal {
			rightIdx = i
		}
	}

	num, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{Val: num}
	if rightIdx != 0 {
		root.Right = this.deserialize(strings.Join(nodes[rightIdx:], ","))
	}
	if leftIdx != 0 {
		if rightIdx == 0 {
			rightIdx = len(nodes)
		}
		root.Left = this.deserialize(strings.Join(nodes[leftIdx:rightIdx], ","))
	}

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
