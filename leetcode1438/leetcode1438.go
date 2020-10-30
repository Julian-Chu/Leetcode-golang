package leetcode1438

func longestSubarray(nums []int, limit int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	var s, e, ans int = 0, 1, 1
	var min, max []int = []int{0}, []int{0}

	for e < l {
		for len(min) > 0 && nums[e] < nums[min[len(min)-1]] {
			min = min[:len(min)-1]
		}

		for len(max) > 0 && nums[e] > nums[max[len(max)-1]] {
			max = max[:len(max)-1]
		}

		min = append(min, e)
		max = append(max, e)

		for nums[max[0]]-nums[min[0]] > limit {
			s++
			if min[0] < s {
				min = min[1:]
			}
			if max[0] < s {
				max = max[1:]
			}
		}

		if ans < e-s+1 {
			ans = e - s + 1
		}
		e++
	}
	return ans
}
