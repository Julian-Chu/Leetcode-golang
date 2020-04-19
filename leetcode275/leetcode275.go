package leetcode275

func hIndex(citations []int) int {
	n := len(citations)
	hIdx := 0
	for idx := 0; idx <= n; idx++ {
		gt := 0
		eq := 0
		for i := range citations {
			switch {
			case citations[i] > idx:
				gt++
			case citations[i] == idx:
				eq++
			}
		}
		if gt+eq >= idx {
			hIdx = idx
		}
	}

	return hIdx
}
