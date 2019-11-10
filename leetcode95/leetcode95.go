package leetcode95

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return helper(1, n+1, map[string][]*TreeNode{})
}

func helper(start int, end int, visited map[string][]*TreeNode) []*TreeNode {
	key := fmt.Sprintf("%d-%d", start, end)
	if _, ok := visited[key]; ok {
		return visited[key]
	}

	if start == end {
		return []*TreeNode{nil}
	}

	var result []*TreeNode
	for i := start; i < end; i++ {
		left := helper(start, i, visited)
		right := helper(i+1, end, visited)
		for l := 0; l < len(left); l++ {
			for r := 0; r < len(right); r++ {
				root := &TreeNode{
					Val:   i,
					Left:  left[l],
					Right: right[r],
				}
				result = append(result, root)
			}
		}
	}
	visited[key] = result
	return result
}
