package leetcode653

import (
	"os"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	if root.Left == nil && root.Right == nil {
		return false
	}

	return (root.Val*2 != k && findNode(root, k-root.Val)) || (root.Left != nil && iterateNode(root, root.Left, k)) || (root.Right != nil && iterateNode(root, root.Right, k))
}

func iterateNode(root *TreeNode, newRoot *TreeNode, k int) bool {
	return (newRoot.Val*2 != k && findNode(root, k-newRoot.Val)) || (newRoot.Left != nil && iterateNode(root, newRoot.Left, k)) || (newRoot.Right != nil && iterateNode(root, newRoot.Right, k))
}

func findNode(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	if root.Val == target {
		return true
	}
	if root.Val < target {
		return findNode(root.Right, target)
	}
	return findNode(root.Left, target)
}

var root *TreeNode

func TestMain(m *testing.M) {
	root = &TreeNode{
		5,
		&TreeNode{
			3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil},
		},
		&TreeNode{
			6, nil, &TreeNode{7, nil, nil},
		},
	}
	os.Exit(m.Run())
}

func Test_Given8_ReturnTrue(t *testing.T) {
	res := findTarget(root, 8)
	if res == false {
		t.Error("Not Found")
	}
}

func Test_Given28_ReturnFalse(t *testing.T) {
	res := findTarget(root, 28)
	if res == true {
		t.Error("Not Found")
	}
}

func Test_Given11_ReturnTrue(t *testing.T) {
	res := findTarget(root, 11)
	if res == false {
		t.Error("Not Found")
	}
}

func Test_Given7_ReturnTrue(t *testing.T) {
	res := findTarget(root, 7)
	if res == false {
		t.Error("Not Found")
	}
}

func Test_Given1_OnlyOneNode(t *testing.T) {
	root = &TreeNode{1, nil, nil}
	res := findTarget(root, 1)
	if res == true {
		t.Error("only one node")
	}
}

func Test_Given4_ValsInSubNodes(t *testing.T) {
	root1 := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	res := findTarget(root1, 4)
	if res == false {
		t.Error("Should be found!")
	}
}
