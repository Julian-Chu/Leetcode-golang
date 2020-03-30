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

	return this.buildTree(nodes)
}

func (this *Codec) buildTree(nodes []string) *TreeNode {
	leftIdx := 0
	rightIdx := 0

	rootVal := nodes[0]
	for i := 1; i < len(nodes); i++ {
		if leftIdx == 0 && this.compareValInString(nodes[i], rootVal) < 0 {
			leftIdx = i
		}
		if rightIdx == 0 && this.compareValInString(nodes[i], rootVal) > 0 {
			rightIdx = i
		}
	}

	v, _ := strconv.Atoi(rootVal)
	root := &TreeNode{Val: v}
	if rightIdx != 0 {
		root.Right = this.buildTree(nodes[rightIdx:])
	}
	if leftIdx != 0 {
		if rightIdx == 0 {
			rightIdx = len(nodes)
		}
		root.Left = this.buildTree(nodes[leftIdx:rightIdx])
	}

	return root
}

func (this Codec) compareValInString(str1, str2 string) int {
	switch {
	case len(str1) > len(str2):
		return 1
	case len(str1) < len(str2):
		return -1
	}

	for i := range str1 {
		switch {
		case str1[i] > str2[i]:
			return 1
		case str1[i] < str2[i]:
			return -1
		}
	}
	return 0
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
