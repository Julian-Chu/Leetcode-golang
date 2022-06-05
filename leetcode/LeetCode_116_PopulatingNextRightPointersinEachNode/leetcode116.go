package LeetCode_116_PopulatingNextRightPointersinEachNode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		size := len(queue)
		var prev *Node
		for i := 0; i < size; i++ {
			cur := queue[i]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			if prev != nil {
				prev.Next = cur
			}
			prev = cur

		}

		queue = queue[size:]
	}

	return root
}
