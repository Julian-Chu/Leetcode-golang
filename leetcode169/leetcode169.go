package leetcode169

func majorityElement(nums []int) int {
	x, cnt := 0, 0

	for _, v := range nums {
		switch {
		case x == v:
			cnt++
		case cnt > 0:
			cnt--
		default:
			x = v
		}
	}
	return x
}
