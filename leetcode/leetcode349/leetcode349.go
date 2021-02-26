package leetcode349

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m := make(map[int]bool)

	for _, key := range nums1 {
		if _, ok := m[key]; ok {
			continue
		}
		m[key] = false
	}

	for _, key := range nums2 {
		if _, ok := m[key]; ok {
			m[key] = true
			continue
		}
	}

	res := make([]int, 0, len(m))

	for key, v := range m {
		if v == true {
			res = append(res, key)
		}
	}

	return res

}
