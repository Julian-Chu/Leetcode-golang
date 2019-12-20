package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NULL = -1 << 63

func Ints2Tree(ints []int) (root *TreeNode) {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root = &TreeNode{
		Val: ints[0],
	}
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root
	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]
		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func compareTrees(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	switch {
	case node1 == nil && node2 != nil:
	case node1 != nil && node2 == nil:
		return false
	}
	if node1.Val != node2.Val {
		return false
	}

	return compareTrees(node1.Left, node2.Left) && compareTrees(node1.Right, node2.Right)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
