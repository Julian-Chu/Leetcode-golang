package leetcode95

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
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
	if len(in) < 2 {
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
		for j := 0; j < len(left); j++ {
			for k := 0; k < len(right); k++ {
				pre := []int{root}
				pre = append(pre, left[j]...)
				pre = append(pre, right[k]...)
				res = append(res, pre)
			}
		}
	}
	return res
}
