package LeetCode_1356_SortIntegersByTheNumberOf1Bits

import "sort"

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		if countBits(arr[i]) == countBits(arr[j]) {
			return arr[i] < arr[j]
		}
		return countBits(arr[i]) < countBits(arr[j])
	})
	return arr
}

func countBits(num int) int {
	cnt := 0
	for num > 0 {
		cnt += num & 1
		num >>= 1
	}
	return cnt
}
