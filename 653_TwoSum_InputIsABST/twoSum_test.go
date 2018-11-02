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

type HashSet struct {
	set map[int]bool
}

func NewHashSet() *HashSet {
	return &HashSet{make(map[int]bool)}
}

func (set *HashSet) Add(i int) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found
}

func (set *HashSet) Get(i int) bool {
	_, found := set.set[i]
	return found
}

func (set *HashSet) Remove(i int) {
	delete(set.set, i)
}

func findTarget(root *TreeNode, k int) bool {
	tree := NewHashSet()

	tree.Add(root.Val)
	if root.Left != nil {
		AddNodeToHashSet(tree, root.Left)
	}

	if root.Right != nil {
		AddNodeToHashSet(tree, root.Right)
	}

	for key := range tree.set {
		target := k - key
		if target == key {
			continue
		}
		if tree.Get(target) == true {
			return true
		}
	}

	return false
}

func AddNodeToHashSet(set *HashSet, newRoot *TreeNode) {
	set.Add(newRoot.Val)
	if newRoot.Left != nil {
		AddNodeToHashSet(set, newRoot.Left)
	}
	if newRoot.Right != nil {
		AddNodeToHashSet(set, newRoot.Right)
	}
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
	root = &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	res := findTarget(root, 4)
	if res == false {
		t.Error("Should be found!")
	}
}
