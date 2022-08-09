package LeetCode_349_IntersectionOfTwoArrays

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m := make(map[int]struct{})

	for _, key := range nums1 {
		if _, ok := m[key]; ok {
			continue
		}
		m[key] = struct{}{}
	}

	res := make([]int, 0, len(m))
	for _, key := range nums2 {
		if _, ok := m[key]; ok {
			res = append(res, key)
			delete(m, key)
		}
	}

	return res
}

func intersection_1(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m := map[int]bool{}

	for _, num := range nums1 {
		m[num] = true
	}

	var res []int
	for _, num := range nums2 {
		if m[num] {
			res = append(res, num)
			m[num] = false
		}
	}

	return res
}

func intersection_2(nums1 []int, nums2 []int) []int {
	n1 := [1001]int{}
	n2 := [1001]int{}

	for i := range nums1 {
		n1[nums1[i]]++
	}

	for i := range nums2 {
		n2[nums2[i]]++
	}

	res := make([]int, 0, 1001)
	for i := 0; i < 1001; i++ {
		if n1[i] == 0 || n2[i] == 0 {
			continue
		}

		res = append(res, i)
	}
	return res
}

func intersection_3(nums1 []int, nums2 []int) []int {
	set := [1001]int{}

	for _, num := range nums1 {
		set[num]++
	}

	var res []int
	for _, num := range nums2 {
		if set[num] != 0 {
			res = append(res, num)
			set[num] = 0
		}
	}
	return res
}
