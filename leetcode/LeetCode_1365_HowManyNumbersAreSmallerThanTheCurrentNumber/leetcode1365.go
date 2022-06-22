package LeetCode_1365_HowManyNumbersAreSmallerThanTheCurrentNumber

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {

	m := make(map[int]int)

	cloned := append([]int{}, nums...)

	sort.Ints(cloned)

	for i, num := range cloned {
		if _, ok := m[num]; !ok {
			m[num] = i - 0
		}
	}

	res := make([]int, len(nums))

	for i, num := range nums {
		res[i] = m[num]
	}
	return res
}
