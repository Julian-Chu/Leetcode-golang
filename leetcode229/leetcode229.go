package leetcode229

func majorityElement(nums []int) []int {
	n := len(nums)
	m := make(map[int]int)

	for i := range nums {
		m[nums[i]]++
	}

	res := make([]int, 0)
	cnt := n / 3
	for k, v := range m {
		if v > cnt {
			res = append(res, k)
		}
	}

	return res
}
