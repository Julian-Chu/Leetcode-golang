package lintcode71

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func zigzagLevelOrder(root *TreeNode) [][]int {
//	res := make([][]int, 0)
//	if root == nil {
//		return res
//	}
//
//	queue := []*TreeNode{root}
//	level := 0
//	for len(queue) > 0 {
//		level++
//		var next_q []*TreeNode
//		levelRes := make([]int, 0, len(queue))
//		for i := range queue {
//			node := queue[i]
//			levelRes = append(levelRes, node.Val)
//			if node.Left != nil {
//				next_q = append(next_q, node.Left)
//			}
//			if node.Right != nil {
//				next_q = append(next_q, node.Right)
//			}
//		}
//		if level&1 == 0 {
//			i, j := 0, len(levelRes)-1
//			for i <= j {
//				levelRes[i], levelRes[j] = levelRes[j], levelRes[i]
//				i++
//				j--
//			}
//		}
//		res = append(res, levelRes)
//		queue = next_q
//	}
//	return res
//}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	backwards := true
	for len(queue) > 0 {
		backwards = !backwards
		level := make([]int, 0, len(queue))
		for range queue {
			node := queue[0]
			copy(queue, queue[1:])
			queue = queue[:len(queue)-1]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if backwards {
			i, j := 0, len(level)-1
			for i <= j {
				level[i], level[j] = level[j], level[i]
				i++
				j--
			}
		}
		res = append(res, level)
	}
	return res
}
