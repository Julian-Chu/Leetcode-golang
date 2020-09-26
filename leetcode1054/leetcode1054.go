package leetcode1054

import "sort"

func rearrangeBarcodes(barcodes []int) []int {
	n := len(barcodes)
	cnt := [10001]int{}
	for _, b := range barcodes {
		cnt[b]++
	}

	sort.Slice(barcodes, func(i, j int) bool {
		if cnt[barcodes[i]] == cnt[barcodes[j]] {
			return barcodes[i] < barcodes[j]
		}
		return cnt[barcodes[i]] > cnt[barcodes[j]]
	})

	res := make([]int, n)
	i := 0
	for _, b := range barcodes {
		res[i] = b
		i += 2
		if i > n-1 {
			i = 1
		}
	}
	return res
}
