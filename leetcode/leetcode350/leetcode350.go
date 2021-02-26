package leetcode350

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m1 := make(map[int]int)

	for _, num := range nums1 {
		m1[num]++
	}

	res := make([]int, 0, len(nums1))
	for _, num := range nums2 {
		cnt, ok := m1[num]
		if !ok {
			continue
		}
		if cnt <= 0 {
			continue
		}

		m1[num]--
		res = append(res, num)

	}

	return res
}
