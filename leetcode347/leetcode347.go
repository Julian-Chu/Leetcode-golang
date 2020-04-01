package leetcode347

import "sort"

func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}
	res := make([]int, 0, k)
	if len(m) <= k {
		for key := range m {
			res = append(res, key)
		}
		return res
	}

	minCnt := 1 << 31
	var keyMinCnt = 0
	for key, cnt := range m {
		if len(res) < k {
			if cnt < minCnt {
				minCnt = cnt
				keyMinCnt = key
			}

			res = append(res, key)
			continue
		}

		if cnt > minCnt {
			minCnt = 1 << 31
			for i := 0; i < k; i++ {
				if res[i] == keyMinCnt {
					res[i] = key
					break
				}
			}
			for i := 0; i < k; i++ {
				c := m[res[i]]
				if c < minCnt {
					minCnt = c
					keyMinCnt = res[i]
				}
			}
		}
	}

	sort.Ints(res)
	return res
}
