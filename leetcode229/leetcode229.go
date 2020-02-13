package leetcode229

func majorityElement(nums []int) []int {

	n := len(nums)
	if n == 1 {
		return nums
	}

	e1, e2, c1, c2 := 0, 1, 0, 0

	for _, v := range nums {
		switch {
		case e1 == v:
			c1++
		case e2 == v:
			c2++
		case c1 == 0:
			e1 = v
			c1 = 1
		case c2 == 0:
			e2 = v
			c2 = 1
		default:
			c1--
			c2--
		}
	}
	res := make([]int, 0, 2)
	c1 = 0
	for _, v := range nums {
		if v == e1 {
			c1++
		}
	}
	if c1 > n/3 {
		res = append(res, e1)
	}
	c2 = 0
	for _, v := range nums {
		if v == e2 {
			c2++
		}
	}
	if c2 > n/3 {
		res = append(res, e2)
	}
	return res
}
