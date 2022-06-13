package LeetCode_763_PartitionLables

import "sort"

func partitionLabels(s string) []int {
	m := map[byte][]int{}

	for i := range s {
		b := s[i]
		if _, ok := m[b]; ok {
			continue
		}

		j := len(s) - 1
		for ; j > i && s[j] != b; j-- {

		}

		m[b] = []int{i, j}
	}

	r := [][]int{}
	for _, interval := range m {
		r = append(r, interval)
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i][0] < r[j][0]
	})

	start := r[0][0]
	end := r[0][1]
	res := []int{}
	for i := 1; i < len(r); i++ {
		if r[i][0] < end {
			end = max(r[i][1], end)
		} else {
			res = append(res, end-start+1)
			start = r[i][0]
			end = r[i][1]
		}
	}

	res = append(res, end-start+1)
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func partitionLabels_1(s string) []int {
	var marks [26]int
	left, right := 0, 0
	for i := range s {
		marks[s[i]-'a'] = i
	}

	var res []int
	for i := range s {
		right = max(right, marks[s[i]-'a'])
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}
	return res
}
