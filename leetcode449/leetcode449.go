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

const NULL string = "NULL"

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := ""
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		q_next := make([]*TreeNode, 0)
		for i := range q {
			if q[i] != nil {
				res += strconv.Itoa(q[i].Val)
				res += ","
				if q[i].Left != nil {
					q_next = append(q_next, q[i].Left)
				} else {
					q_next = append(q_next, nil)
				}

				if q[i].Right != nil {
					q_next = append(q_next, q[i].Right)
				} else {
					q_next = append(q_next, nil)
				}
			} else {
				res += NULL
				res += ","
				q_next = append(q_next, nil)
				q_next = append(q_next, nil)
			}
		}
		q = nil
		for i := range q_next {
			if q_next[i] != nil {
				q = q_next
				break
			}
		}
	}

	if len(res) > 0 {
		return res[:len(res)-1]
	}

	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	nodes := strings.Split(data, ",")

	l := len(nodes)
	batchsize := 1
	var root *TreeNode
	start := 0
	var q []*TreeNode
	for batchsize <= l {
		var q_next []*TreeNode
		ns := nodes[start : start+batchsize]
		if len(q) == 0 {
			num, _ := strconv.Atoi(ns[0])
			root = &TreeNode{Val: num}
			q_next = append(q_next, root)
		} else {
			for i := range q {
				if q[i] == nil {
					q_next = append(q_next, nil)
					q_next = append(q_next, nil)
					continue
				}
				if ns[i*2] != NULL {
					num, _ := strconv.Atoi(ns[i*2])
					node := &TreeNode{Val: num}
					q[i].Left = node
					q_next = append(q_next, node)
				} else {
					q_next = append(q_next, nil)
				}
				if ns[i*2+1] != NULL {
					num, _ := strconv.Atoi(ns[i*2+1])
					node := &TreeNode{Val: num}
					q[i].Right = node
					q_next = append(q_next, node)
				} else {
					q_next = append(q_next, nil)
				}
			}
		}
		start += batchsize
		batchsize *= 2
		q = q_next
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
