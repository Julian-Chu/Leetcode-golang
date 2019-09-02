package leetcode334

func increasingTriplet(nums []int) bool {
	n := len(nums)
	seq := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				seq[i] = max(seq[j]+1, seq[i])
			}
		}
		if seq[i]+1 >= 3 {
			return true
		}
	}
	return false
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//func increasingTriplet(nums []int) bool {
//func increasingTriplet(nums []int) bool {
//	seq := make([]int, 0, len(nums))
//	for _, v := range nums {
//		at := sort.SearchInts(seq, v)
//		if at == len(seq) {
//			seq = append(seq, v)
//		} else if seq[at] > v {
//			seq[at] = v
//		}
//		if len(seq) >= 3 {
//			fmt.Printf("%v", seq)
//			return true
//		}
//	}
//	return false
//}

//func increasingTriplet(nums []int) bool {
//	n1 := len(nums)
//	if n1 == 0 {
//		return false
//	}
//	seq := []int{nums[0]}
//	for _, v := range nums {
//		if v <= seq[0] {
//			seq[0] = v
//		} else if v > seq[len(seq)-1] {
//			seq = append(seq, v)
//		} else {
//			l, r := 0, len(seq)-1
//			for r-l > 1 {
//				mid := (l + r) / 2
//				if v > seq[mid] {
//					l = mid
//				} else if v < seq[mid] {
//					r = mid
//				}
//			}
//			seq[r] = v
//		}
//		if len(seq) >= 3 {
//			return true
//		}
//	}
//	return false
//
//}
