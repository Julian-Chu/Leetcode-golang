package leetcode95

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {

	in := make([]int, n)
	for i := 0; i < n; i++ {
		in[i] = i + 1
	}

	pres := getPres(in)

	return buildTreeNodes(in, pres)
}

func buildTreeNodes(in []int, pres [][]int) []*TreeNode {
	nodes := make([]*TreeNode, len(pres))

	for i, pre := range pres {
		nodes[i] = buildNode(in, pre)
	}
	return nodes
}

func buildNode(in []int, pre []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	node := &TreeNode{Val: pre[0]}
	nodeInIdx := 0
	for i := range in {
		if in[i] == pre[0] {
			nodeInIdx = i
			break
		}
	}
	node.Left = buildNode(in[0:nodeInIdx], pre[1:nodeInIdx+1])
	node.Right = buildNode(in[nodeInIdx+1:], pre[nodeInIdx+1:])
	return node
}

func getPres(in []int) [][]int {
	if len(in) == 0 {
		return [][]int{}
	}
	if len(in) == 1 {
		return [][]int{in}
	}
	if len(in) == 2 {
		return [][]int{{in[0], in[1]}, {in[1], in[0]}}
	}
	res := make([][]int, 0)
	for i := 0; i < len(in); i++ {
		root := in[i]
		left := getPres(in[:i])
		right := getPres(in[i+1:])
		n1 := len(left)
		if n1 == 0 {
			n1 = 1
		}

		for j := 0; j < n1; j++ {
			n2 := len(right)
			if n2 == 0 {
				n2 = 1
			}
			for k := 0; k < n2; k++ {
				pre := append([]int{}, root)
				if len(left) != 0 {
					pre = append(pre, left[j]...)
				}
				if len(right) != 0 {
					pre = append(pre, right[k]...)
				}
				res = append(res, pre)
			}
		}
	}
	return res
}
