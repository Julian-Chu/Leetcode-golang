package leetcode451

import "sort"

func frequencySort(s string) string {
	var (
		counts = make([]int, 128)
		bin    = make([]byte, 0, 128)
		res    = make([]byte, 0, len(s))
	)

	for i := 0; i < len(s); i++ {
		counts[s[i]]++
	}

	for k := byte(0); k < 128; k++ {
		if counts[k] > 0 {
			bin = append(bin, k)
		}
	}

	sort.Slice(bin, func(i, j int) bool {
		return counts[bin[i]] > counts[bin[j]]
	})

	for i := 0; i < len(bin); i++ {
		ln := counts[bin[i]]
		for j := 0; j < ln; j++ {
			res = append(res, bin[i])
		}
	}

	return string(res)
}
