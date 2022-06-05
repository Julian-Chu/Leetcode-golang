package LeetCode_429_N_aryTreeLevelOrderTraversal

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, 0, size)
		for i := 0; i < size; i++ {
			cur := queue[i]
			if cur.Children != nil {
				queue = append(queue, cur.Children...)
			}
			tmp = append(tmp, cur.Val)
		}
		res = append(res, tmp)
		queue = queue[size:]
	}
	return res
}
