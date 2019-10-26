package leetcode88

func merge(nums1 []int, m int, nums2 []int, n int) {

	nums3 := append([]int{}, nums1[:m]...)
	j, k := 0, 0
	for i := 0; i < m+n; i++ {
		switch {
		case j >= m:
			nums1[i] = nums2[k]
			k++
		case k >= n:
			nums1[i] = nums3[j]
			j++
		default:
			if nums2[k] >= nums3[j] {
				nums1[i] = nums3[j]
				j++
			} else {
				nums1[i] = nums2[k]
				k++
			}
		}
	}
}
