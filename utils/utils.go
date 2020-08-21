package utils

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
	for i < n && len(queue) > 0 {
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

func ListNodeToInts(node *ListNode) []int {
	res := make([]int, 0)

	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}

	return res
}

func IntsToListNode(ints []int) *ListNode {
	preHead := &ListNode{}

	cur := preHead
	for _, v := range ints {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return preHead.Next
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

// TreeNode is type of treenode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type DoublyLinkedNode struct {
	Val   int
	Prev  *DoublyLinkedNode
	Next  *DoublyLinkedNode
	Child *DoublyLinkedNode
}
