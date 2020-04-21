package leetcode378

func kthSmallest(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	n := len(matrix)

	res := make([]int, 0, n*n)
	for i := range matrix {
		res = mergeSort(res, matrix[i])
	}

	return res[k-1]
}

func mergeSort(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}

	if len(nums2) == 0 {
		return nums1
	}

	i, j := 0, 0

	var res []int
	for i < len(nums1) || j < len(nums2) {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] > nums2[j] {
				res = append(res, nums2[j])
				j++
			} else if nums1[i] < nums2[j] {
				res = append(res, nums1[i])
				i++
			} else {
				res = append(res, nums1[i])
				res = append(res, nums2[j])
				i++
				j++
			}
			continue
		}

		for i < len(nums1) {
			res = append(res, nums1[i])
			i++
		}

		for j < len(nums2) {
			res = append(res, nums2[j])
			j++
		}
	}
	return res
}
