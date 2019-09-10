package leetcode491

import (
	"strconv"
	"strings"
)

func findSubsequences(nums []int) [][]int {
	n1 := len(nums)
	table := make([][][]int, n1)
	res := make([][]int, 0)
	for i := 0; i < n1; i++ {
		table[i] = make([][]int, 0)
	}
	keys := make(map[string]bool)
	for i := 0; i < n1; i++ {
		for j := 0; j < i; j++ {
			if nums[i] >= nums[j] {
				first := []int{nums[j], nums[i]}
				key := IntToString(first)
				if _, v1 := keys[key]; !v1 {
					keys[key] = true
					res = append(res, first)
					table[i] = append(table[i], first)
				}
				for _, v2 := range table[j] {
					arr := make([]int, len(v2))
					copy(arr, v2)
					arr = append(arr, nums[i])
					key := IntToString(arr)
					if _, v3 := keys[key]; !v3 {
						keys[key] = true
						res = append(res, arr)
						table[i] = append(table[i], arr)
					}
				}
			}
		}
	}
	return res
}

func IntToString(nums []int) string {
	b := make([]string, len(nums))
	for i, v := range nums {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, ",")
}
