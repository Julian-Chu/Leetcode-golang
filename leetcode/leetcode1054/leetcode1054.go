package leetcode1054

import "sort"

func rearrangeBarcodes(barcodes []int) []int {
	if len(barcodes) == 0 {
		return []int{}
	}

	cnt := [10001]int{}

	for _, bc := range barcodes {
		cnt[bc]++
	}

	sort.Slice(barcodes, func(i, j int) bool {
		if cnt[barcodes[i]] == cnt[barcodes[j]] {
			return barcodes[i] < barcodes[j]
		}
		return cnt[barcodes[i]] > cnt[barcodes[j]]
	})

	res := make([]int, len(barcodes))
	i := 0
	for _, bc := range barcodes {
		res[i] = bc
		i += 2
		if i >= len(barcodes) {
			i = 1
		}
	}
	return res
}
