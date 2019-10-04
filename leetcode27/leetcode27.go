package leetcode27

import "fmt"

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1

	for {
		for i < len(nums) && nums[i] != val {
			i++
		}

		for j >= 0 && nums[j] == val {
			j--
		}
		if j < i {
			break
		}
		fmt.Println(i, j)
		nums[i], nums[j] = nums[j], nums[i]
	}

	return i
}
