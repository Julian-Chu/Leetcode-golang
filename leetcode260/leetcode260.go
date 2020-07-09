package leetcode260

func singleNumber(nums []int) []int {
	set := make(map[int]int)

	for _, num := range nums {
		set[num]++
	}

	res := make([]int, 0)

	for key, n := range set {
		if n == 1 {
			res = append(res, key)
		}
		if len(res) == 2 {
			return res
		}
	}

	return res
}
