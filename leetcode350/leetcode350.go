package leetcode350

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m1 := make(map[int]int)

	for _, num := range nums1 {
		m1[num]++
	}

	resMap := make(map[int]int)

	for _, num := range nums2 {
		cnt, ok := m1[num]
		if !ok {
			continue
		}
		if resMap[num] >= cnt {
			continue
		}

		resMap[num]++
	}

	res := make([]int, 0, len(nums1))

	for key, cnt := range resMap {
		for i := 0; i < cnt; i++ {
			res = append(res, key)
		}
	}

	return res
}
