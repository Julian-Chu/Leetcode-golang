package leetcode310

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	links := make([][]int, n)

	count := make([]int, n)
	for _, edge := range edges {
		links[edge[0]] = append(links[edge[0]], edge[1])
		links[edge[1]] = append(links[edge[1]], edge[0])
		count[edge[0]]++
		count[edge[1]]++
	}

	leafs := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if count[i] == 1 {
			leafs = append(leafs, i)
		}
	}

	for n > 2 {
		newLeafs := make([]int, 0, len(leafs))
		for _, leaf := range leafs {
			n--
			for _, linknode := range links[leaf] {
				count[linknode]--
				if count[linknode] == 1 {
					newLeafs = append(newLeafs, linknode)
				}
			}
		}
		leafs = newLeafs
	}

	res := make([]int, 0, 2)
	for _, root := range leafs {
		res = append(res, root)
	}

	return res
}
