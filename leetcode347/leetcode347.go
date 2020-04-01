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
		sort.Ints(res)
		return res
	}

	type kp struct {
		key int
		cnt int
	}
	kps := make([]kp, 0, len(m))
	for key := range m {
		kps = append(kps, kp{
			key: key,
			cnt: m[key],
		})
	}

	sort.Slice(kps, func(i, j int) bool {
		if kps[i].cnt > kps[j].cnt {
			return true
		}
		return false
	})

	kps = kps[0:k]
	for i := range kps {
		res = append(res, kps[i].key)
	}

	sort.Ints(res)
	return res
}
