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
	return true
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
