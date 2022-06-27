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

func connect_dfs(root *Node) *Node {
	if root == nil {
		return nil
	}

	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}

		if node.Left != nil {
			node.Left.Next = node.Right
		}

		if node.Right != nil {
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}

		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return root
}

func connect_forloop(root *Node) *Node {
	if root == nil {
		return nil
	}

	for cur := root; cur.Left != nil; cur = cur.Left {
		for node := cur; node != nil; node = node.Next {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}
