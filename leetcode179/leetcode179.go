package leetcode179

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	numStrs := make([]string, len(nums))

	isZeros := true
	for i := range nums {
		if nums[i] != 0 {
			isZeros = false
		}
		numStrs[i] = strconv.Itoa(nums[i])
	}
	if isZeros {
		return "0"
	}

	sort.Slice(numStrs, func(i, j int) bool {
		return numStrs[i]+numStrs[j] > numStrs[j]+numStrs[i]
	})

	return strings.Join(numStrs, "")
}
