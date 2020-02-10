package leetcode209

func minSubArrayLen(s int, nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	res := length + 1
	for l, r, sum := 0, 0, nums[0]; l < length && r < length; {
		if sum < s {
			r++
			if r >= length {
				break
			}
			sum += nums[r]
			continue
		}

		if r-l+1 < res {
			res = r - l + 1
		}

		sum -= nums[l]
		l++
	}
	if res > length {
		return 0
	}

	return res
}
